# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base bash
ADD . /go/src/github.com/thingful/simple
WORKDIR /go/src/github.com/thingful/simple
RUN make test && make compile

# final stage
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build-env /go/src/github.com/thingful/simple /app/
ENTRYPOINT ./simple
EXPOSE 8080
