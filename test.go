package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// GetRelativePath - returns paths joined onto the releative executable
func GetRelative(paths ...string) string {
	// get the currently running executable path
	base := BasePath()
	if len(paths) > 0 {
		// append the paths to the relative base path
		loc := base
		for i, _ := range paths {
			filepath.Join(loc, paths[i])
		}
		return loc
	}
	return filepath.Join(base, paths[0])
}

// BasePath - return the path where the app is running
func BasePath() string {
	ex, _ := os.Executable()
	return filepath.Dir(ex)
}

func testConfig() {
	configFile := filepath.Join(BasePath(), "config-files", "config.yaml")
	if inf, _ := os.Stat(configFile); inf != nil {
		fmt.Printf("Name: %s\n", inf.Name())
		fmt.Printf("Is Dir: %x\n", inf.IsDir())
		fmt.Printf("Size: %x\n", inf.Size())
	}
	
	bytes, err := os.ReadFile(configFile); if err != nil {
		log.Fatalf("erroring reading .yaml: %s", err)
	} else {
		log.Println(string(bytes))
	}
}

func main() {
	testConfig()
}