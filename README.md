# AWS ECS Agent Introspection Client
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
