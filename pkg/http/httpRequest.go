package golang_ippanel_http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Post(destinationUrl string, apiKey string, requestBody any) ([]byte, error) {
	postBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.NewRequest(http.MethodPost, destinationUrl, responseBody)
	resp.Header.Add("apikey", apiKey)
	resp.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(resp)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	//Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
