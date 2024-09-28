package service_test

import (
	"context"
	"testing"

	"github.com/kubediscovery/platform-customer-registry/internal/core/entity"
	"github.com/kubediscovery/platform-customer-registry/internal/core/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegistryService_Create(t *testing.T) {
	var cr entity.CustomerRegistry
	rs, _ := service.NewRegistryService(nil, nil, nil)
	t.Run("success", func(t *testing.T) {
		cr = entity.CustomerRegistry{
			ProjectName: "project",
			Repository:  "repo",
			UserName:    "user",
			UserEmail:   "user@domain.com"}
	})
	crr, err := rs.Create(context.Background(), &entity.CustomerRegistryResponse{CustomerRegistry: cr})
	require.NoError(t, err)
	assert.NotNil(t, crr)

}
