package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

const mandir string = "/usr/share/man/man1/"

func checkForErrors(e error) {
	if e != nil {
		log.Printf("Error occured %e", e)
		os.Exit(1)
	}
}

func createDir(dirname string) {
	_, err := os.Stat(dirname)

	if err != nil {
		err := os.Mkdir(dirname, os.ModePerm)

		if err != nil {
			log.Fatal("Could not create directory")
		}

		log.Printf("Succesfully created a directory %s", dirname)
	} else {
		log.Printf("Path %s already exists", dirname)
	}
}

func readAndCopyGzip(sourcePath string, destinationPath string) {
	// open the file and uncompress it through gzip reader
	f, err := os.Open(sourcePath)
	checkForErrors(err)

	defer f.Close()

	gr, err := gzip.NewReader(f)
	checkForErrors(err)

	defer gr.Close()

	// create destination file and copy the contents of uncompressed file to it
	dest, err := os.Create(destinationPath)
	checkForErrors(err)

	defer dest.Close()

	nbytes, err := io.Copy(dest, gr)
	checkForErrors(err)

	log.Printf("Written %d bytes to %s", nbytes, destinationPath)

	dest.Close()
	gr.Close()
	f.Close()
}

func main() {
	data, err := os.ReadDir(mandir)
	checkForErrors(err)

	// create destination directory to output unzipped files
	createDir("files")

	// remove the .gz extension from the file when extracting
	r, _ := regexp.Compile(".gz$")

	for i := 0; i < len(data); i++ {
		filePath := filepath.Join(mandir, data[i].Name())
		destFileName := r.ReplaceAllString((data[i].Name()), "")
		destFilePath := filepath.Join("./files", destFileName)

        readAndCopyGzip(filePath, destFilePath);
	}
}
