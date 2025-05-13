package api

import (
	"encoding/json"
	"net/http"

	"github.com/prs-io/openobserve-go/models"
)

func (u *API) CreateUser(orgID string, user models.UserRequest) (*models.HttpResponse, error) {

	reqPath := buildPath(orgID, "users")
	resp, err := u.client.Do(http.MethodPost, reqPath, user)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, u.Error(resp)
	}

	var response models.HttpResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (u *API) AddUserToOrg(orgID string, email string, role models.UserOrgRole) (*models.HttpResponse, error) {
	reqPath := buildPath(orgID, "users", email)
	resp, err := u.client.Do(http.MethodPut, reqPath, role)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, u.Error(resp)
	}

	var response models.HttpResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (u *API) UpdateUser(orgID string, email string, updateReq models.UpdateUser) (*models.HttpResponse, error) {

	reqPath := buildPath(orgID, "users", email)
	resp, err := u.client.Do(http.MethodPut, reqPath, updateReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, u.Error(resp)
	}

	var response models.HttpResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (u *API) DeleteUser(orgID string, email string) (*models.HttpResponse, error) {

	reqPath := buildPath(orgID, "users", email)
	resp, err := u.client.Do(http.MethodDelete, reqPath, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return nil, u.Error(resp)
	}
	var response models.HttpResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (u *API) ListUser(orgID string) (*models.UserList, error) {
	reqPath := buildPath(orgID, "users")
	resp, err := u.client.Do(http.MethodGet, reqPath, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, u.Error(resp)
	}

	var response models.UserList
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
