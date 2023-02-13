package event

type CreateEventDto struct {
	Name        string `json:"name" validate:"required, min=3, max=255" example:"Why Go is better than everything"`
	Description string `json:"description" validate:"required, min=3, max=255" example:"Sample description"`
	City        string `json:"city,omitempty" validate:"min=3, max=255"`
	Country     string `json:"country,omitempty" validate:"min=3, max=255"`
}
