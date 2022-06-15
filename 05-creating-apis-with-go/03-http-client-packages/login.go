package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// loginClient. 
func loginClient(url, email, password string) LoginResponse {
	login := Login{
		Email:    email,
		Password: password,
	}

	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(&login)
	if err != nil {
		log.Fatalf("Error at marshal login: %v", err)
	}
	
	resp := httpClient(http.MethodPost, url, "", data)
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body %v", err)
	}
	
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Status code 200 not received, obtained: %d, \n\tresponse: %s", resp.StatusCode, body)
	}
	
	dataResponse := LoginResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("Error at unmarshal body: %v", err)
	}

	return dataResponse
}