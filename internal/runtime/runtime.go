package runtime

import (
	"fmt"
)

func Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no command specified")
	}

	switch args[0] {
	case "child":
		return runChild(args[1:])
	default:
		return runParent(args)
	}
}
