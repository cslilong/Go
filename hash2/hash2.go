package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	TestFile := "a.txt"
	infile, inerr := os.Open(TestFile)
	if inerr == nil {
		md5h := md5.New()
		io.Copy(md5h, infile)
		fmt.Printf("%s md5 is : %x\n\n", TestFile, md5h.Sum([]byte("")))

		sha1h := sha1.New()
		io.Copy(sha1h, infile)
		fmt.Printf("%s sha1 is : %x\n\n", TestFile, sha1h.Sum([]byte("")))
	} else {
		fmt.Println(inerr)
		os.Exit(1)
	}
}
