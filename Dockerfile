FROM golang:1.17.3-alpine3.14

COPY ./go.* ./*.go /build/
COPY ./cmd /build/cmd
RUN cd /build && go build -o envinfo ./cmd/envinfo/main.go

CMD "/build/envinfo"
