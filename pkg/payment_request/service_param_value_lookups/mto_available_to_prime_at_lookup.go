package serviceparamvaluelookups

import (
	"database/sql"

	"github.com/transcom/mymove/pkg/appcontext"
	"github.com/transcom/mymove/pkg/apperror"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services/ghcrateengine"
)

// MTOAvailableToPrimeAtLookup does lookup on the MTOAvailableToPrime timestamp
type MTOAvailableToPrimeAtLookup struct {
}

func (m MTOAvailableToPrimeAtLookup) lookup(appCtx appcontext.AppContext, keyData *ServiceItemParamKeyData) (string, error) {
	db := appCtx.DB()

	// Get the MoveTaskOrder
	moveTaskOrderID := keyData.MoveTaskOrderID
	var moveTaskOrder models.Move
	err := db.Find(&moveTaskOrder, moveTaskOrderID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return "", apperror.NewNotFoundError(moveTaskOrderID, "looking for MoveTaskOrderID")
		default:
			return "", apperror.NewQueryError("Move", err, "")
		}
	}

	availableToPrimeAt := moveTaskOrder.AvailableToPrimeAt
	if availableToPrimeAt == nil {
		return "", apperror.NewBadDataError("This move task order is not available to prime")
	}

	return (*availableToPrimeAt).Format(ghcrateengine.TimestampParamFormat), nil
}
