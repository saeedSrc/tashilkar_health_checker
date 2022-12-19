package Services

import (
	"fmt"
	"net/http"
)

func alert() {
	client := &http.Client{}
	webhookUrl := "http://127.0.0.1:3000/api/v1/update"

	req, err := http.NewRequest("GET", webhookUrl, http.NoBody)
	if err != nil {
		fmt.Println(err)
	}
	req.Close = true
	req.Header.Set("Accept-Encoding", "identity")

	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
}
