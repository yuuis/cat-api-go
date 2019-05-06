FROM golang:alpine AS build-env
ADD . /work
WORKDIR /work
ENV GO111MODULE=on
RUN apk --update add --no-cache git mercurial
RUN go build

FROM alpine:3.9
COPY --from=build-env /work/cat-api-go /usr/local/bin/cat-api-go
CMD ["usr/local/bin/cat-api-go"]
