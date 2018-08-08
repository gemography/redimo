# redimo
[![Go Report Card](https://goreportcard.com/badge/github.com/hiddenfounders/redimo)](https://goreportcard.com/report/github.com/hiddenfounders/redimo)

An Open-Source Utility for Mongo-Redis Synchronization using Change Streams

## Description
redimo is a utility that synchronizes mongodb and redis using change streams. It can be used also in case of failure of mongodb instances.

## Structure
```
config
  config.go
handlers
  userHandler.go
models
  changeDocument.go
  dataStore.go
  user.go
utils
  mongo.go
  redis.go
app.go
```
To add new `handlers` of new collections. You should create your model under `models` and your handler under `handlers`. For more information about creating handlers for new collections please check the Wiki.

## Installation
### By Cloning the repository
`$ git clone https://github.com/hiddenfounders/redimo.git`

`$ cd redimo`

### By go get
`$ go get clone https://github.com/hiddenfounders/redimo.git`

Then go to your `$GOPATH/src/github.com/hiddenfounders/redimo`

### Configuration
The general config file is at `config/config.go`, if you do not use username and password, just let it empty.

Mongo configuration is at `utils/mongo.go`. If you are using a database with authentication, please make the boolean `withPassword = true` and add the credentials to `config/config.go`.

Redis configuration is at `utils/redis.go`.

### Run the application
`$ go run app.go`

`$ go build app.go && ./app`

## Development
If you want to add another handler, add the model to `models` and follow the template used in `userHandler.go` to make your own handler.

The function `ListenOnPipeline()` is generic and it can be called for any model.

## Wiki
[redis for golang] https://godoc.org/github.com/go-redis/redis

[mongo for golang] https://github.com/globalsign/mgo
