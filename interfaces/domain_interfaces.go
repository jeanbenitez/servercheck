package interfaces

import (
	"context"

	"github.com/jeanbenitez/servercheck/models"
)

// IDomainController interface...
type IDomainController interface {
	Fetch(ctx context.Context, num int64) ([]*models.Domain, error)
	GetByDomain(ctx context.Context, domain string) (*models.Domain, error)
	Create(ctx context.Context, d *models.Domain) (bool, error)
	Update(ctx context.Context, d *models.Domain) (*models.Domain, error)
	Delete(ctx context.Context, domain string) (bool, error)
}

// IServerController interface...
type IServerController interface {
	FetchByDomain(ctx context.Context, domain string) ([]*models.Server, error)
	Create(ctx context.Context, d *models.Server) (bool, error)
	Delete(ctx context.Context, domain string) (bool, error)
}
