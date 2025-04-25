package zip

import (
	"archive/zip"
	"fmt"
	"practice/chapter_10/10.2/arch"
)

func ReadZipArchive(name string) error {
	zipListing, err := zip.OpenReader(name)
	if err != nil {
		return err
	}
	for num, file := range zipListing.File {
		fmt.Printf("File %d: %v\n", num+1, file.Name)
	}
	return nil
}

func init()  {
	arch.RegisterFormat("zip", ReadZipArchive)
}
