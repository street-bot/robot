package core

import "fmt"

var VersionString string

func printVersion() {
	if VersionString == "" {
		VersionString = "development"
	}

	fmt.Println("Streetbot Robot, Version: " + VersionString)
}
