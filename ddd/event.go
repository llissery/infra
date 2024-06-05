// Package ddd Domain-Driven Design definition
package ddd

type DomainEvent interface {
	EventType() string
	EventID() string
	OccurredOn() int64
}
