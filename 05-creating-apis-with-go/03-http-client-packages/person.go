package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// createPerson .
func createPerson(url, token string, person *Person) GeneralResponse{
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(person)
	if err != nil {
		log.Fatalf("Error at marshal person: %v", err)
	}
	
	resp := httpClient(http.MethodPost, url, token, data)
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body %v", err)
	}
	
	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Status code 201 not received, obtained: %d, \n\tresponse: %s", resp.StatusCode, body)
	}

	dataResponse := GeneralResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("Error at unmarshal body: %v", err)
	}

	return dataResponse
}