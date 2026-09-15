package internal

// BaseService provides common functionality for all service clients.
type BaseService struct {
	Client *HTTPClient
}

// NewBaseService creates a new base service.
func NewBaseService(client *HTTPClient) *BaseService {
	return &BaseService{
		Client: client,
	}
}
