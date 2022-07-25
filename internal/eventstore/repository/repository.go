package repository

import (
	"context"
)

//Repository pushes and filters events
type Repository interface {
	//Health checks if the connection to the storage is available
	Health(ctx context.Context) error
	// Push adds all events of the given aggregates to the event streams of the aggregates.
	// if unique constraints are pushed, they will be added to the unique table for checking unique constraint violations
	// This call is transaction save. The transaction will be rolled back if one event fails
	Push(ctx context.Context, events []*Event, uniqueConstraints ...*UniqueConstraint) error
	// Filter returns all events matching the given search query
	Filter(ctx context.Context, searchQuery *SearchQuery) (events []*Event, err error)
	//LatestSequence returns the latest sequence found by the search query
	LatestSequence(ctx context.Context, queryFactory *SearchQuery) (uint64, error)
	//InstanceIDs returns the instance ids found by the search query
	InstanceIDs(ctx context.Context, queryFactory *SearchQuery) ([]string, error)
	//CreateInstance creates a new sequence for the given instance
	CreateInstance(ctx context.Context, instanceID string) error
}
