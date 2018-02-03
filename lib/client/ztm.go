package umwarsawclient

import (
	"time"

	"github.com/kaweue/api-um-warsaw-client/lib/types"
)

const (
	BusStopRequestId        = "b27f4c17-5c50-4a5b-89dd-236b282bc499"
	LinesOnBusStopRequestId = "88cd555f-6f31-43ca-9de4-66c479ad5942"
	TimeTableRequestId      = "e923fa0e-d96c-43f9-ae6e-60518c9f3238"
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

func (c *Client) GetTimeTable(busStopId string, busStopNr string, line string) (types.TimeTable, error) {

	query := struct {
		id        string
		busstopId string
		busstopNr string
		line      string
	}{
		TimeTableRequestId,
		busStopId,
		busStopNr,
		line,
	}

	result, err := c.executeQuery(query)
	if err != nil {
		return types.TimeTable{}, err
	}

	var timeTable types.TimeTable
	for _, res := range result.Result {
		var record types.TimeTableRecord
		for _, val := range res.Values {
			switch val.Key {
			case "brygada":
				record.Brigade = val.Value
			case "kierunek":
				record.Direction = val.Value
			case "czas":
				record.Time, err = time.Parse("15:04:05", val.Value)
				if err != nil {
					return types.TimeTable{}, err
				}
			}
		}
		timeTable.Record = append(timeTable.Record, record)
	}
	return timeTable, nil
}
