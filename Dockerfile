FROM golang:1.12

WORKDIR /go/src/cat
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]
