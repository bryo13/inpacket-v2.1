package ips

import "regexp"

var (
	classAPrivate, classBPrivate, classCPrivate *regexp.Regexp
)

func init() {
	classAPrivate = regexp.MustCompile(`^10\.(([1-9]?\d|[12]\d\d)\.){2}([1-9]?\d|[12]\d\d)$`)
	classBPrivate = regexp.MustCompile(`^172\.(1[6-9]|2\d|3[01])(\.([1-9]?\d|[12]\d\d)){2}$`)
	classCPrivate = regexp.MustCompile(`^192\.16[6-8](\.([1-9]?\d|[12]\d\d)){2}$`)
}

// CheckIPClass checks the IP class read from the chosen interface
func CheckIPClass(input string) {
	// First get the packet
	// Read its IP
	// check Ip class
}
