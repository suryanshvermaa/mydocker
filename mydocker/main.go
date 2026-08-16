package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"golang.org/x/sys/unix"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: mydocker run <command>")
		os.Exit(1)
	}

	if os.Args[1] == "child" {
		child()
		return
	}
	switch os.Args[1] {
	case "run":
		run()
	default:
		fmt.Println("unknown command")
		os.Exit(1)
	}
}

func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// creating process
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: unix.CLONE_NEWPID | unix.CLONE_NEWUTS | unix.CLONE_NEWNS,
	}
	if err := cmd.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func child() {
	fmt.Println("running inside container")
	// set hostname
	if err :=
		unix.Sethostname([]byte("mycontainer")); err != nil {
		panic(err)
	}

	err := unix.Mount(
		"proc",
		"/proc",
		"proc",
		0,
		"",
	)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(os.Args[2])

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
