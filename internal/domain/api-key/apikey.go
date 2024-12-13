package apikey

type APIKey struct {
	Key        string
	SystemName string
}

type Repository interface {
	GetSystemByKey(apiKey string) (*APIKey, error)
	CreateAPIKey(systemName string) (*APIKey, error)
}
