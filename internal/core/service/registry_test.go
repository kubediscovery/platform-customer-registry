package service_test

import (
	"github.com/kubediscovery/platform-customer-registry/internal/core/entity"
	"github.com/kubediscovery/platform-customer-registry/internal/core/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRegistryService_Create(t *testing.T) {
	var cr entity.CustomerRegistry
	rs := service.NewRegistryService(nil)
	t.Run("success", func(t *testing.T) {
		cr = entity.CustomerRegistry{
			ProjectName: "project",
			Repository:  "repo",
			UserName:    "user",
			UserEmail:   "user@domain.com"}
	})
	crr, err := rs.Create(cr)
	require.NoError(t, err)
	assert.NotNil(t, crr)

}
