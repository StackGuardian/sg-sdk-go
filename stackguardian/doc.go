// Package stackguardian provides the StackGuardian SDK for Go.
//
// The StackGuardian SDK allows you to interact with the StackGuardian API
// for infrastructure orchestration, policy management, and workflow automation.
//
// # Getting Started
//
// To use the SDK, first create a Config and then create a Client:
//
//	config := stackguardian.DefaultConfig()
//	config.APIKey = "apikey " + os.Getenv("SG_API_TOKEN")
//
//	client, err := stackguardian.NewClient(config)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// # Making API Calls
//
// Use the service clients to make API calls:
//
//	ctx := context.Background()
//	org, err := client.Organizations.ReadOrganization(ctx, "my-org")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// # Error Handling
//
// The SDK provides typed errors with helper functions:
//
//	org, err := client.Organizations.ReadOrganization(ctx, "my-org")
//	if err != nil {
//	    if stackguardian.IsNotFoundError(err) {
//	        // Handle not found
//	    } else if stackguardian.IsUnauthorizedError(err) {
//	        // Handle unauthorized
//	    }
//	}
//
// # Context and Cancellation
//
// All API methods accept a context.Context for cancellation and timeouts:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//
//	org, err := client.Organizations.ReadOrganization(ctx, "my-org")
//
// # Retry Logic
//
// The SDK includes automatic retry with exponential backoff for transient errors.
// Configure retry behavior via Config:
//
//	config := stackguardian.DefaultConfig()
//	config.MaxRetries = 5
//	config.RetryWaitMin = 2 * time.Second
//	config.RetryWaitMax = 60 * time.Second
//
// # Testing
//
// For testing, define minimal interfaces for the operations you use:
//
//	type OrganizationReader interface {
//	    ReadOrganization(ctx context.Context, org string) (*api.OrgGetResponse, error)
//	}
//
//	func MyFunc(ctx context.Context, client OrganizationReader) error {
//	    org, err := client.ReadOrganization(ctx, "my-org")
//	    // ...
//	}
//
// This allows you to mock only the methods you need in tests.
package stackguardian

// Version is the current version of the SDK.
const Version = "2.0.0"
