package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	//"github.com/weppos/pslint"
)

func usage() {
	bin := path.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", bin)
	fmt.Fprintf(os.Stderr, "\t%s [flags] # check from stdin\n", bin)
	fmt.Fprintf(os.Stderr, "\t%s [flags] --file path # check the content of file \n", bin)
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func init() {
	flag.Usage = usage
	flag.Parse()
}

func main() {
	switch flag.NArg() {
	case 0:
		fi, _ := os.Stdin.Stat()
		if fi.Mode()&os.ModeNamedPipe == 0 {
			flag.Usage()
			os.Exit(2)
		} else {
			bytes, _ := ioutil.ReadAll(os.Stdin)
			fmt.Print(string(bytes))
		}
	default:
		flag.Usage()
		os.Exit(2)
	}
}
