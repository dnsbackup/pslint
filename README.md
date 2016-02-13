# pslint

PSLint is a linter for [Public Suffix list](https://publicsuffix.org/).


## Compile

You need Go to compile the binary.

```
$ cd /path/to/repo
$ go build
```

You can also [cross-compile for several platforms](http://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5).


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
$ ./pslint --file lists/master
Found 0 problems!
```

```
$ cat lists/master | ./pslint
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
