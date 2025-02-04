package adminuser

import (
	"github.com/gobuffalo/validate/v3"

	"github.com/transcom/mymove/pkg/appcontext"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services"
)

type adminUserQueryBuilder interface {
	FetchOne(appCtx appcontext.AppContext, model interface{}, filters []services.QueryFilter) error
	CreateOne(appCtx appcontext.AppContext, model interface{}) (*validate.Errors, error)
	UpdateOne(appCtx appcontext.AppContext, model interface{}, eTag *string) (*validate.Errors, error)
}

type adminUserFetcher struct {
	builder adminUserQueryBuilder
}

// FetchAdminUser fetches an admin user given a slice of filters
func (o *adminUserFetcher) FetchAdminUser(appCtx appcontext.AppContext, filters []services.QueryFilter) (models.AdminUser, error) {
	var adminUser models.AdminUser
	error := o.builder.FetchOne(appCtx, &adminUser, filters)
	return adminUser, error
}

// NewAdminUserFetcher return an implementation of the AdminUserFetcher interface
func NewAdminUserFetcher(builder adminUserQueryBuilder) services.AdminUserFetcher {
	return &adminUserFetcher{builder}
}
