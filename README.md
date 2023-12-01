# Description
> A simple repository with basic CRUD api and DDD pattern to serve as a showcase. 

## Getting started
### Project structure
- `/cmd` will have all the code that will be built and shipped as binaries. they are most likely to be named `main.go` and be found at `/cmd/<executable-name>/main.go`
- `/pkg` will have the actual shared codebase of the project. all of the modules used by the entrypoints will come from this folder
- `/scripts` will have useful utilitary scripts such as `dev.sh` that usually runs the project at a dev-local environment
- `/pkg/*<storage|repository>*/migrations` (optional) this one is optional and if found it will contain eventual migrations. We use the [migrate CLI](https://github.com/golang-migrate/migrate) to create and manage our migrations
- `/misc` (optional) will have miscellaneous scripts and files that do not relate to the project directly but are (or were) used to build and maintain the project. (such as a migration script e.g.)

