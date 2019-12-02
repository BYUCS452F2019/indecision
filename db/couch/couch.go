package couch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Constants .
const (
	CHOICES     = "choices"
	USERS       = "users"
	RESTAURANTS = "restaurants"
	HISTORY     = "history"
)

// DB is the MongoDB proxy object that implements our DB interface
type DB struct {
}

// NewDB returns a pointer to the DB object
func NewDB() *DB {
	return &DB{}
}

// OpenConnection opens a connection with the database
func (c *DB) OpenConnection() error {

	return nil
}

// CloseConnection closes the connection to the database
func (c *DB) CloseConnection() error {
	return nil
}

func (c *DB) req(method, endpoint, contentType string, body []byte) (string, []byte, error) {
	errMsg := "unable to make request against couch"

	url := fmt.Sprintf("http://localhost:5984/%s", endpoint)
	url = strings.TrimSpace(url)

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return "", nil, fmt.Errorf("%s: %s", errMsg, err)
	}

	if len(contentType) > 0 {
		req.Header.Add("Content-Type", contentType)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", nil, fmt.Errorf("%s: %s", errMsg, err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil, fmt.Errorf("%s: %s", errMsg, err)
	}

	if resp.StatusCode/100 != 2 {
		ce := CouchError{}
		err = json.Unmarshal(b, &ce)
		if err != nil {
			return "", nil, fmt.Errorf("%s: received a non-200 response from %s. body: %s", errMsg, url, b)
		}

		return "", nil, CheckCouchErrors(ce)
	}

	return resp.Header.Get("content-type"), b, nil
}

// MakeRequest .
func (c *DB) MakeRequest(method, endpoint, contentType string, body []byte, toFill interface{}) error {
	respType, body, err := c.req(method, endpoint, contentType, body)
	if err != nil {
		return err
	}

	if !strings.EqualFold(respType, "application/json") {
		return fmt.Errorf("unexpected response content-type: expected %s, but got %s", "application/json", respType)
	}

	if toFill == nil {
		return nil
	}

	// unmarshal body
	err = json.Unmarshal(body, toFill)
	if err != nil {
		return fmt.Errorf("unable to make request against couch: %s", err)
	}

	return nil
}

// CheckCouchErrors gives more readable messages for the couch errors
func CheckCouchErrors(ce CouchError) error {
	switch strings.ToLower(ce.Error) {
	case "not_found":
		return &NotFound{fmt.Sprintf("The ID requested was unknown. Message: %v.", ce.Reason)}
	case "conflict":
		return &Conflict{fmt.Sprintf("There was a conflict updating/creating the document: %v", ce.Reason)}
	case "bad_request":
		return &BadRequest{fmt.Sprintf("The request was bad: %v", ce.Reason)}
	default:
		return fmt.Errorf("unknown error type: %v. Message: %v", ce.Error, ce.Reason)
	}
}

// CouchError is a struct for the error response from couch
type CouchError struct {
	Error  string `json:"error"`
	Reason string `json:"reason"`
}

// NotFound .
type NotFound struct {
	msg string
}

// Error gets the error message
func (n NotFound) Error() string {
	return n.msg
}

// Conflict .
type Conflict struct {
	msg string
}

// Error gets the error message
func (c Conflict) Error() string {
	return c.msg
}

// BadRequest .
type BadRequest struct {
	msg string
}

// Error gets the error message
func (br BadRequest) Error() string {
	return br.msg
}
