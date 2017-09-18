package ecsagent

// AgentMetadata contains information about cluster, container instance and version of the ECS Agent
type AgentMetadata struct {
	Cluster              *string
	ContainerInstanceArn *string
	Version              *string
}

// GetMetadata returns information about ECS Agent running on the system
func (i *Introspector) GetMetadata() (*AgentMetadata, error) {

	metadata := &AgentMetadata{}
	if err := i.get("v1/metadata", metadata); err != nil {
		return nil, err
	}

	return metadata, nil
}
