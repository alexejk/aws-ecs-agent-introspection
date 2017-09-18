package ecsagent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntrospector_GetMetadata(t *testing.T) {

	ts, api := testServerAndClient(t, "metadata.json")
	defer ts.Close()

	metadata, err := api.GetMetadata()

	assert.NoError(t, err)
	assert.Equal(t, "default", *metadata.Cluster)
	assert.Equal(t, "Amazon ECS Agent - v1.14.4 (f94beb4)", *metadata.Version)
}

func TestIntrospector_GetMetadataFailure(t *testing.T) {

	ts, api := failingServerAndClient(t, 500, "expected failure")
	defer ts.Close()

	metadata, err := api.GetMetadata()

	assert.Error(t, err)
	assert.Nil(t, metadata)
}
