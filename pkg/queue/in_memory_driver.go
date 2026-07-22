package queue
import "context"
type InMemoryQueueDriver struct{}
func (i *InMemoryQueueDriver) Publish(ctx context.Context, topic string, payload any) (string, error) { return "msg_123", nil }
func (i *InMemoryQueueDriver) Subscribe(ctx context.Context, topic string, handler func(msg any) error) error { return nil }
func (i *InMemoryQueueDriver) GetDriverName() string { return "in_memory" }
