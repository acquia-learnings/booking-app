# booking-app
learning go with Nana



```gotemplate
go list ... 
```

Executing in any folder lists all the packages, including packages of the standard library first followed by external 
libraries in your go workspace.

```gotemplate
go install app
```

looks for `app` subdirectory inside `src` directory of `GOPATH` (since `GOROOT` doesn’t have it).


There is a set of programs to build and process Go source code. Instead of being run directly, programs in that set are 
usually invoked by the go program. `GOPATH` and `GOROOT` are environment variables that define a certain arrangement and 
organization for the Go source code. The paths of `GOPATH` and `GOROOT` can be modified explicitly if required.

```gotemplate
go env GOROOT
go env GOPATH
```

to check the current `GOROOT`, `GOPATH`.

`GOPATH`, also called the workspace directory, is the directory where the Go code belongs. It is implemented by and 
documented in the go/build package and is used to resolve import statements. The go get tool downloads packages to the 
first directory in `GOPATH`. If the environment variable is unset, `GOPATH` defaults to a subdirectory named “go” in the 
user’s home directory.

```
└── GOPATH
    ├── src     # holds source code. The path below this directory determines the import path or the executable name.
    ├── pkg     # holds installed package objects. Each target OS and architecture pair has its own subdirectory of pkg.
    └── bin     # holds compiled commands. Every command is named for its source directory.
```

__Note:__ When using modules in Go, the `GOPATH` is no longer used to determine imports. However, it is still used to 
store downloaded source code in `pkg` and compiled commands `bin`.

`GOROOT` is for compiler and tools that come from go installation and is used to find the standard libraries. It should 
always be set to the installation directory.

__Note:__ It is possible to install the Go tools to a different location. This can be done by setting the `GOROOT` 
environment variable to point to the directory in which it was installed, although this is not recommended as it comes 
preset with the tooling.