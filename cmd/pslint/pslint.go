package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	//"github.com/weppos/pslint"
	"github.com/weppos/pslint"
)

var (
	flagFile *string
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
	flagFile = flag.String("file", "", "Read the PSL from file")

	flag.Usage = usage
	flag.Parse()
}

func main() {
	linter := pslint.NewLinter()

	switch flag.NArg() {
	case 0:
		if flagFile != nil {
			lintFile(linter, *flagFile)
		} else {
			lintPipe(linter)
		}
	default:
		flag.Usage()
		os.Exit(2)
	}
}

func lintFile(linter *pslint.Linter, path string) {
	printLint(linter.LintFile(path))
}

func lintPipe(linter *pslint.Linter) {
	fi, _ := os.Stdin.Stat()
	if fi.Mode()&os.ModeNamedPipe == 0 {
		flag.Usage()
		os.Exit(2)
	} else {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		printLint(linter.LintString(string(bytes)))
	}
}

func printLint(problems []pslint.Problem, err error) {
	fmt.Println(problems)
	fmt.Println(err)
}
