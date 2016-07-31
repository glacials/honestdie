FROM golang:onbuild
EXPOSE 8080
RUN make run
