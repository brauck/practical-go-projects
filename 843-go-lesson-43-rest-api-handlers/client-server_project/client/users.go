package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *APIClient) GetUsers() ([]User, error) {
	resp, err := api.Client.Get(api.BaseURL + "/users")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func (api *APIClient) CreateUser(u User) (*User, error) {
	body, _ := json.Marshal(u)

	resp, err := api.Client.Post(
		api.BaseURL+"/users",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var created User
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return nil, err
	}

	return &created, nil
}

func (api *APIClient) UpdateUser(u User) (*User, error) {
	body, _ := json.Marshal(u)

	req, err := http.NewRequest(http.MethodPut, api.BaseURL+"/users", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var updated User
	if err := json.NewDecoder(resp.Body).Decode(&updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

func (api *APIClient) DeleteUser(id int) error {
	body, _ := json.Marshal(struct {
		ID int `json:"id"`
	}{ID: id})

	req, err := http.NewRequest(http.MethodDelete, api.BaseURL+"/users", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("delete failed: status %d", resp.StatusCode)
	}

	return nil
}
