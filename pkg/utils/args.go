package utils

import (
	"fmt"
	"runtime"
)

// RebuildArgs move "--help" or "-h" flag to the end of the arg list
func RebuildArgs(args []string) []string {
	helpFlag := "--help"
	argList := []string{}
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			helpFlag = arg
		} else {
			argList = append(argList, arg)
		}
	}
	argList = append(argList, helpFlag)
	return argList
}

// PlatformString returns a string identifying the current system.
func PlatformString() string {
	return fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
}
