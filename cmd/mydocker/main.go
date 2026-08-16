package main

import (
	"fmt"
	"os"

	"github.com/suryanshvermaa/mydocker/internal/runtime"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: mydocker <command>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("usage mydocker run <command>")
			os.Exit(1)
		}
		if err := runtime.Run(os.Args[2:]); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("unknown command")
		os.Exit(1)
	}
}
