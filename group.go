package gomarathon

import (
	"fmt"
)

/*
* Group support
 */

// ListApps is listing all apps
func (c *Client) ListGroups() (*Response, error) {
	return c.GetGroup("/")
}


// ListGroupVersions list app versions from a define appid
func (c *Client) ListGroupVersions(groupID string) (*Response, error) {
	options := &RequestOptions{
		Path: fmt.Sprintf("groups/%s/versions", groupID),
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetGroup gets a group
func (c *Client) GetGroup(groupID string) (*Response, error) {
	options := &RequestOptions{
		Path: fmt.Sprintf("groups/%s", groupID),
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return nil, fmt.Errorf("request rrror")
	}
	return r, nil
}

// GetAppVersion get a single version from a single app
func (c *Client) GetGroupVersion(groupID string, version string) (*Response, error) {
	options := &RequestOptions{
		Path: fmt.Sprintf("groups/%s/versions/%s", groupID, version),
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return nil, fmt.Errorf("request error")
	}
	return r, nil
}

// CreateGroup Create a new group
func (c *Client) CreateGroup(group *Group) (*Response, error) {
	// TODO : VALIDATE DATAS
	options := &RequestOptions{
		Path:   "groups",
		Datas:  group,
		Method: "POST",
	}
	r, err := c.request(options)
	if r != nil {
		if r.Code == 201 {
			return r, nil
		}
	}
	return nil, err
}


func (c *Client) UpdateGroup(groupID string, group *Group) (*Response, error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("groups/%s", groupID),
		Datas:  group,
		Method: "PUT",
	}
	r, err := c.request(options)
	if r != nil {
		if r.Code == 204 {
			return r, nil
		}
	}
	return nil, err

}

// DeleteGroup delete this app from the cluster
func (c *Client) DeleteGroup(groupID string) (*Response, error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("groups/%s", groupID),
		Method: "DELETE",
	}
	r, err := c.request(options)
	if r != nil {
		if r.Code == 204 {
			return r, nil
		}
	}
	return nil, err
}

// DeleteApp delete this app from the cluster
func (c *Client) RollbackGroup(groupID string, version string) (*Response, error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("groups/%s/version/%s", groupID,version),
		Method: "PUT",
	}
	r, err := c.request(options)
	if r != nil {
		if r.Code == 204 {
			return r, nil
		}
	}
	return nil, err
}
