package completer

import (
	"strings"

	"github.com/BANKEX/plasma-research/src/node/verifier/cli/options"
	"github.com/c-bata/go-prompt"
)

func Completer(d prompt.Document) []prompt.Suggest {
	args := strings.Split(d.TextBeforeCursor(), " ")
	return argumentsCompleter(args)
}

func argumentsCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(options.InitialOptions, args[0], true)
	}
	first := args[0]
	switch first {
	case "eth":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(options.EthOptions, args[1], true)
		}
	case "plasma":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(options.PlasmaOptions, args[1], true)
		}
	}
	return []prompt.Suggest{}
}
