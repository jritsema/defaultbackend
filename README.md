# defaultbackend

An HTTP server that returns 200 OK for health checks and 501 Not Implemented for all other requests.

Configuration via the following environment variables (all are optional):

- `PORT` (default: 8080) - the port to listen on
- `HEALTHCHECK` (default: /health) - the URL that should return 200 OK


## Development

```
 Choose a make command to run

  vet           vet code
  test          run unit tests
  build         build a binary
  autobuild     auto build when source files change
  dockerbuild   build project into a docker container image
  start         build and run local project
  deploy        build code into a container and deploy it to the cloud dev environment
```
