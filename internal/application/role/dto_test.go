package role

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleRequest_ToEntity(t *testing.T) {
	tests := []struct {
		name     string
		req      RoleRequest
		expected int8
	}{
		{
			name:     "Empty Status",
			req:      RoleRequest{Status: ""},
			expected: 1, // Default to 1
		},
		{
			name:     "Status String 0",
			req:      RoleRequest{Status: "0"},
			expected: 0,
		},
		{
			name:     "Status String 1",
			req:      RoleRequest{Status: "1"},
			expected: 1,
		},
		{
			name:     "Status Int 0",
			req:      RoleRequest{Status: 0},
			expected: 0,
		},
		{
			name:     "Status Int 1",
			req:      RoleRequest{Status: 1},
			expected: 1,
		},
		{
			name:     "Status Float 0",
			req:      RoleRequest{Status: float64(0)},
			expected: 0,
		},
		{
			name:     "Status Float 1",
			req:      RoleRequest{Status: float64(1)},
			expected: 1,
		},
		{
			name:     "Invalid Status Type",
			req:      RoleRequest{Status: true},
			expected: 1, // Default to 1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entity := tt.req.ToEntity()
			assert.Equal(t, tt.expected, entity.Status)
		})
	}
}
