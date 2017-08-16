// unzipOpenXML go prog to unzip a microsoft docx, xlsx,... files
//
// +build
//
package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var targetDir string
var debug bool
var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "%s [-o=outputDir] some.docx\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "")
	flag.PrintDefaults()
}

// init initializes flag variables and Usage
func init() {
	flag.Usage = Usage
	flag.StringVar(&targetDir, "o", "", "dir to unpack files to (optional), default: 'unzipped_<FILENAME>'")
	flag.BoolVar(&debug, "debug", false, "spills debug info when set to 'true', default: 'false'")
}

// main unpacks given archive
func main() {
	log.SetPrefix("unzipOpenXML: ")
	flag.Parse()
	fileName := flag.Arg(0)
	if targetDir == "" {
		targetDir = "unzipped_" + fileName
		os.MkdirAll(targetDir, 0700)
	}
	r, err := zip.OpenReader(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	if debug {
		log.Printf("opened zip archive '%s'\n", fileName)
	}

	// iterate through the files in the archive...
	for _, file := range r.File {
		if debug {
			log.Printf("processing file '%s' from archive.\n", file.Name)
		}
		// create a full qualified path for the file
		path := filepath.Join(targetDir, file.Name)
		os.MkdirAll(targetDir, 0700)
		// if it is a directory, create it and go on with next file.
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			log.Printf("directory '%s' created.\n", path)
			continue
		}
		// open file from the archive for reading.
		fileReader, err := file.Open()
		if err != nil {
			log.Println(err)
			continue
		}
		defer fileReader.Close()
		if debug {
			log.Printf("'%s' opened from archive.\n", file.Name)
		}

		// create a targetfile with appropriate permissions
		dir, _ := filepath.Split(path)
		os.MkdirAll(dir, 0700)
		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			log.Println(err)
			continue
		}
		defer targetFile.Close()
		if debug {
			log.Printf("targetFile '%s' created.\n", path)
		}

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			log.Println(err)
			continue
		}
		if debug {
			log.Printf("file content from '%s' copied to '%s'\n", file.Name, path)
		}
	}
	if debug {
		log.Printf("unpacking '%s' done.\n", fileName)
	}
}
