package gomarathon

import (
	"fmt"
)

// ListTasks get all the tasks running on the cluster
func (c *Client) ListTasks() (*DefaultResponse, error) {
	var r DefaultResponse
	options := &RequestOptions{
		Path: "tasks",
	}
	err := c.request(options, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// GetAppTasks get the tasks across the cluster from the appid
func (c *Client) GetAppTasks(appID string) (*DefaultResponse, error) {
	var r DefaultResponse
	options := &RequestOptions{
		Path: fmt.Sprintf("apps/%s/tasks", appID),
	}
	err := c.request(options, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// KillTasks will kill the tasks from the host for the appid
func (c *Client) KillTasks(appID string, host string, scale bool) (*DefaultResponse, error) {
	var r DefaultResponse
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s/tasks", appID),
		Method: "DELETE",
		Params: &Parameters{
			Host:  host,
			Scale: scale,
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

// KillTask kill a particular taskid
func (c *Client) KillTask(appID string, taskID string, scale bool) (*DefaultResponse, error) {
	var r DefaultResponse
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s/tasks/%s", appID, taskID),
		Method: "DELETE",
		Params: &Parameters{
			Scale: scale,
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
