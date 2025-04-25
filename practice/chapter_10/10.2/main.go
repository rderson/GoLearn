package main

import (
	"fmt"
	"os"

	"practice/chapter_10/10.2/arch"
	_ "practice/chapter_10/10.2/arch/tar"
	_ "practice/chapter_10/10.2/arch/zip"
)

func main() {
	fmt.Println("10.2: create a generic function that can read ZIP and TAR archives.")
	fmt.Println()
	zipArchive := "test_archive.zip"
	tarArchive := "test_archive.tar"
	unsupportedArchive := "test_archive.7z"
	if err := arch.ReadArchive(zipArchive); err != nil {
		fmt.Fprintf(os.Stderr, "10.2: %v\n", err)
		os.Exit(1)
	}
	fmt.Println()
	if err := arch.ReadArchive(tarArchive); err != nil {
		fmt.Fprintf(os.Stderr, "10.2: %v\n", err)
		os.Exit(1)
	}
	fmt.Println()
	if err := arch.ReadArchive(unsupportedArchive); err != nil {
		fmt.Fprintf(os.Stderr, "10.2: %v\n", err)
		os.Exit(1)
	}
}






