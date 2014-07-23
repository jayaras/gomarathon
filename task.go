package gomarathon

import (
	//	"encoding/json"
	"fmt"
)

func (c *Client) ListTasks() (*Response, error) {
	options := &RequestOptions{
		Path: "tasks",
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) GetAppTasks(appID string) (*Response, error) {
	options := &RequestOptions{
		Path: fmt.Sprintf("apps/%s/tasks", appID),
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) KillTasks(appID string, host string, scale bool) (*Response, error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s/tasks", appID),
		Method: "DELETE",
		Params: &Parameters{
			Host:  host,
			Scale: scale,
		},
	}
	r, err := c.request(options)
	if r != nil {
		if r.Code == 204 {
			return r, nil
		}
	}
	return nil, err
}

func (c *Client) KillTask(appID string, taskID string, scale bool) (*Response, error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s/tasks/%s", appID, taskID),
		Method: "DELETE",
		Params: &Parameters{
			Scale: scale,
		},
	}
	r, err := c.request(options)
	if r != nil {
		if r.Code == 204 {
			return r, nil
		}
	}
	return nil, err
}
