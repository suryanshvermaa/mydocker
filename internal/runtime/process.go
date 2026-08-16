package runtime

import (
	"os"
	"os/exec"
	"syscall"

	"golang.org/x/sys/unix"
)

func runParent(args []string) error {
	cmd := exec.Command(
		"/proc/self/exe",
		append([]string{"run", "child"},
			args...)...,
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// creating process
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: unix.CLONE_NEWPID | unix.CLONE_NEWUTS | unix.CLONE_NEWNS,
	}
	return cmd.Run()
}
