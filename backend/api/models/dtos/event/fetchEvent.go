package event

import "time"

type FetchEventDto struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location,omitempty"`
	Created  time.Time `json:"created"`
}
