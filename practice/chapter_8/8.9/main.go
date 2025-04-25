package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show progress messages")

func main() {

	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	
	var wg sync.WaitGroup	
	for _, root := range roots {
		wg.Add(1)
		go func (dir string)  {
			defer wg.Done()
			fileSizes := make(chan int64)
			go func ()  {
				walkDir(dir, fileSizes)
				close(fileSizes)
			}()

			var tick <-chan time.Time
			if *verbose {
				tick = time.Tick(500 * time.Millisecond)
			}
			var nfiles, nbytes int64
			loop: 
				for {
					select {
					case size, ok := <- fileSizes:
						if !ok {
							break loop
						}	
						nfiles++
						nbytes += size
					case <-tick:
						printDiskUsage(dir, nfiles, nbytes)
					}
				}
				printDiskUsage(dir, nfiles, nbytes)
		}(root)
	}

	wg.Wait()
}

func printDiskUsage(dir string, nfiles, nbytes int64) {
	fmt.Printf("%s: %d files  %.1f GB\n", dir, nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "8.9: %v\n", err)
		return nil
	}
	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "8.9: %v\n", err)
			return nil
		}
		infos = append(infos, info)
	}
	return infos
}