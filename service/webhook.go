package services

import (
	"bytes"
	"fmt"
	"net/http"
)

type Service struct {
}

func New() *Service {
	service := &Service{}
	return service
}

func (s *Service) Alert(message string) {
	client := &http.Client{}
	webhookUrl := "http://127.0.0.1:3000/api/v1/update"
	jsonMessage := `{"message":"` + message + `"}`

	fmt.Println(jsonMessage)
	jsonBody := []byte(jsonMessage)
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", webhookUrl, bodyReader)
	if err != nil {
		fmt.Println("service could not find", err)
	}
	req.Close = true
	req.Header.Set("Accept-Encoding", "identity")

	_, err = client.Do(req)
	if err != nil {
		fmt.Println("service could not find 2", err)
	}

	fmt.Println("inja mire")
}
