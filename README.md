# defaultbackend

An HTTP server that returns 200 OK for health checks and 501 Not Implemented for all other requests.

Configuration via the following environment variables (all are optional):

- `PORT` (default: 8080) - the port to listen on
- `HEALTHCHECK` (default: /health) - the URL that should return 200 OK

A pre-built image is published here https://gallery.ecr.aws/jritsema/defaultbackend

## Usage

```sh
docker run -it --rm public.ecr.aws/jritsema/defaultbackend
```


## Development

```
 Choose a make command to run

  vet           vet code
  test          run unit tests
  build         build a binary
  autobuild     auto build when source files change
  start         build and run local project
  dockerbuild   build project into a container image
  login         login to ecr public
  push          push container image to registry
```
