package services

import (
	"backend-template/utilities"
	"net/http"
	"strconv"
)

// PROCESS

func AuthUser(auth string) (int, error) {
	auth_url := "http://" + utilities.GoDotEnvVariable("AUTH_URL") + "/api/auth/user" //URL TO AUTH SERVICE

	//auth_url := "http:\\\\" + utilities.GoDotEnvVariable("AUTH_URL") + "\\api\\auth\\user" //URL TO AUTH SERVICE
	var response struct {
		Data struct {
			Sub string `json:"sub"`
		} `json:"data"`
	}
	//fmt.Println(auth_url)
	err := utilities.HTTPRequest(
		"GET",
		auth_url,
		http.Header{
			"Authorization": {auth},
			"Content-Type":  {"application/json"},
		},
		nil,
		&response,
	)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(response.Data.Sub)
}
