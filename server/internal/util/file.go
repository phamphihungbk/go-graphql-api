package util

import (
	"fmt"
	"os"
	"path"
)

func OpenFile(directory string, fileName string) *os.File {
	fileLocation := path.Join(directory, fileName)
	src, err := os.OpenFile(fileLocation, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Can not open log file", err)
	}

	return src
}
