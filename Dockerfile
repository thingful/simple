# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base
ADD . /go/src/github.com/thingful/growser
WORKDIR /go/src/github.com/thingful/growser
RUN make test && make compile

# final stage
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build-env /go/src/github.com/thingful/growser /app/
ENTRYPOINT ./growser
EXPOSE 8000
