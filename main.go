package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	ga "github.com/ozgur-soft/google-analytics/src"
)

func main() {
	http.HandleFunc("/", view)
	server := http.Server{Addr: ":8080", ReadTimeout: 30 * time.Second, WriteTimeout: 30 * time.Second}
	server.ListenAndServe()
}

func view(w http.ResponseWriter, r *http.Request) {
	println("got request from", r.UserAgent())
	api := new(ga.API)
	api.UserAgent = r.UserAgent()
	api.ContentType = "application/x-www-form-urlencoded"

	client := new(ga.Client)
	client.ProtocolVersion = "1"
	client.ClientID = uuid.New().String()
	client.TrackingID = "UA-203067455-2"
	client.HitType = "event"
	client.DocumentLocationURL = "https://www.example.com/payment"
	client.DocumentTitle = "New product"
	client.DocumentEncoding = "UTF-8"

	client.EventAction = "Purchase"
	client.EventCategory = "Purchase"
	client.EventValue = "123"
	client.EventLabel = "Success"

	if r.URL.Path == "/" {
		println(api.Send(client))
		w.Write([]byte("Payment success"))
	}
}
