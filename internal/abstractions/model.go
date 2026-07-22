package abstractions

import "time"

type IModel interface {
	GetID() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}
