package helpers

import (
	"os"
	"strings"
)

func ReadInputFile(file string) []string {
	dat, err := os.ReadFile(file)
	if err != nil {
			panic(err)
	}

	return  strings.Split(string(dat), "\r\n")
}
