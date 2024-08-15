package graph

import (
	"context"
)

func (r *mutationResolver) DeleteFact(ctx context.Context, id int32) (bool, error) {
	_, err := r.DeleteFact(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
