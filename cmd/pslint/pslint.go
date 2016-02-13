package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

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
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(2)
	}

	if len(problems) == 0 {
		return
	}

	maxs := []int{0, 0, 0, 0}
	rows := [][]string{}
	for _, problem := range problems {
		row := []string{fmt.Sprintf("%v", problem.Line), problem.Message, fmt.Sprintf("%v", problem.Level), problem.LineSource}
		rows = append(rows, row)

		for i := 0; i < 4; i += 1 {
			if n := len(row[i]); n > maxs[i] {
				maxs[i] = n
			}
		}
	}

	fmt.Printf("Found %d problems:\n", len(problems))

	for _, row := range rows {
		fmt.Printf("%v: %v | %v (%v)\n",
			rightPad(row[0], maxs[0]),
			rightPad(row[3], maxs[3]),
			rightPad(row[1], maxs[1]),
			row[2])
	}

	os.Exit(1)
}

func leftPad(s string, length int) string {
	return strings.Repeat(" ", length-len(s)) + s
}

func rightPad(s string, length int) string {
	return s + strings.Repeat(" ", length-len(s))
}
