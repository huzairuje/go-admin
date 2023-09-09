# README

## Description
this repository is basically a boilerplate admin on golang using an echo framework (https://echo.labstack.com/) and postgresql.


## Install PreRequisite
1. run `make local` to install or up or run the postgres container and redis container.
2. make sure the docker postgres or container postgres is running and make a database on the postgres container with name `coba`
   and run the sql script on the directory `db/migrations` (this migration is still TODO)
3. make sure the config on config.local.yaml is correct

## How to Run Locally
```shell
make run
```
or simply with go command
```shell
go run main.go
```
