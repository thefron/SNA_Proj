Social Network Analysis Term Proj
=================================

## Installation

Clone the repository under your `GOPATH`:

```
$ git clone git@github.com:thefron/SNA_Proj.git $GOPATH/src/github.com/thefron/SNA_Proj.git
```

Restore dependencies:

```
# If you dont have godep installed, install godep first.
$ go get -u github.com/tools/godep
# Restore dependencies
$ godep restore
```

## Development

### Dependency Management

We are using [godep](https://github.com/tools/godep) to manage external
dependencies.

When adding a new dependency, mgo(MongoDB driver) for instance:

```
$ go get gopkg.in/mgo.v2
# Make sure you include mgo.v2 in your source code
# goimports is encouraged to use.
$ godep save ./cmd/...
# mgo.v2 is copied to vendor directory.
```
