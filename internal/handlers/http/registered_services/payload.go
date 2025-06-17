package registered_services

type CreateRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	HealthUrl   string   `json:"health_url"`
	OwnerEmail  string   `json:"owner_email"`
	Tags        []string `json:"tags"`
}
