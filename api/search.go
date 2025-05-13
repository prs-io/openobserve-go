package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prs-io/openobserve-go/models"
)

func (a *API) Search(orgID string, searchRequest *models.SearchRequest) (*models.SearchResponse, error) {

	var reqPath UrlPath
	var err error

	if err = reqPath.Set(orgID, "_search"); err != nil {
		return nil, err
	}
	resp, err := a.client.Do(http.MethodPost, reqPath.String(), searchRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, a.Error(resp)
	}

	var response *models.SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}

func (a *API) SearchHistory(orgID string, searchRequest *models.SearchHistoryRequest) (*models.SearchResponse, error) {

	var reqPath UrlPath
	var err error

	if err = reqPath.Set(orgID, "_search_history"); err != nil {
		return nil, err
	}
	fmt.Println("reqPath", reqPath.String())
	resp, err := a.client.Do(http.MethodPost, reqPath.String(), searchRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, a.Error(resp)
	}

	var response *models.SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}

func (a *API) SearchPartition(orgID string, searchRequest *models.SearchPartitionRequest) (*models.SearchPartitionResponse, error) {

	var reqPath UrlPath
	var err error

	if err = reqPath.Set(orgID, "_search_partition"); err != nil {
		return nil, err
	}
	fmt.Println("reqPath", reqPath.String())
	resp, err := a.client.Do(http.MethodPost, reqPath.String(), searchRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, a.Error(resp)
	}

	var response *models.SearchPartitionResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}
