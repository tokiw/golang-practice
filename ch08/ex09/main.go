package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type disk struct {
	root   string
	nfiles int64
	nbytes int64
}

//!+
func main() {
	// ...determine roots...

	//!-
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileBytesChans := make([]chan int64, len(roots))

	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	dus := make([]disk, len(roots))
	for i, root := range roots {
		fileBytesChans[i] = make(chan int64)
		callWarkDir(root, fileBytesChans[i])
	loop:
		for {
			dus[i].root = roots[i]
			select {
			case size, ok := <-fileBytesChans[i]:
				if !ok {
					break loop // fileSizes was closed
				}
				dus[i].nfiles++
				dus[i].nbytes += size
			case <-tick:
				printDiskUsage(dus) // Wasteful output
			}
		}
	}

	printDiskUsage(dus) // final totals
}

//!-

func printDiskUsage(dus []disk) {
	for _, d := range dus {
		fmt.Printf("%s %d files  %.1f GB\n", d.root, d.nfiles, float64(d.nbytes)/1e9)
	}
}

func callWarkDir(root string, fileSizes chan<- int64) {
	var n sync.WaitGroup
	n.Add(1)
	go walkDir(root, &n, fileSizes)
	go func() {
		n.Wait()
		close(fileSizes)
	}()
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+walkDir
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

//!-walkDir

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
