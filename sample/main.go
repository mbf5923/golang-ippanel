package main

import (
	"github.com/mbf5923/golang-ippanel"
	"log"
)

func main() {
	conf := golang_ippanel.Config{
		ApiKey: "YOUR API KEY",
	}
	x := golang_ippanel.New(conf)
	resp, err := x.SendPattern("YOUR PATTERN CODE", "+9810005923", "+989333333333", map[string]any{
		"code": 123654,
	})
	if err != nil {
		log.Fatalln(err)
	}
	if resp.Code == 200 {
		log.Println(resp.Data.MessageId)
	} else {
		log.Println(resp.ErrorMessage)
	}

}
