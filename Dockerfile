FROM golang:1.19-buster

RUN go version
ENV GOPATH=/

COPY ./go.mod ./

RUN go mod download

COPY ./ ./

RUN go build -o test-app ./main.go

CMD ["./test-app"]