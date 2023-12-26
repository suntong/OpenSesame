package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// ".JPG", ".MOV" is for PowerShot SX40
var validMedias = []string{".JPG", ".MOV", ".jpg", ".mov", ".mp4"}

// dirPath: Directory to walk through
func listMedias(dirPath, p string) string {
	var files []string

	// Walk through the directory and collect file names
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && slices.Contains(validMedias, filepath.Ext(info.Name())) {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking through directory:", err)
		return ""
	}

	cutl := len(dirPath) // cut length from the left of the path
	if cutl == 2 {       // special case when dirPath is just './'
		cutl = 0
	}
	b := &strings.Builder{}
	fmt.Fprintf(b, "[")
	for i, f := range files {
		if i != 0 {
			fmt.Fprintf(b, ",")
		}
		fmt.Fprintf(b, "{src: '%s/%s'}", p, f[cutl:])
	}
	fmt.Fprintf(b, "]")
	return b.String()
}
