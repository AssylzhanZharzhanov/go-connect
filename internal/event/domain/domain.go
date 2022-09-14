package domain

// Events table name.
const (
	TableName = "events"
)

// EventID - represents the event id.
type EventID int

// Event - represents event
type Event struct {
	ID   EventID `json:"id"`
	Name string  `json:"name"`
}

// Validate - validates record struct
func (e Event) Validate() error {
	return nil
}

// TableName - returns name of table
func (e Event) TableName() string {
	return TableName
}
