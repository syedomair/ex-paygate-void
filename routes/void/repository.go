package void

import (
	"github.com/syedomair/ex-paygate-lib/lib/models"
)

// Repository interface
type Repository interface {
	SetRequestID(requestID string)
	VoidApprove(inputApproveKey map[string]interface{}) (*models.Approve, error)
}
