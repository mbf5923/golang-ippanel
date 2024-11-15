package golang_ippanel

import (
	"encoding/json"
	"fmt"
	golang_ippanel_http "github.com/mbf5923/golang-ippanel/pkg/http"
)

type Config struct {
	ApiKey string `json:"apikey"`
	Url    string
}

type Response struct {
	Status       string       `json:"status"`
	Code         int          `json:"code"`
	ErrorMessage string       `json:"error_message"`
	Data         ResponseData `json:"data"`
}

type ResponseData struct {
	MessageId uint64 `json:"message_id"`
}

type Ippanel struct {
	config Config
}

func New(config Config) Ippanel {
	if len(config.Url) == 0 {
		config.Url = "https://api2.ippanel.com/api/v1/"
	}
	return Ippanel{
		config: config,
	}
}

func (ip Ippanel) SendPattern(patternCode string, sendNumber string, recipientNumber string, patternVariables map[string]any) (Response, error) {
	patternReq := PatternRequest{
		Code:      patternCode,
		Sender:    sendNumber,
		Recipient: recipientNumber,
		Variable:  patternVariables,
	}
	var response Response
	post, err := golang_ippanel_http.Post(fmt.Sprintf("%s%s", ip.config.Url, "sms/pattern/normal/send"), ip.config.ApiKey, patternReq)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(post, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
