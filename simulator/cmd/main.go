package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-faker/faker/v4"
)

const webhookUrl string = "http://web/salesforce/webhook"
const numCustomers int = 10

type CustomerData struct {
	AccountManager  string `faker:"name"`
	Id              string `faker:"uuid_hyphenated"`
	Name            string `faker:"domain_name"`
	SalesEngineerId string `faker:"uuid_hyphenated"`
	SignedTos       bool
	SKU1            bool
	SKU2            bool
	SKU3            bool
}

func main() {
	fmt.Println("salesforce random data producer...start")

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)

	customers := make([]CustomerData, numCustomers)
	for i := 0; i < numCustomers; i++ {
		customers[i] = CustomerData{}
		err := faker.FakeData(&customers[i])
		if err != nil {
			fmt.Println("faker.FakeData.fail.(%s)", err)
		}
	}
	rand.Seed(time.Now().UnixNano())

	go func() {
		// 1. POST randomly form message to webhook_url
		customerId := rand.Intn(numCustomers)

		jsonBody, err := json.Marshal(customers[customerId])
		if err != nil {
			fmt.Println("json.Marshal.fail.(%s)", err)
		}
		bodyReader := bytes.NewReader(jsonBody)

		req, err := http.NewRequest(http.MethodPost, webhookUrl, bodyReader)
		if err != nil {
			fmt.Println("http.newrequest.fail.(%s)", err)
			// Ignore on error because the url is not around
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("http.DefaultClient.Do.fail.(%s)", err)
			// Ignore on error because the url is not around
		}

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("ioutil.ReadAll.fail.(%s)", err)
			// Ignore on error because the url is not around
		}
		fmt.Println("http.DefaultClient.Do.(%d):(%s)", resp.StatusCode, respBody)

		// 2. Sleep for 10 seconds
		time.Sleep(10 * time.Second)
	}()

	// Block until we have a signal
	<-cancelChan

	fmt.Println("salesforce random data producer...end")
}
