package db

import (
	"context"
	"time"
)

type Team struct {
	ID            string `gorm:"primaryKey"`
	Name          string
	Slug          string
	CreationTime  time.Time
	MarkedDeleted bool
}

type ListOpts struct {
	Pagination

	OrderBy    string
	SearchTerm string
}

type TeamRepository interface {
	List(ctx context.Context, opts ListOpts) ([]*Team, error)

	Get(ctx context.Context, id string) (*Team, error)
	GetByUser(ctx context.Context, userID string) (*Team, error)

	Create(ctx context.Context, team Team) (*Team, error)

	Delete(ctx context.Context, id string) (*Team, error)
}
