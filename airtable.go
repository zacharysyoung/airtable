package airtable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	apiUrl = "https://api.airtable.com/v0"
)

type Airtable struct {
	apiKey string `json:"apiKey"`
	base   string `json:"base"`
}

func New(apiKey, base string) *Airtable {
	return &Airtable{
		apiKey: apiKey,
		base:   base,
	}
}

type Table struct {
	Name       string `json:"name"`       // table name
	MaxRecords string `json:"maxRecords"` // max 100
	View       string `json:"view"`       // Grid view
}

func (a *Airtable) List(table Table, response interface{}) error {
	if err := a.call(GET, table, nil, nil, response); err != nil {
		return err
	}

	return nil
}

func (a *Airtable) Get(table Table, id string, response interface{}) error {
	if err := a.call(GET, table, &id, nil, response); err != nil {
		return err
	}

	return nil
}
func (a *Airtable) Create(table Table, data []byte, response interface{}) error {
	if err := a.call(POST, table, nil, data, response); err != nil {
		return err
	}

	return nil
}

func (a *Airtable) Update(table Table, id string, data []byte, response interface{}) error {
	if err := a.call(PATCH, table, &id, data, response); err != nil {
		return err
	}

	return nil
}

func (a *Airtable) Delete(table Table, id string) error {
	if err := a.call(DELETE, table, &id, nil, nil); err != nil {
		return err
	}

	return nil
}

type methodHttp string

const (
	GET    methodHttp = "GET"
	POST   methodHttp = "POST"
	PUT    methodHttp = "PUT"
	PATCH  methodHttp = "PATCH"
	DELETE methodHttp = "DELETE"
)

func (a *Airtable) call(method methodHttp, table Table, id *string, payload []byte, response interface{}) error {

	if table.MaxRecords == "" {
		table.MaxRecords = "100"
	}

	if table.View == "" {
		table.View = "Grid view"
	}

	table.View = url.QueryEscape(table.View)
	table.Name = url.QueryEscape(table.Name)

	var path string

	// list
	if method == GET && id == nil {
		path = fmt.Sprintf("%s/%s/%s?maxRecords=%s&view=%s", apiUrl, a.base, table.Name, table.MaxRecords, table.View)
	}

	// get || delete || update
	if (method == GET && id != nil) || (method == DELETE && id != nil || (method == PUT && id != nil || method == PATCH && id != nil)) {
		path = fmt.Sprintf("%s/%s/%s/%s", apiUrl, a.base, table.Name, *id)
	}

	// create
	if method == POST {
		path = fmt.Sprintf("%s/%s/%s", apiUrl, a.base, table.Name)
	}

	req, err := http.NewRequest(string(method), path, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.apiKey))
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{
		Timeout: time.Second * 10,
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	log.Println(method, path, res.StatusCode)

	if res.StatusCode == http.StatusUnprocessableEntity {
		return fmt.Errorf("%v Unprocessable Entity", res.StatusCode)
	}

	if res.StatusCode == http.StatusTooManyRequests {
		time.Sleep(time.Second * 5)
		return a.call(method, table, id, payload, response)
	}

	if method == DELETE {
		return nil
	}

	return json.NewDecoder(res.Body).Decode(response)
}
