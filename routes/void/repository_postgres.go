package void

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/syedomair/ex-paygate-lib/lib/models"
	"github.com/syedomair/ex-paygate-lib/lib/tools/logger"
)

type postgresRepo struct {
	client    *gorm.DB
	logger    logger.Logger
	requestID string
}

// NewPostgresRepository Public.
func NewPostgresRepository(c *gorm.DB, logger logger.Logger) Repository {
	return &postgresRepo{client: c, logger: logger, requestID: ""}
}

func (p *postgresRepo) SetRequestID(requestID string) {
	p.requestID = requestID
}

// VoidApprove Public
func (p *postgresRepo) VoidApprove(inputApproveKey map[string]interface{}) (*models.Approve, error) {
	methodName := "VoidApprove"
	p.logger.Debug(p.requestID, "M:%v start", methodName)
	start := time.Now()

	approveKey := ""
	if approveKeyValue, ok := inputApproveKey[ApproveKey]; ok {
		approveKey = approveKeyValue.(string)
	}

	approveObj := models.Approve{}
	if err := p.client.Table("approve").
		Where("approve_key = ?", approveKey).
		Find(&approveObj).Error; err != nil {
		return nil, errors.New("invalid approve_key")
	}

	inputApproveKey["status"] = 0

	if err := p.client.
		Table("approve").
		Where("approve_key = ?", approveKey).
		Updates(inputApproveKey).Error; err != nil {
		return nil, err
	}

	p.logger.Debug(p.requestID, "M:%v ts %+v", methodName, time.Since(start))
	return &approveObj, nil
}
