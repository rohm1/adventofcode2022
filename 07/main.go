package main

import (
	"aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type dir struct {
	name   string
	parent *dir
	dirs   map[string]*dir
	files  map[string]int
}

func createDir(name string, parent *dir) *dir {
	newDir := dir{name: name, parent: parent}
	newDir.dirs = map[string]*dir{}
	newDir.files = map[string]int{}
	return &newDir
}

const MaxDirSize = 100_000
const UpdateSize = 30_000_000
const FsSize = 70_000_000

func interpretCommand(lines []string, index int, currentDir *dir, rootDir *dir) (*dir, int) {
	args := strings.Split(lines[index], " ")
	switch args[1] {
	case "cd":

		switch args[2] {
		case "..":
			return interpretCommand(lines, index+1, currentDir.parent, rootDir)
		case "/":
			return interpretCommand(lines, index+1, rootDir, rootDir)
		default:
			return interpretCommand(lines, index+1, currentDir.dirs[args[2]], rootDir)
		}
	case "ls":
		for {
			index++
			if lines[index] == "" {
				return currentDir, index
			}

			args = strings.Split(lines[index], " ")
			if args[0] == "$" {
				return interpretCommand(lines, index, currentDir, rootDir)
			}

			if args[0] == "dir" {
				currentDir.dirs[args[1]] = createDir(args[1], currentDir)
			} else {
				size, _ := strconv.Atoi(args[0])
				currentDir.files[args[1]] = size
			}
		}
	}

	return currentDir, index
}

func calcDirSize(currentDir *dir) int {
	size := 0
	for _, fileSize := range currentDir.files {
		size += fileSize
	}
	for _, childDir := range currentDir.dirs {
		size += calcDirSize(childDir)
	}
	return size
}

func calcSumSmallDirs(currentDir *dir) int {
	sum := 0

	currentDirSize := calcDirSize(currentDir)
	if currentDirSize <= MaxDirSize {
		sum += currentDirSize
	}

	for _, childDir := range currentDir.dirs {
		sum += calcSumSmallDirs(childDir)
	}

	return sum
}

func deleteSmallestDirForUpdate(currentDir *dir, freeSize int) int {
	size := calcDirSize(currentDir)
	for _, childDir := range currentDir.dirs {
		smallestDirToDeleteForUpdate := deleteSmallestDirForUpdate(childDir, freeSize)
		if smallestDirToDeleteForUpdate < size && freeSize+smallestDirToDeleteForUpdate >= UpdateSize {
			size = smallestDirToDeleteForUpdate
		}
	}
	return size
}

func main() {
	lines, err := utils.ReadFile("07", false)
	if err != nil {
		return
	}

	root := createDir("/", nil)
	interpretCommand(lines, 1, root, root)

	fmt.Println("Part 1: ", calcSumSmallDirs(root))
	fmt.Println("Part 2: ", deleteSmallestDirForUpdate(root, FsSize-calcDirSize(root)))
}
