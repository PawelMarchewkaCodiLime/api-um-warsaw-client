package umwarsawclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/kaweue/api-um-warsaw-client/lib/authenticator"
	"github.com/kaweue/api-um-warsaw-client/lib/types"
)

type Client struct {
	httpClient    *http.Client
	authenticator *authenticator.Authenticator
	apiURL        string
}

func NewClient(apiURL string, auth *authenticator.Authenticator, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		httpClient,
		auth,
		apiURL}
}

func parseQuery(values *url.Values, query interface{}) error {
	val := reflect.Indirect(reflect.ValueOf(query))
	for i := 0; i < val.Type().NumField(); i++ {
		values.Add(val.Type().Field(i).Name, val.Field(i).String())
	}
	return nil
}

func (c *Client) executeQuery(query interface{}) (*types.Result, error) {
	request, err := http.NewRequest(http.MethodGet, c.apiURL+"/action/dbtimetable_get", nil)
	if err != nil {
		return nil, err
	}

	q := request.URL.Query()
	err = parseQuery(&q, query)
	if err != nil {
		return nil, err
	}

	c.authenticator.Authenticate(&q)
	request.URL.RawQuery = q.Encode()
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Http status = " + response.Status)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result types.Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return &result, nil
}
