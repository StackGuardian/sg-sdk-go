package stackguardian

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	assert.Equal(t, "https://api.app.stackguardian.io", config.BaseURL)
	assert.Equal(t, 3, config.MaxRetries)
	assert.NotNil(t, config.HTTPClient)
	assert.Contains(t, config.UserAgent, "sg-sdk-go/v")
}

func TestConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid config",
			config: &Config{
				APIKey:  "apikey test-key",
				BaseURL: "https://api.app.stackguardian.io",
			},
			wantErr: false,
		},
		{
			name: "missing API key",
			config: &Config{
				BaseURL: "https://api.app.stackguardian.io",
			},
			wantErr: true,
			errMsg:  "APIKey is required",
		},
		{
			name: "missing base URL",
			config: &Config{
				APIKey: "apikey test-key",
			},
			wantErr: true,
			errMsg:  "BaseURL is required",
		},
		{
			name:    "empty config",
			config:  &Config{},
			wantErr: true,
			errMsg:  "APIKey is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	t.Run("creates client with valid config", func(t *testing.T) {
		config := DefaultConfig()
		config.APIKey = "apikey test-key"

		client, err := NewClient(config)

		require.NoError(t, err)
		assert.NotNil(t, client)
		assert.NotNil(t, client.Organizations)
		assert.NotNil(t, client.Stacks)
		assert.NotNil(t, client.Workflows)
		assert.NotNil(t, client.WorkflowGroups)
	})

	t.Run("fails with invalid config", func(t *testing.T) {
		config := &Config{
			BaseURL: "https://api.app.stackguardian.io",
			// Missing APIKey
		}

		client, err := NewClient(config)

		require.Error(t, err)
		assert.Nil(t, client)
	})

	t.Run("uses default config if nil", func(t *testing.T) {
		// This would fail validation without APIKey
		client, err := NewClient(nil)

		require.Error(t, err) // Should fail validation
		assert.Nil(t, client)
	})
}

func TestClient_Config(t *testing.T) {
	config := DefaultConfig()
	config.APIKey = "apikey test-key"

	client, err := NewClient(config)
	require.NoError(t, err)

	gotConfig := client.Config()
	assert.Equal(t, config, gotConfig)
}

// Example of how to define minimal interfaces for testing
type OrganizationReader interface {
	ReadOrganization(ctx interface{}, org string) (interface{}, error)
}

// This demonstrates the pattern for mocking in tests
func TestExampleMocking(t *testing.T) {
	// In your application code, you would define a minimal interface:
	//
	// type OrganizationReader interface {
	//     ReadOrganization(ctx context.Context, org string) (*api.OrgGetResponse, error)
	// }
	//
	// Then your function accepts the interface:
	//
	// func MyFunc(ctx context.Context, client OrganizationReader) error {
	//     org, err := client.ReadOrganization(ctx, "my-org")
	//     // ...
	// }
	//
	// In tests, you can create a mock that implements only that interface:
	//
	// type mockOrgReader struct {
	//     mock.Mock
	// }
	//
	// func (m *mockOrgReader) ReadOrganization(ctx context.Context, org string) (*api.OrgGetResponse, error) {
	//     args := m.Called(ctx, org)
	//     return args.Get(0).(*api.OrgGetResponse), args.Error(1)
	// }

	t.Skip("This is a documentation example, not a real test")
}
