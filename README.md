# vizimind

## Description

Vizimind is the monorepo for the core api and the dashboard.

## Prerequisites

- [docker](https://docs.docker.com/engine/install/)
- [docker-compose](https://docs.docker.com/compose/install/)
- [buildx](https://github.com/docker/buildx)
- [golang](https://go.dev/doc/install)
- [nodejs](https://nodejs.org/en)
- [npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm)
- [task](https://taskfile.dev/installation/)

## Core

The core is developed with golang + grpc_gateway and use mongo db to store data.

## Dashboard

The dashboard is developed with solidjs + bootstrap with no typescript.

## How to run

- copy the `core/config.dist.yaml` to `core/config.yaml` and fill the missing values.
- copy the `app/dashboard/.env.dist` to `app/dashboard/.env` and fill the missing values.
- run task core:run to start the core api.
- run task dashboard:run to start the backoffice.

## Build

- run task core:build to build the core api.
- run task dashboard:build to build the dashboard.

## Test

- run task core:test to test the core api.
- run task dashboard:test to test the dashboard.

## Lint

- run task core:lint to lint the core api.
- run task dashboard:lint to lint the dashboard.

## Generate

- run task core:gen to generate the core api.

## Docker

- run TAG=latest task core:docker latest to build the core api docker image.
- run task dashboard:docker latest to build the dashboard docker image.

## Deploy

To deploy into staging or production build the docker images and push them to the registry.
Then go into the desired server and do `docker compose pull && docker compose restart` at the root directory.

## Project Recommendations

TODO
