package umwarsawclient

import (
	"github.com/kaweue/api-um-warsaw-client/lib/types"
)

const (
	BusStopRequestId        = "b27f4c17-5c50-4a5b-89dd-236b282bc499"
	LinesOnBusStopRequestId = "88cd555f-6f31-43ca-9de4-66c479ad5942"
)

func (c *Client) GetBusStop(busStopName string) (types.BusStop, error) {

	query := struct {
		id   string
		name string
	}{
		BusStopRequestId,
		busStopName,
	}

	result, err := c.executeQuery(query)
	if err != nil {
		return types.BusStop{}, err
	}
	busStop := types.BusStop{}
	busStop.BusID = result.Result[0].Values[0].Value
	busStop.Name = result.Result[0].Values[1].Value

	return busStop, nil
}

func (c *Client) GetLinesOnBusStop(busStopId string, busStopNr string) ([]types.Line, error) {

	query := struct {
		id        string
		busstopId string
		busstopNr string
	}{
		LinesOnBusStopRequestId,
		busStopId,
		busStopNr,
	}

	result, err := c.executeQuery(query)
	if err != nil {
		return nil, err
	}

	var lines []types.Line
	for _, res := range result.Result {
		for _, val := range res.Values {
			if val.Key == "linia" {
				lines = append(lines, types.Line(val.Value))
			}
		}
	}
	return lines, nil
}
