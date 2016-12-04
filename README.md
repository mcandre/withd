# withd - a cross-platform alternative to pushd && command && popd

# EXAMPLE

```
$ pwd
/Users/andrew/go/src/github.com/mcandre/withd

$ ls *.txt
ls *.txt: No such file or directory

$ withd examples/ ls *.txt
names.txt

$ pwd
/Users/andrew/go/src/github.com/mcandre/withd
```

# DOWNLOAD

https://github.com/mcandre/withd/releases

# REQUIREMENTS

* [Go](https://golang.org) 1.7+ with [$GOPATH configured](https://gist.github.com/mcandre/ef73fb77a825bd153b7836ddbd9a6ddc)

## Optional

* [coreutils](https://www.gnu.org/software/coreutils/coreutils.html)
* [make](https://www.gnu.org/software/make/)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [gox](https://github.com/mitchellh/gox) (e.g. `go get github.com/mitchellh/gox`)
* [shlint](https://rubygems.org/gems/shlint)
* [shellcheck](http://hackage.haskell.org/package/ShellCheck)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/withd/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone git@github.com:mcandre/withd.git $GOPATH/src/github.com/mcandre/withd
$ cd $GOPATH/src/github.com/mcandre/withd
$ git submodule update --init --recursive
$ sh -c 'cd cmd/withd && go install'
```

# TEST

```
$ make integration-test
```

# PORT

```
$ make port
```

# LINT

Keep the code tidy:

```
$ make lint
```

# GIT HOOKS

See `hooks/`.
