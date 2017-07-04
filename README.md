# simple

This repository contains a simple little Golang HTTP server which I'm using to
experiment with hosting services on Amazon's ECS.

## Requirements

* Docker 17.05 or later
* `make`

### Optional

* Local go environment (we build inside the container so no need really)
* `glide` - this project uses glide to manage vendored dependencies, so if you
  want to add new dependencies it would be easiest to install glide locally

## Building the container

The project is set up to use a multistage build for Docker. This requires
Docker version 17.05 or later.

The project contains a Makefile with a couple of tasks for building and
deploying the container, so to build the container run the following:

```bash
$ make build
```

This command invokes `docker-compose` which then builds the image defined in
our Dockerfile. This means install required components inside the
`golang:alpine` build container, running the tests, and compiling the code
before copying the compiled binary into a minimal alpine container.

## Publishing the container

Our `docker-compose` file also defines the image on docker hub, so to publish a
new version run the following command:

```bash
$ make push
```

This command pushes our image over to Docker hub ready for use.

## Running the container

The container is set up to run using docker-compose, so run the following
command to start it running:

```bash
$ docker-compose up
```

## TODO

* Figure out how to deploy updated containers within ECS
* Add CI integration, and have the CI build server publish the final container
