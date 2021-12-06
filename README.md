
# Fakestore Sandbox

## Introduction
This application is just a starting point on how to setup VSCode in order to debug a `golang` application (using [delve](https://github.com/go-delve/delve)) that runs inside a container.

## tl;dr

For the impatient ones just issue:

```
make copy-dvl-from-container
make run-application
```

## The remote API

The remote API that we use in order to fetch data is the following:
* https://fakestoreapi.com/docs

## Useful notes

### Start using modules

```
make make debug-container-development
I have no name!@fa5e69365006:/src$ go mod init github.com/giannoul/golang-fakestore
```
### Debugger 

First use the VSCode in order to create some breakpoints inside the source code. Issue:

```
make run-container-development
```

which will invoke `delve`. Then click Play in the VSCode debugger and that's it!

Since `delve` will ignore `Ctrl+C` you may use the following command to stop the container:

```
make stop-container-development
```