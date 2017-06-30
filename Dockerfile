# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base
ADD . /src
RUN cd /src && make build

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/growser /app/
ENTRYPOINT ./growser
