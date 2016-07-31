# honestdie

honestdie is a simple web app that picks a random element from a given group of elements, but does it on a page that
everyone can see in realtime from separate devices, with the same result witnessed by all.

## Running locally

### Prerequisites
* Go 1.6+

```bash
go get -u github.com/jteeuwen/go-bindata/...
go get -u github.com/elazarl/go-bindata-assetfs/...
```

### Building
```bash
make
```

### Running
```bash
make run
```

Access at `localhost:8080`.
