package gomarathon

import (
	"fmt"
	"net/url"
)

// RegisterCallbackURL register a new callback url
func (c *Client) RegisterCallbackURL(uri string) (*DefaultResponse, error) {

	var r DefaultResponse

	options := &RequestOptions{
		Path:   "eventSubscriptions",
		Method: "POST",
		Params: &Parameters{
			CallbackURL: url.QueryEscape(uri),
		},
	}
	err := c.request(options, &r)
	if &r != nil {
		if r.GetCode() == 201 {
			return &r, nil
		}
	}
	return nil, err
}

// GetEventSubscriptions gets all registered callback url
func (c *Client) GetEventSubscriptions() (*DefaultResponse, error) {
	var r DefaultResponse
	options := &RequestOptions{
		Path: "eventSubscriptions",
	}
	err := c.request(options, &r)
	if err != nil {
		return nil, err
	}
	if r.GetCode() != 200 {
		return nil, fmt.Errorf("request error")
	}
	return &r, nil
}

// DeleteCallbackURL delete a particular callback url
func (c *Client) DeleteCallbackURL(uri string) (*DefaultResponse, error) {
	var r DefaultResponse

	options := &RequestOptions{
		Path:   "eventSubscriptions",
		Method: "DELETE",
		Params: &Parameters{
			CallbackURL: url.QueryEscape(uri),
		},
	}
	err := c.request(options, &r)
	if &r != nil {
		if r.GetCode() == 204 {
			return &r, nil
		}
	}
	return nil, err
}
