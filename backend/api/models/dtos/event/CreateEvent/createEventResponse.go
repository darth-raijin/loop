package createEventDto

type CreateEventResponse struct {
	ID          uuid   `json:"id""`
	Name        string `json:"name""`
	Description string `json:"description"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
}
