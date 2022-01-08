package void

import (
	"github.com/syedomair/ex-paygate-lib/lib/models"
)

// Payment Interface
type Payment interface {
	VoidPayment(*models.Approve) (string, error)
}
