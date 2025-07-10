package commands

import "sort"

type cliCommand struct {
	name        string
	description string
	Callback    func(config *Config, parameters []string) error
}

func GetCommandsSorted() map[string]cliCommand {
	commands := GetCommands()
	keys := make([]string, 0, len(commands))
	sortedCommands := make(map[string]cliCommand)

	for key := range commands {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		sortedCommands[key] = commands[key]
	}

	return sortedCommands
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show this help message",
			Callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			Callback:    commandMapB,
		},
		"explore": {
			name:        "explore <location>",
			description: "Explore a given location",
			Callback:    commandExplore,
		},
	}
}
