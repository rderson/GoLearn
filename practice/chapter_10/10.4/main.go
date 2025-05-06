package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Package struct {
	Name string		`json:"Name"`
 	Imps []string 	`json:"Imports"`
	Deps []string 	`json:"Deps"`
}

func main() {
	// Vice versa:

	// key := os.Args[1]

	// cmd := exec.Command("go", "list", "-f", "'{{ .Deps }}'", key)
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "10.4: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(string(output))

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <target-package> [...more targets]")
		os.Exit(4)
	}

	cmd := exec.Command("go", "list", "-json", "...")
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "10.4: %v\n", err)
		os.Exit(1)
	}

	dec := json.NewDecoder(strings.NewReader(string(output)))
	var workspacePkgs []Package

	for  {
		var pkg Package
		if err := dec.Decode(&pkg); err == io.EOF {
			break
		} else if err != nil{
			fmt.Fprintf(os.Stderr, "10.4: %v\n", err)
			os.Exit(2)
		}
		workspacePkgs = append(workspacePkgs, pkg)
	}

	// for _, pkg := range workspacePkgs {
	// 	fmt.Printf("Package %s\nImports: %v\nDeps:%v\n\n", pkg.Name, pkg.Imps, pkg.Deps)
	// }

	var finalPkgs []string
	
	for _, target := range os.Args[1:] {
		for _, pkg := range workspacePkgs {
			for _, imp := range pkg.Imps{
				if imp == target {
					finalPkgs = append(finalPkgs, pkg.Name)
				}
			}
			for _, dep := range pkg.Deps {
				if dep == target {
					finalPkgs = append(finalPkgs, pkg.Name)
				}
			}
		}
	}

	for _, pkg := range removeDuplicates(finalPkgs) {
		fmt.Println(pkg)
	}
}

func removeDuplicates(slice []string) (result []string) {
	seen := make(map[string]bool)

	for _, str := range slice {
		if !seen[str] {
			seen[str] = true
			result = append(result, str)
		}
	}

	return result
}
