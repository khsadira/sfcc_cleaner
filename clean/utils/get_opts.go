package utils

import "strings"

func GetOpts(optsName string) []string {
	return strings.Split(optsName, ",")
}
