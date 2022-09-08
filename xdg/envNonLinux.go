//go:build !linux

package xdg

// DataHome returns an empty string for non-Linux systems
func DataHome() string {
	return ""
}

// DataDirs returns an empty slice for non-Linux systems
func DataDirs() []string {
	return []string{}
}

// ConfigDirs returns an empty slice for non-Linux systems
func ConfigDirs() []string {
	return []string{}
}

// CacheHome returns an empty string for non-Linux systems
func CacheHome() string {
	return ""
}
