package main

import (
	"os"
	"path"
	"regexp"
)

// IsValidDiscordID - returns if a string is a valid channel, user, role, emoji ID etc.
func IsValidDiscordID(str string) bool {
	re := regexp.MustCompile("^[0-9]{18,19}$")
	return re.MatchString(str)
}

// GetRelativePath - returns paths joined onto the releative executable
func GetRelativePath(paths ...string) string {
	// get the currently running executable path
	binary, _ := os.Executable()
	loc := binary
	if len(paths) > 0 {
		// append the paths to the relative base path
		for _, i := range paths {
			path.Join(loc, i)
		}
		return path.Join(path.Base(binary))
	}
	return path.Join(path.Base(binary), paths[0])
}
