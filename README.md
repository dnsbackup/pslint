# pslint

PSLint is a linter for [Public Suffix list](https://publicsuffix.org/).

[![Build Status](https://travis-ci.org/weppos/pslint.svg?branch=master)](https://travis-ci.org/weppos/pslint)

PSLint is both a binary and a library. You can run the linter using the built-in `pslint` command or include the library as a dependency in your Go code to call the exported functions.


## Compile

You need Go to compile the binary.

```
$ cd /path/to/repo
$ go build
```

You can also [cross-compile for several platforms](https://github.com/mitchellh/gox).


## Usage

```shell
$ ./pslint

Usage of pslint:
    pslint [flags] # check from stdin
    pslint [flags] --file path # check the content of file
Flags:
  -fail-fast
        Stop checking on first error
  -fail-first
        Stop checking line on first error (default true)
  -file string
        Read the PSL from file
```

**Examples**

```
$ ./pslint --file lists/master.txt
Found 0 problems!
```

```
$ cat lists/master.txt | ./pslint
Found 0 problems!
```

```
$ ./pslint --file lists/errors.txt
Found 5 problems:
 40:  airline.aero | leading space        (warning)
183: aRpa          | non-lowercase suffix (error)
243:  sa.gov.au    | leading space        (warning)
327: q.BG          | non-lowercase suffix (error)
507: cOm.by        | non-lowercase suffix (error)
```


## License

Copyright (c) 2016 Simone Carletti. This is Free Software distributed under the MIT license.
