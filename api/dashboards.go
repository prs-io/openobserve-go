package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prs-io/openobserve-go/models"
)

//TODO:// Implement Below the methods
// func (u *API) CreateDashboard(orgID string, dashboard models.CreateDashboardRequest) (*models.HttpResponse, error) {}
// func (u *API) UpdateDashboard(orgID string, dashboard models.UpdateDashboardRequest) (*models.HttpResponse, error) {}
// func (u *API) MoveDashboard(orgID string, dashboard models.UpdateDashboardRequest) (*models.HttpResponse, error) {}
// func (u *API) DeleteDashboard(orgID string, dashboardID string) (*models.HttpResponse, error) {}

func (a *API) GetDashboard(orgID string, dashboardID string) (*models.DashboardResponse, error) {

	reqPath := buildPath(orgID, "dashboards", dashboardID)
	resp, err := a.client.Do(http.MethodGet, reqPath, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf(" %s", resp.Status)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get dashboards: %s", resp.Status)
	}

	var response models.DashboardResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	fmt.Println("Response...:", response)
	return &response, nil
}

func (a *API) Dashboards(orgID string, folder *string, title *string, pageSize *int64) (*models.DashboardsResponse, error) {

	queryParams := make(map[string]string)

	if folder != nil {
		queryParams["folder"] = *folder
	}
	if title != nil {
		queryParams["title"] = *title
	}
	//TODO: may need to do pagination
	// if pageSize != nil {
	// 	queryParams["pageSize"] = fmt.Sprintf("%d", *pageSize)
	// }

	reqPath := buildPath(orgID, "dashboards")
	if len(queryParams) > 0 {
		q := "?"
		first := true
		for k, v := range queryParams {
			if !first {
				q += "&"
			}
			q += fmt.Sprintf("%s=%s", k, v)
			first = false
		}
		reqPath += q
	}
	resp, err := a.client.Do(http.MethodGet, reqPath, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user: %s", resp.Status)
	}

	var response models.DashboardsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
