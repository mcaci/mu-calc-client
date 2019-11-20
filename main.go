package main

import "encoding/json"

import "log"

import "net/http"

import "bytes"

import "io/ioutil"

func main() {

	requestBody, err := json.Marshal(map[string]int{
		"a": 40,
		"b": 4,
	})

	if err != nil {
		log.Fatalln(err)
	}

	// docker exec -it add-cont ip addr show eth0
	addr := "http://172.17.0.2:8080/add"
	log.Println(addr)
	resp, err := http.Post(addr, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
