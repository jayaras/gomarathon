package gomarathon

import (
	"fmt"
)

/*
* Group support
 */

// ListApps is listing all apps
func (c *Client) ListGroups() (*DefaultResponse, error) {
	return c.GetGroup("/")
}

// ListGroupVersions list app versions from a define appid
func (c *Client) ListGroupVersions(groupID string) (*DefaultResponse, error) {
	var r DefaultResponse
	options := &RequestOptions{
		Path: fmt.Sprintf("groups/%s/versions", groupID),
	}
	err := c.request(options, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

// GetGroup gets a group
func (c *Client) GetGroup(groupID string) (*DefaultResponse, error) {
	var r DefaultResponse

	options := &RequestOptions{
		Path: fmt.Sprintf("groups/%s", groupID),
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
func (c *Client) GetGroupVersion(groupID string, version string) (*DefaultResponse, error) {
	var r DefaultResponse

	options := &RequestOptions{
		Path: fmt.Sprintf("groups/%s/versions/%s", groupID, version),
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

// CreateGroup Create a new group
func (c *Client) CreateGroup(group *Group) (*DefaultResponse, error) {
	var r DefaultResponse

	// TODO : VALIDATE DATAS
	options := &RequestOptions{
		Path:   "groups",
		Datas:  group,
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

func (c *Client) UpdateGroup(groupID string, group *Group) (*DefaultResponse, error) {
	var r DefaultResponse

	options := &RequestOptions{
		Path:   fmt.Sprintf("groups/%s", groupID),
		Datas:  group,
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

// DeleteGroup delete this group of apps from the cluster
func (c *Client) DeleteGroup(groupID string, force bool) (*DefaultResponse, error) {
	var r DefaultResponse

	options := &RequestOptions{
		Path:   fmt.Sprintf("groups/%s", groupID),
		Method: "DELETE",
		Params: &Parameters{Force: force},
	}
	err := c.request(options, &r)
	if &r != nil {
		if r.GetCode() == 204 {
			return &r, nil
		}
	}
	return nil, err
}

func (c *Client) ForceDeleteGroup(groupID string) (*DefaultResponse, error) {
	//var r DefaultResponse

	return c.DeleteGroup(groupID, true)
}

// RollbackGroup  rollback group to a previous version
func (c *Client) RollbackGroup(groupID string, version string) (*DefaultResponse, error) {
	var r DefaultResponse

	options := &RequestOptions{
		Path:   fmt.Sprintf("groups/%s/version/%s", groupID, version),
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
