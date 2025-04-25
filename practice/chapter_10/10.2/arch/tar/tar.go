package tar

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"practice/chapter_10/10.2/arch"
)

func ReadTarArchive(name string) error {
	arch, err := os.Open(name)
	defer arch.Close()
	if err != nil {
		return err
	}
	tarListing := tar.NewReader(arch)
	i := 0
	for {
		file, err := tarListing.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("File %d: %v\n", i+1, file.Name)
		i += 1
	}
	return nil
}

func init()  {
	arch.RegisterFormat("tar", ReadTarArchive)
}