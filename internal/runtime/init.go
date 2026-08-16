package runtime

import (
	"fmt"
	"os"
	"os/exec"
)

func runChild(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no command specified")
	}
	if err := setupNamespaces(); err != nil {
		return err
	}

	// here rootfs is relative path
	if err := setupRootfs("rootfs"); err != nil {
		return err
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
