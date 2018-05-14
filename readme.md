# Go starter

## Dependencies

- [Go (Golang)](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Make for windows](http://gnuwin32.sourceforge.net/packages/make.htm)

---

## Build Commands

### Complete Build:

This will set the gopath, pull all of the dependencies and build the project.

```
make
```

### Only Build:

This will set the gopath and build the project. This will not install dependencies.

```
make build
```

### Install Dependencies:

This will set the gopath, pull all of the dependencies and build the project.

```
make install
```

### Watch & Reload on Changes:

This will watch your `app` directory for changes and reload the app when detected.

```
make watch
```

### To run:

The project will build into the `/bin` folder.

```
./bin/app
/bin/app.exe
```

---

## Docker:

Using docker compose

```
docker-compose build
docker-compose up
```

---

## Development

The source code lives in `/src/app`.
Dependencies will be installed into the parent directory.

Imports must be absolute from the makefile.