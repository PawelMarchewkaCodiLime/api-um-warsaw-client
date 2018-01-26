package main

import (
	"fmt"

	"github.com/kaweue/api-um-warsaw-client/lib/authenticator"
	"github.com/kaweue/api-um-warsaw-client/lib/client"
)

func main() {
	auth := authenticator.NewAuthenticator("bb9fea97-6310-49a2-baa0-b0fdd3ed44dd")
	client := umwarsawclient.NewClient("https://api.um.warszawa.pl/api/", auth, nil)

	busStop, _ := client.GetBusStop("znana")
	fmt.Printf("BusStopId = %d BusStopName = %s\n", busStop.BusID, busStop.Name)
}
