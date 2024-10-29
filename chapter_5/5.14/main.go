package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func breadthFirst(f func(item string, depth int,) []string, worklist []string) {
	seen := make(map[string]bool)
	depthMap := make(map[string]int)
	depthMap[worklist[0]] = 0

	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				newItems := f(item, depthMap[item])
				for _, newItem := range newItems {
					worklist = append(worklist, newItem)
					depthMap[newItem] = depthMap[item] + 1
				}
			}
		}
	}
}


func directoryTree(dirname string) []string {
	var results []string

	entries, err := os.ReadDir(dirname)
	if err != nil {
		for _, entry := range entries {
			if !entry.IsDir() {
				results = append(results, filepath.Join(dirname, entry.Name()))
			} else {
				log.Printf("Error reading directory: %v", err)
			}
		}
	}

	for _, entry := range entries {
		results = append(results, filepath.Join(dirname, entry.Name()))
		if entry.IsDir() {
			results = append(results, directoryTree(filepath.Join(dirname, entry.Name()))...)
		}
	}

	return results
}

func printItem(item string, depth int) {
	prefix := ""
	for i := 0; i < depth-1; i++ {
		prefix += "│   "
	}

	if depth > 0 {
		prefix += "└── "

	}

	fmt.Printf("%s%s\n", prefix, filepath.Base(item))
}

func main()  {
	path := "C:\\Users\\4uro4ka\\Desktop\\GoLearn"

	root := []string{path}

	fmt.Println(root[0])
	breadthFirst(func(item string, depth int) []string {
		printItem(item, depth)
		return directoryTree(item)
	}, root)
}