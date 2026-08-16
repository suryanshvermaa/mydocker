package runtime

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func setupNamespaces() error {
	if err :=
		unix.Sethostname([]byte("mycontainer")); err != nil {
		return fmt.Errorf("set hostname: %w", err)
	}
	return nil
}
