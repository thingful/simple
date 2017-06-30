# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base
ADD . /src
RUN cd /src && make compile

# final stage
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build-env /src/growser /app/
ENTRYPOINT ./growser
