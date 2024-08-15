package graph

import (
	"context"
	"user-management/internal/db"
)

func (r *queryResolver) GetFact(ctx context.Context, id int32) (*db.Fact, error) {
	fact, err := r.GetFact(ctx, id)
	if err != nil {
		return nil, err
	}
	return fact, nil
}

func (r *queryResolver) ListFacts(ctx context.Context) ([]*db.Fact, error) {
	facts, err := r.ListFacts(ctx)
	if err != nil {
		return nil, err
	}
	return facts, nil
}
