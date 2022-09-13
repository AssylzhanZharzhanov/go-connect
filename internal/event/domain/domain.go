package domain

const (
	TableName = "events"
)

type EventID int

type Event struct {
	ID   EventID `json:"id"`
	Name string  `json:"name"`
}

func (e Event) Validate() error {
	return nil
}

func (e Event) TableName() string {
	return TableName
}
