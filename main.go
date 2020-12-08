package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	src := os.Args[1]
	dst := os.Args[2]

	fmt.Println(src, dst)

	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fmt.Printf("visited file or dir: %q\n", path)

			f, err := w.Create(path)
			if err != nil {
				log.Fatal(err)
			}

			body, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatal(err)
			}

			_, err = f.Write(body)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(dst+"/demo.zip", buf.Bytes(), 0600)
}
