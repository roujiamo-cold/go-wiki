package main

import (
	"fmt"
	"os"
	"sort"

	digest "github.com/roujiamo-cold/go-wiki/golang.org/blog/pipelines/digestingtree"
)

func main() {
	// Calculate the MD5 sum of all files under the specified directory,
	// then print the results sorted by path name.
	m, err := digest.MD5AllBounded(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}
