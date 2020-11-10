package goversion

import "runtime"

//GetVersion Returns go version
func GetVersion() string {
	return runtime.Version()
}
