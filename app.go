package gomarathon

import (
	"fmt"
)

/*
* Application requests
 */

// ListApps is listing all apps
func (c *Client) ListApps() (*DefaultResponse, error) {
	return c.ListAppsByCmd("")
}

// ListAppsByCmd list all apps by cmd filter
func (c *Client) ListAppsByCmd(filter string) (*DefaultResponse, error) {
	var r DefaultResponse

	options := &RequestOptions{
		Path: fmt.Sprintf("apps"),
		Params: &Parameters{
			Cmd: filter,
		},
	}
	err := c.request(options, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// ListAppVersions list app versions from a define appid
func (c *Client) ListAppVersions(appID string) (*DefaultResponse, error) {
	var r DefaultResponse

	options := &RequestOptions{
		Path: fmt.Sprintf("apps/%s/versions", appID),
	}
	err := c.request(options, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// GetApp gets a single app with production version
func (c *Client) GetApp(appID string) (*DefaultResponse, error) {
	var r DefaultResponse

	options := &RequestOptions{
		Path: fmt.Sprintf("apps/%s", appID),
	}
	err := c.request(options, &r)
	if err != nil {
		return nil, err
	}
	if r.GetCode() != 200 {
		return nil, fmt.Errorf("request rrror")
	}
	return &r, nil
}

// GetAppVersion get a single version from a single app
func (c *Client) GetAppVersion(appID string, version string) (*DefaultResponse, error) {
	var r DefaultResponse
	options := &RequestOptions{
		Path: fmt.Sprintf("apps/%s/versions/%s", appID, version),
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

// CreateApp Create a new Application
func (c *Client) CreateApp(app *Application) (*DefaultResponse, error) {
	var r DefaultResponse

	// TODO : VALIDATE DATAS
	options := &RequestOptions{
		Path:   "apps",
		Datas:  app,
		Method: "POST",
	}
	err := c.request(options, &r)
	if &r != nil {
		if r.GetCode() == 201 {
			return &r, nil
		}
	}
	return nil, err
}

// UpdateApp update the app but thoses changes are made for the next running app and does
// not shut down the production applications
func (c *Client) UpdateApp(appID string, app *Application) (*DefaultResponse, error) {
	var r DefaultResponse

	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s", appID),
		Datas:  app,
		Method: "PUT",
	}
	err := c.request(options, &r)
	if &r != nil {
		if r.GetCode() == 204 {
			return &r, nil
		}
	}
	return nil, err

}

// DeleteApp delete this app from the cluster
func (c *Client) DeleteApp(appID string) (*DefaultResponse, error) {
	var r DefaultResponse

	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s", appID),
		Method: "DELETE",
	}
	err := c.request(options, &r)
	if &r != nil {
		if r.GetCode() == 204 {
			return &r, nil
		}
	}
	return nil, err
}
