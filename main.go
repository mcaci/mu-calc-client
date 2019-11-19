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

	resp, err := http.Post("http://localhost:4000/add", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	// var (
	// 	httpCalcr = flag.String("http", ":8080", "http listen calcress")
	// )
	// flag.Parse()
	// ctx := context.Background()

	// srv := serv.NewFibonacciService()
	// errChan := make(chan error)

	// go func() {
	// 	c := make(chan os.Signal, 1)
	// 	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	// 	errChan <- fmt.Errorf("%s", <-c)
	// }()

	// // mapping endpoints
	// endpoints := serv.Endpoints{
	// 	FibonacciEndpoint: serv.MakeFibonacciEndpoint(srv),
	// }

	// // HTTP transport
	// go func() {
	// 	log.Println("fibonacci is listening on port:", *httpCalcr)
	// 	handler := serv.NewHTTPServer(ctx, endpoints)
	// 	errChan <- http.ListenAndServe(*httpCalcr, handler)
	// }()

	// log.Fatalln(<-errChan)
}
