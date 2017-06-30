# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o growser

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/growser /app/
ENTRYPOINT ./growser
