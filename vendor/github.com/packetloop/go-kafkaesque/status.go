package gokafkaesque

import "fmt"

// GetStatus returns status kafka-admin-service /health endpoint.
func (client *Client) GetStatus() (GenericResponse, error) {
	//func (client *Client) getHealth() string {
	//return client.Rest.HostURL
	resp, err := client.Rest.R().Get("/health")
	if err != nil {
		return GenericResponse{}, fmt.Errorf("ERROR: %s", err.Error())
	}
	if resp.StatusCode() >= 200 && resp.StatusCode() <= 299 {
		var data GenericResponse
		err := client.Rest.JSONUnmarshal(resp.Body(), &data)
		if err != nil {
			return GenericResponse{}, fmt.Errorf("ERROR: %s", err.Error())
		}
		return data, nil
	}
	return GenericResponse{}, fmt.Errorf("ERROR: %v", resp.Status())
}

// GetResponse is a method that returns actual health status of "ok".
func (h *GenericResponse) GetResponse() string {
	return h.Response
}
