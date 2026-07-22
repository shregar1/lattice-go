package queue
import "context"
type RabbitMQQueueDriver struct{}
func (r *RabbitMQQueueDriver) Publish(ctx context.Context, topic string, payload any) (string, error) { return "amqp_123", nil }
func (r *RabbitMQQueueDriver) Subscribe(ctx context.Context, topic string, handler func(msg any) error) error { return nil }
func (r *RabbitMQQueueDriver) GetDriverName() string { return "rabbitmq" }
