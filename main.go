package main

import (
	"fmt"
	iconv "github.com/djimenez/iconv-go"
	"io/ioutil"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("error: no input file")
		return
	}

	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := iconv.ConvertString(string(input), "gbk", "utf-8")
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(os.Args) == 3 {
		ioutil.WriteFile(os.Args[2], []byte(out), 0600)
		return
	}
	ioutil.WriteFile(os.Args[1]+".utf8", []byte(out), 0600)
}
