package registered_services

type CreateRequest struct {
	Name        string   `json:"name" validate:"required,min=3,max=128"`
	Description string   `json:"description"`
	HealthUrl   string   `json:"health_url" validate:"required,url"`
	OwnerEmail  string   `json:"owner_email" validate:"required,email"`
	Tags        []string `json:"tags"`
}
