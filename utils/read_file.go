package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(day string, ensureTrailingLine bool) ([]string, error) {
	file := "real"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	fileName := day + "/" + file + ".txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("file ", fileName, " is not readable")
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	if ensureTrailingLine {
		if lines[len(lines)-1] != "" {
			lines = append(lines, "")
		}
	} else {
		for {
			if lines[len(lines)-1] != "" {
				break
			}

			lines = lines[0 : len(lines)-1]
		}
	}

	return lines, nil
}
