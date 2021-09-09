package v1

import (
	"os"
	"whatsapp-clone/logger"

	"go.uber.org/zap"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SetAPMSampleRate(rate string) {
	logger.Client().Debug("SetAPMSampleRate", zap.String("rate", rate))
	os.Setenv("ELASTIC_APM_TRANSACTION_SAMPLE_RATE", rate)
}
