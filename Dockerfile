FROM golang:onbuild
EXPOSE 8080
RUN /bin/bash -c 'go get -u github.com/jteeuwen/go-bindata/... ;\
go get -u github.com/elazarl/go-bindata-assetfs/... ;\
go-bindata -pkg handler -o handler/bindata.go assets/...'
