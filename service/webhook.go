package services

import (
	"bytes"
	"go.uber.org/zap"
	"net/http"
	"tashilkar_health_checker/config"
)

type Service struct {
	config *config.Config
	logger *zap.SugaredLogger
}

func New(config *config.Config, logger *zap.SugaredLogger) *Service {
	service := &Service{
		config: config,
		logger: logger,
	}
	return service
}

func (s *Service) Alert(message string) {
	client := &http.Client{}
	jsonBody := []byte(`{"message":"` + message + `"}`)
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", s.config.WebhookUrl, bodyReader)
	if err != nil {
		s.logger.Errorf("service could not find. error is: %v", err)
	}
	req.Close = true
	req.Header.Set("Accept-Encoding", "identity")

	_, err = client.Do(req)
	if err != nil {
		//s.logger.Errorf("service could not send request. error is: %v", err)
	}
}
