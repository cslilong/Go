package main

import (
	"fmt"
	"os"
	"os/exec"
)

var helloWorld = []byte(`package main

import "fmt"

func main() {
	fmt.Println("hello, world")
}
`)

var hellostr = `package main

import "fmt"

func main() {
	fmt.Println("hello, world")
}
`

func err(e error) bool {
	if e != nil {
		fmt.Println(e)
		return true
	}
	return false
}

func main() {
	f, e := os.Create("x.go")
	if err(e) {
		return
	}
	defer os.Remove("x.go")
	defer f.Close()

	//_, e = io.Copy(f, helloWorld)
	//_, e = f.Write(helloWorld)
	_, e = f.WriteString(hellostr)
	if err(e) {
		return
	}
	f.Close()

	cmd := exec.Command("go", "run", "x.go")

	o, e := cmd.CombinedOutput()
	if err(e) {
		return
	}
	fmt.Printf("%s", o)

}
