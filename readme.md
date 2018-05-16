# Golong - A Simple Template

#### A Go starter template, built on top of the Gin framework, with a focus on go modules and API development

<br>

![Gopher](https://i.imgur.com/yMWOR71.png)

---

## Dependencies

- [Go (Golang)](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Make](http://gnuwin32.sourceforge.net/packages/make.htm)
- [GNU CoreUtils](http://gnuwin32.sourceforge.net/packages/coreutils.htm)

---

### Windows Users

If you have the Choco package manager, the following command should install the environment dependencies

```
choco install make gnuwin32-coreutils.install -y
```
<i>You may need to add `C:\Program Files (x86)\GnuWin32\bin` to your path</i>

Otherwise use the following download links to manually install<br>
- [Make](http://gnuwin32.sourceforge.net/packages/make.htm)
- [GNU CoreUtils](http://gnuwin32.sourceforge.net/packages/coreutils.htm)

<i>Note: You may use git-bash, cygwin or WSL as these come with the required utilities</i>

---

## Getting started

This uses Go 1.11's modules. This project does not need to be in your GOPATH.

- Download this repo as a zip
- Extract it where you'd like
- Run `make`

---

## Build Commands

### Complete Build:

This will pull all of your dependencies and build the project into your OS's binary format.

```
make
```

### Only Build:

This will build the project. This will not install dependencies.

```
make build
```

### Install Dependencies:

This will pull all of the dependencies and build the project.

```
make get
```

### To run:

The project will build into the `/bin` folder.

```
make run
```

---

## Docker:

Included is a basic dockerfile, I will expand it to include docker-based development

---

## Development

The source code lives in `/src`.
Dependencies will be installed using go modules available from `go 1.11` and above.

Imports must be from the `go.mod`.

Folder structure:

```
src/
    cmd/
        servi/
            **/* -> While this isn't included, you can create a
                 -> comand line interface here
        servid/
            routes/
                handlers/
                    hander-name/ 
                        handler.go -> Handler for your HTTP requests 
                        response.go
                        request.go
                routes.go -> Your routes are defined here
                lib.go
        main.go -> This is your bootstrapping go project
    platform/
        **/* -> You add your platform specific code here
             -> Anything too general can be split up into it's own package 
    .env
    go.mod -> Dependencies are defined here
``` 