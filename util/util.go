package util

import (
	"os"
	"strconv"
)

func IncrementDownload(path string) string {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return path
	} else {
		modifier := 0

		for os.IsExist(err) {
			modifier++
		}
		path := path + "(" + strconv.Itoa(modifier) + ")"
		return path
	}
}
