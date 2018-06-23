// Package exitcode : OS exit code enumeration
//
// These codes are licensed under CC0.
// http://creativecommons.org/publicdomain/zero/1.0/
package exitcode

import "os"

//ExitCode is OS exit code enumeration class
type ExitCode int

const (
	//Normal is OS exit code "normal"
	Normal ExitCode = iota
	//Abnormal is OS exit code "abnormal"
	Abnormal
)

var exitcodeMap = map[ExitCode]string{
	Normal:   "normal end",
	Abnormal: "abnormal end",
}

//Exit calls os.Exit()
func (c ExitCode) Exit() {
	os.Exit(int(c))
}

//Stringer method
func (c ExitCode) String() string {
	if str, ok := exitcodeMap[c]; ok {
		return str
	}
	return "unknown"
}
