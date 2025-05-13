package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/prs-io/openobserve-go/models"
)

//TODO: Implement the rest of the API methods

// func (a *API) CreateAlert(orgID string, alert Alert) error                     {}
// func (a *API) UpdateAlert(orgID string, alertID string, alert Alert) error     {}
// func (a *API) MoveAlerts(orgID string, alertID string, alert Alert) error      {}
// func (a *API) DeleteAlert(orgID string, alertID string) error                  {}

// func (a *API) EnableAlerts(orgID string, page int, size int) ([]Alert, error)  {}
// func (a *API) TriggerAlerts(orgID string, page int, size int) ([]Alert, error) {}

// func (a *API) ListAlertTemplates(orgID string) ([]AlertTemplate, error)                          {}
// func (a *API) GetAlertTemplate(orgID string, templateID string) (*AlertTemplate, error)          {}
// func (a *API) CreateAlertTemplate(orgID string, template AlertTemplate) error                    {}
// func (a *API) UpdateAlertTemplate(orgID string, templateID string, template AlertTemplate) error {}
// func (a *API) DeleteAlertTemplate(orgID string, templateID string) error                         {}

// func (a *API) ListDestinations(orgID string) ([]AlertRule, error)                          {}
// func (a *API) GetDestination(orgID string, destinationID string) (*AlertRule, error)     {}

func (a *API) ListAlerts(orgID string, folder *string, owner *string, enabled *bool, pageSize *int64, streamType *string, streamName *string) (*models.ListAlertsResponse, error) {

	queryParams := make(map[string]string)

	if folder != nil {
		queryParams["folder"] = *folder
	}
	if owner != nil {
		queryParams["owner"] = *owner
	}
	if enabled != nil {
		queryParams["enabled"] = strconv.FormatBool(*enabled)
	}
	if pageSize != nil {
		queryParams["pageSize"] = strconv.FormatInt(*pageSize, 10)
	}
	if streamType != nil {
		queryParams["streamType"] = *streamType
	}
	if streamName != nil {
		queryParams["streamName"] = *streamName
	}

	reqPath := buildPath("v2", orgID, "alerts")
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
		return nil, fmt.Errorf("failed to get Alerts: %s", resp.Status)
	}

	var response models.ListAlertsResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (a *API) GetAlert(orgID string, alertID string) (*models.Alert, error) {
	reqPath := buildPath("v2", orgID, "alerts", alertID)

	resp, err := a.client.Do(http.MethodGet, reqPath, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get Alerts: %s", resp.Status)
	}

	var response models.Alert
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil

}
