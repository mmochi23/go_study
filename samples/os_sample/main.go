package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	for i, s := range os.Args {
		if i > 0 {
			fmt.Println(i, s)
		}
	}

	f, err := os.Open(os.Args[1])
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	fi, _ := f.Stat()
	fmt.Printf("name=%s size=%d mode=%s updated=%s dir=%v\n", fi.Name(), fi.Size(), fi.Mode(), fi.ModTime(), fi.IsDir())

	f2, _ := os.Create("bar.txt")
	fi2, _ := f2.Stat()
	fmt.Printf("name=%s size=%d mode=%s updated=%s dir=%v\n", fi2.Name(), fi2.Size(), fi2.Mode(), fi2.ModTime(), fi2.IsDir())
	f2.WriteAt([]byte("Golang"), 0)
	f2.Seek(0, os.SEEK_END)
	f2.WriteString("bar\n")
	f2.Close()

	f3, _ := os.OpenFile("bar.txt", os.O_RDWR|os.O_APPEND, 0666)
	f3.WriteString("Hello World!\n")
	f3.Close()

	err = os.Rename("bar.txt", "files/bar.txt.bak")
	if err != nil {
		fmt.Println(err)
	}

	err = os.Remove("bar.txt")
	if err != nil {
		fmt.Println(err)
	}

	err = os.Chdir("files")
	d, _ := os.Open(".")
	defer d.Close()
	dl, _ := d.Readdir(0)
	for i, f := range dl {
		fmt.Println(i, f.Name())
	}

	fmt.Println(os.TempDir())
	fmt.Println(os.Hostname())
	e := os.Environ()
	for i, k := range e {
		if i < 5 {
			fmt.Println(k)
		}
	}
}
