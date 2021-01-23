FROM golang:1.15-alpine3.12 as build-env
RUN apk add build-base linux-headers
RUN mkdir /build
WORKDIR /build
ADD . /build
RUN go mod download
RUN go build

FROM alpine:3.12
COPY --from=build-env /build/fingy /bin/fingy
ENTRYPOINT [ "/bin/fingy" ]