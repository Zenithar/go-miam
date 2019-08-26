package repositories

import (
	"context"

	"go.zenithar.org/miam/internal/models"
)

// ApplicationCreator describes application creator contract.
type ApplicationCreator interface {
	Create(ctx context.Context, entity *models.Application) error
}
