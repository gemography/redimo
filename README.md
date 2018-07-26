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
`git clone https://github.com/hiddenfounders/redimo.git`

`cd redimo`

`go run app.go`

## Wiki
[redis for golang] https://godoc.org/github.com/go-redis/redis

[mongo for golang] https://github.com/globalsign/mgo
