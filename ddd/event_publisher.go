package ddd

import (
	"context"
)

type DomainEventPublisher interface {
	// Publish 实现领域事件发布.
	// 消息发布异常会抛出 error.
	Publish(ctx context.Context, aggregateType, aggregateID string, events ...DomainEvent)
}
