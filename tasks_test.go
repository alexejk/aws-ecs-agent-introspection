package ecsagent

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntrospector_GetTasks(t *testing.T) {

	ts, api := testServerAndClient(t, "tasks.json")
	defer ts.Close()

	tasks, err := api.GetTasks()

	assert.NoError(t, err)
	assert.Len(t, tasks.Tasks, 2)

	task := tasks.Tasks[0]
	assert.Equal(t, "arn:aws:ecs:us-east-1:123456789012:task/example5-58ff-46c9-ae05-543f8example", *task.Arn)
	assert.Equal(t, "8", *task.Version)
	assert.Equal(t, "RUNNING", *task.DesiredStatus)
	assert.Equal(t, "RUNNING", *task.KnownStatus)

	assert.Len(t, task.Containers, 2)
}

func TestIntrospector_GetTaskByTaskArn(t *testing.T) {

	ts, api := testServerAndClient(t, "task.json")
	defer ts.Close()

	taskArn := "arn:aws:ecs:us-east-1:123456789012:task/e01d58a8-151b-40e8-bc01-22647b9ecfec"
	task, err := api.GetTaskByTaskArn(taskArn)

	assert.NoError(t, err)
	assert.Equal(t, taskArn, *task.Arn)

}

func TestIntrospector_GetTaskByDockerID(t *testing.T) {

	ts, api := testServerAndClient(t, "task.json")
	defer ts.Close()

	dockerID := "79c796ed2a"
	task, err := api.GetTaskByDockerID(dockerID)

	assert.NoError(t, err)
	//assert.Equal(t, dockerId, *task.)

	found := false

	for _, c := range task.Containers {
		if strings.HasPrefix(*c.DockerID, dockerID) {
			found = true
			break
		}
	}

	assert.True(t, found, "required docker container should be found")

}
