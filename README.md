# AWS ECS Agent Introspection Client

[![Build Status](https://travis-ci.org/alexejk/aws-ecs-agent-introspection.svg?branch=master)](https://travis-ci.org/alexejk/aws-ecs-agent-introspection)
[![GoDoc](https://godoc.org/github.com/alexejk/aws-ecs-agent-introspection?status.svg)](https://godoc.org/github.com/alexejk/aws-ecs-agent-introspection)
[![License](https://img.shields.io/github/license/alexejk/aws-ecs-agent-introspection.svg)](https://github.com/alexejk/aws-ecs-agent-introspection)

A simple Golang client for accessing ECS Introspection APIs.  
AWS Documentation: [ECS Container Agent Introspection](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-introspection.html) 

## Usage

Simply create a new `Introspector` instance and query the agent.
```go
api := ecsagent.NewIntrospector()
metadata, err := api.GetMetadata()
if err != nil {
  fmt.Printf("failed getting agent metadata: %v\n", err)
  return
}

fmt.Printf("Agent of instance %s is part of %s cluster", *metadata.ContainerInstanceArn, *metadata.Cluster)
```

## License

The Introspection client library for Amazon ECS Container Agent is licensed under the MIT License.
