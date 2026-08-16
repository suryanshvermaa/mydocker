package types

type ContainerConfig struct {
	ID       string
	Command  []string
	RootFS   string
	Hostname string
	// Memory   `uint64`
	// CPUQuota uint64
	// Network  string
	// Volumes  []Volume
}
