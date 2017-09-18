package ecsagent

import "fmt"

// AgentTasks encapsulates ECS Agent Tasks running on the system
type AgentTasks struct {
	Tasks []*AgentTask
}

// AgentTask encapsulates Tasks and their docker containers on the system
type AgentTask struct {
	Arn           *string
	DesiredStatus *string
	KnownStatus   *string
	Family        *string
	Version       *string
	Containers    []*AgentContainer
}

// AgentContainer encapsulates information about docker container
type AgentContainer struct {
	DockerID   *string `json:"DockerId"`
	DockerName *string
	Name       *string
}

const (
	taskFilterTaskArn  = "taskarn"
	taskFilterDockerID = "dockerid"
)

// GetTasks allows querying agent for all tasks it's aware of
func (i *Introspector) GetTasks() (*AgentTasks, error) {

	tasks := &AgentTasks{}
	if err := i.get("v1/tasks", tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetTaskByTaskArn returns information about a specific task on the system identified by a Task ARN
func (i *Introspector) GetTaskByTaskArn(taskArn string) (*AgentTask, error) {

	return i.getTaskBy(taskFilterTaskArn, taskArn)
}

// GetTaskByDockerID returns information about task associated with a provided Docker container ID
func (i *Introspector) GetTaskByDockerID(dockerID string) (*AgentTask, error) {

	return i.getTaskBy(taskFilterDockerID, dockerID)
}

func (i *Introspector) getTaskBy(param, value string) (*AgentTask, error) {

	task := &AgentTask{}
	url := fmt.Sprintf("v1/tasks?%s=%s", param, value)

	if err := i.get(url, task); err != nil {
		return nil, err
	}

	return task, nil
}
