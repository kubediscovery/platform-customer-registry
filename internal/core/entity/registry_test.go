package entity

import (
	"testing"
	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCustomerRegistry_Validate(t *testing.T) {
	var cr CustomerRegistry
	t.Run("success", func(t *testing.T) {
		cr = CustomerRegistry{
			ProjectName: "project",
			Repository:  "repo",
			UserName:    "user",
			UserEmail:   "email@domian.com"}
	})

	require.NoError(t, cr.Validate())

}
