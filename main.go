package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	dir_to_scan := os.Args[1]
	if os.Args[0] == "" {
		log.Fatal("You need to add a valid path")
	}
	minWidth, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	minHeight, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal("You need to add a valid min height")

	}

	_err := filepath.Walk(dir_to_scan, func(path string, info os.FileInfo, err error) error {

		if reader, err := os.Open(path); err == nil {
			defer reader.Close()
			im, _, err := image.DecodeConfig(reader)
			if err != nil {
				// fmt.Fprintf(os.Stderr, "%s: %v\n", imgFile.Name(), err)
			}
			if im.Width > 0 && (im.Width < minWidth || im.Height < minHeight) {
				fmt.Printf("%s %d %d\n", path, im.Width, im.Height)
				e := os.Remove(path)
				if e != nil {
					log.Fatal(e)
				}
			}
		} else {
			fmt.Println("Impossible to open the file:", err)
		}

		return nil
	})
	if _err != nil {
		log.Println(err)
	}
}
