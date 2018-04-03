# withd - a cross-platform alternative to pushd && command && popd

# EXAMPLES

```
$ pwd
/Users/andrew/go/src/github.com/mcandre/withd

$ ls names.txt
ls names.txt: No such file or directory

$ ls examples/
names.txt

$ withd examples/ ls names.txt
names.txt

$ withd examples/ ls *.txt
ls: *.txt: No such file or directory
2017/06/11 22:33:14 exit status 1
panic: exit status 1

goroutine 1 [running]:
panic(0xb7820, 0xc420010590)
        /usr/local/Cellar/go/1.7/libexec/src/runtime/panic.go:500 +0x1a1
log.Panic(0xc420043f08, 0x1, 0x1)
        /usr/local/Cellar/go/1.7/libexec/src/log/log.go:320 +0xc9
main.main()
        /Users/andrew/go/src/github.com/mcandre/withd/cmd/withd/main.go:46 +0x1ea

$ pwd
/Users/andrew/go/src/github.com/mcandre/withd

$ withd -help
  -help
        Show usage information
  -version
        Show version information
```

Note: The failure of withd to find `*.txt` in `examples/` is a side effect of how shells expand file pattern globs, which are relative to the current working directory, not the directory specified to `withd`.

# DOWNLOAD

https://github.com/mcandre/withd/releases

# DOCUMENTATION

https://godoc.org/github.com/mcandre/withd

# RUNTIME REQUIREMENTS

(None)

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.9+

## Recommended

* [Docker](https://www.docker.com/)
* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [goxcart](https://github.com/mcandre/goxcart) (e.g., `github.com/mcandre/goxcart/...`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/withd/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p "$GOPATH/src/github.com/mcandre"
$ git clone https://github.com/mcandre/withd.git "$GOPATH/src/github.com/mcandre/withd"
$ cd "$GOPATH/src/github.com/mcandre/withd"
$ git submodule update --init --recursive
$ go install ./...
```

# TEST

```
$ mage test
```

# PORT

```
$ mage port
```

# LINT

Keep the code tidy:

```
$ mage lint
```
