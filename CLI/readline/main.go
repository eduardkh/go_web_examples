package main

// https://stackoverflow.com/questions/41895260/implementing-auto-autocomplete-for-cli-application

import (
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/chzyer/readline"
)

// completer defines which commands the user can use
var completer = readline.NewPrefixCompleter()

// categories holding the initial default categories. The user can  add categories.
var categories = []string{"Category A", "Category B", "Category C"}

var l *readline.Instance

func main() {

	// Initialize config
	config := readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold: true,
	}

	var err error
	// Create instance
	l, err = readline.NewEx(&config)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	// Initial initialization of the completer
	updateCompleter(categories)

	log.SetOutput(l.Stderr())
	// This loop watches for user input and process it
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		// Checking which command the user typed
		switch {
		// Add new category
		case strings.HasPrefix(line, "newCategory"):
			// Remove the "newCategory " prefix (including space)
			if len(line) <= 12 {
				log.Println("newCategory <NameOfCategory>")
				break
			}
			// Append everything that comes behind the command as the name of the new category
			categories = append(categories, line[12:])
			// Update the completer to make the new category available in the cmd
			updateCompleter(categories)
		// Program is closed when user types "exit"
		case line == "exit":
			goto exit
		// Log all commands we don't know
		default:
			log.Println("Unknown command:", strconv.Quote(line))
		}
	}
exit:
}

// updateCompleter is updates the completer allowing to add new command during runtime. The completer is recreated
// and the configuration of the instance update.
func updateCompleter(categories []string) {

	var items []readline.PrefixCompleterInterface

	for _, category := range categories {
		items = append(items, readline.PcItem(category))
	}

	completer = readline.NewPrefixCompleter(
		readline.PcItem("newEntry",
			items...,
		),
		readline.PcItem("newCategory"),
	)

	l.Config.AutoComplete = completer
}
