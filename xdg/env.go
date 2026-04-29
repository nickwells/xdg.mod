package xdg

import (
	"os"
)

// ConfigHome now uses the standard library function os.UserConfigDir to
// provide the value. It returns the empty string if that function returns a
// non-nil error.
func ConfigHome() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		return ""
	}

	return dir
}

// CacheHome now uses the standard library function os.UserCacheDir to
// provide the value. It returns the empty string if that function returns a
// non-nil error.
func CacheHome() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		return ""
	}

	return dir
}
