FROM golang:alpine AS build-container
ADD . /work
WORKDIR /work
ENV GO111MODULE=on
RUN apk --update add --no-cache git mercurial
RUN go build -o=migrate ./mysql/migrate.go
RUN go build

FROM alpine:3.9
COPY --from=build-container /work/cat-api-go /usr/local/bin/cat-api-go
COPY --from=build-container /work/migrate /usr/local/bin/migrate
CMD ["usr/local/bin/cat-api-go"]
