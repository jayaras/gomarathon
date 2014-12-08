package gomarathon

import "fmt"

func (c *Client) ListDeployments() (*DeploymentResponse, error) {
	var r DeploymentResponse

	options := &RequestOptions{
		Path: "deployments",
	}
	err := c.request(options, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (c *Client) DeleteDeployment(deploymentID string) (*DefaultResponse, error) {

	var r DefaultResponse

	options := &RequestOptions{
		Path:   fmt.Sprintf("deployments/%s", deploymentID),
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
