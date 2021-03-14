# Golang DDD Boilerplate
---
It's a simple boilerplate for DDD based aplications

This boilerplate is using the following techs:

* HTTP: [Echo](https://github.com/labstack/echo)
* Database: Postgres (default driver)
* Dependency Injection: [Wire](https://github.com/google/wire)
* Mocks generation: [gomock](https://github.com/golang/mock)
> One advantage of splitting your application into layers is that you can change your edges without concern about the application core <3
**Change them as you need/want**  

## Running in locally

1. download and build dependencies

```shell
  $ make build
```
2. Provide a .env file as .env.example or export this environment variables

3. Run generated bin
```shell
  $ ./app
```

## Docker
1. Generate a docker image
```shell
  $ make build-docker
```

2. Run it
```shell
  $ docker run --env-file .env --net=host jeeo/go_api
```
> PS: I was lazy :) using host network and envfile is just for example purpose. 
Run it as you want following good pratices for your environment
---
## Todo
- [x] Define project layers
- [x] Adds deployment artifacts
- [ ] Unit tests example
> there are few examples of unit tests for mapper and application layer =)
