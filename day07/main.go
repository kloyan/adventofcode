package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const (
	maxSpace      = 70000000
	requiredSpace = 30000000
)

type directory struct {
	name    string
	size    int
	subdirs map[string]*directory
	parent  *directory
}

func main() {
	root := buildFs(input)

	fmt.Println("Answer for Part 1:", totalBelowThreshold(root, 100000))

	threshold := requiredSpace - (maxSpace - root.size)
	fmt.Println("Answer for Part 2:", findSmallestSize(root, threshold))
}

func buildFs(input string) *directory {
	currDir := &directory{name: "/", subdirs: make(map[string]*directory)}
	root := currDir

	for _, str := range strings.Split(strings.TrimSpace(input), "\n")[1:] {
		switch {
		case strings.HasPrefix(str, "$ cd"):
			currDir = changeDir(currDir, str)
		case strings.HasPrefix(str, "dir"):
			addSubdir(currDir, str)
		case strings.HasPrefix(str, "$ ls"):
			continue
		default:
			var size int
			fmt.Sscanf(str, "%d", &size)
			updateDirSize(currDir, size)
		}
	}

	return root
}

func changeDir(dir *directory, str string) *directory {
	var dst string
	fmt.Sscanf(str, "$ cd %s", &dst)

	if dst == ".." {
		return dir.parent
	}

	return dir.subdirs[dst]
}

func addSubdir(dir *directory, str string) {
	var sub string
	fmt.Sscanf(str, "dir %s", &sub)

	dir.subdirs[sub] = &directory{name: sub, parent: dir, subdirs: make(map[string]*directory)}
}

func updateDirSize(dir *directory, size int) {
	dir.size += size
	if dir.parent != nil {
		updateDirSize(dir.parent, size)
	}
}

func totalBelowThreshold(dir *directory, threshold int) int {
	size := 0
	if dir.size <= threshold {
		size += dir.size
	}

	for _, sub := range dir.subdirs {
		size += totalBelowThreshold(sub, threshold)
	}

	return size
}

func findSmallestSize(dir *directory, threshold int) int {
	queue := make([]*directory, 0)
	queue = append(queue, dir)
	res := dir.size

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.size >= threshold && curr.size < res {
			res = curr.size
		}

		for _, sub := range curr.subdirs {
			queue = append(queue, sub)
		}
	}

	return res
}
