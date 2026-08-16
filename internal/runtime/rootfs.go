package runtime

import (
	"fmt"
	"path/filepath"

	"golang.org/x/sys/unix"
)

func setupRootfs(rootfs string) error {
	rootfs, err := filepath.Abs(rootfs)
	if err != nil {
		return fmt.Errorf("resolve rootfs: %w", err)
	}

	if err := makeMountsPrivate(); err != nil {
		return err
	}

	if err := pivotRoot(rootfs); err != nil {
		return err
	}

	if err := mountProc(); err != nil {
		return err
	}

	return nil
}

func makeMountsPrivate() error {
	if err := unix.Mount(
		"",
		"/",
		"",
		unix.MS_PRIVATE|unix.MS_REC,
		"",
	); err != nil {
		return fmt.Errorf("make mounts private: %w", err)
	}

	return nil
}

func bindMountRootfs(rootfs string) error {
	if err := unix.Mount(
		rootfs,
		rootfs,
		"",
		unix.MS_BIND|unix.MS_REC,
		"",
	); err != nil {
		return fmt.Errorf("bind mount rootfs: %w", err)
	}

	return nil
}

func pivotRoot(rootfs string) error {
	if err := bindMountRootfs(rootfs); err != nil {
		return err
	}

	oldRoot := filepath.Join(rootfs, "oldroot")

	if err := unix.Mkdir(oldRoot, 0700); err != nil {
		return fmt.Errorf("create oldroot: %w", err)
	}

	if err := unix.PivotRoot(rootfs, oldRoot); err != nil {
		return fmt.Errorf("pivot_root: %w", err)
	}

	if err := unix.Chdir("/"); err != nil {
		return fmt.Errorf("chdir: %w", err)
	}

	if err := unix.Unmount("/oldroot", unix.MNT_DETACH); err != nil {
		return fmt.Errorf("unmount oldroot: %w", err)
	}

	if err := unix.Rmdir("/oldroot"); err != nil {
		return fmt.Errorf("remove oldroot: %w", err)
	}

	return nil
}

func mountProc() error {
	if err := unix.Mount(
		"proc",
		"/proc",
		"proc",
		0,
		"",
	); err != nil {
		return fmt.Errorf("mount proc: %w", err)
	}

	return nil
}
