package queue
import "context"
type BaseQueueClient interface {
	Publish(ctx context.Context, topic string, payload any) (string, error)
	Subscribe(ctx context.Context, topic string, handler func(msg any) error) error
	GetDriverName() string
}
