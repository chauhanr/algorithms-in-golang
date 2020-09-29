# Go Modules

**Go package** is constructed from one or more source files that together declare constants, types ,
functions and variables belonging to the package and which are accessible in all files of the same
pakage. 

**Go Module** is a collection of go packages store in a file tree with a go.mod file at its root.
The `go.mod` file defines the module path which is also the import path used for the root directory. 
Each dependency requirement is written as module path and a specific semantic version. 

To make a project a module we need to run the following command 
```
go mod init github.com/chauhanr/golang-web   # this needs to run on root of golang-web 
					     # this will create a go.mod and go.sum files. 
```

Each time there is a package imported in your code that has a go.mod it will be added to your
project's go.mod with a version. Another file `go.sum` is also added to the root folder of the
package that holds the hashes of the package that you have downloaded and using. 

To list all the modules used by a golang project we can use the following command: 
```
$ go list -m all 
```
In order to get the modules latest or a specific version we can use the following commands.
```
$ go get golang.org/x/test    # gets the latest tagged version of the test module from git repo 

$ go get rsc.io/sampler@v1.3.1  # gets this specific version of the package in question.  
```
In order to get rid of modules and packages that are no longer used in the go.mod we can run the
following command: 

```
$ go mod tidy
```
