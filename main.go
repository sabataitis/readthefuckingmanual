package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const mandir string = "/usr/share/man/man1/"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadDir(mandir)

	check(err)

	var first os.DirEntry = data[3]
	const dest = "unzipped"

	firstFilePath := filepath.Join(mandir, first.Name())

	fmt.Println(firstFilePath, first.Name())

	f, err := os.Open(firstFilePath)
	check(err)

	defer f.Close()

	gr, err := gzip.NewReader(f)
	check(err)

	defer gr.Close()

    const filesPath string = "./files/";

    // create files dir if not exists
    stat, err := os.Stat(filesPath);

    if(err != nil || !stat.IsDir()) {
        err := os.MkdirAll(filesPath, os.ModePerm);
        check(err);

        fmt.Println("Created directory", filesPath);
    }

    destFilePath := filepath.Join(filesPath, first.Name() + ".txt");
	destination, err := os.Create(destFilePath)
    check(err);

	defer destination.Close()

	nbytes, err := io.Copy(destination, gr)
	check(err)
	
	fmt.Println("Written", nbytes, "to", destFilePath)

	destination.Close()
	gr.Close()
	f.Close()

}
