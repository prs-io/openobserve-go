package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/prs-io/openobserve-go/client"
	"github.com/prs-io/openobserve-go/models"
)

// TODO: ctx

const basePath = "/api"

type API struct {
	client *client.Client
}

func NewAPI(c *client.Client) *API {
	return &API{client: c}
}

func buildPath(args ...string) string {
	path, _ := url.JoinPath(basePath, args...)
	return path
}

type UrlPath string

func (p *UrlPath) String() string {
	return string(*p)
}

func (p *UrlPath) Set(args ...string) error {
	path, err := url.JoinPath(basePath, args...)
	if err != nil {
		return fmt.Errorf("failed to Set URL path: %w", err)
	}
	*p = UrlPath(path)
	return nil
}

func (p *UrlPath) AddQueryParams(queryParams map[string]string) error {
	reqPath := string(*p)
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
	*p = UrlPath(reqPath)
	return nil
}

func (a *API) ShortenURL(orgID string, req models.ShortenUrlRequest) (*models.ShortenUrlResponse, error) {

	reqPath := buildPath(orgID, "short")

	resp, err := a.client.Do(http.MethodPost, reqPath, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create shorten url: %s", resp.Status)
	}

	var response models.ShortenUrlResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (a *API) RetrieveShortURL(orgID string, shortID string) (*models.ShortenUrlRedirect, error) {

	reqPath := fmt.Sprintf("/short/%s/short/%s", orgID, shortID)
	resp, err := a.client.Do(http.MethodGet, reqPath, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusInternalServerError {
		return nil, fmt.Errorf("failed to retrieve shorten url: %s", resp.Status)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("%s", resp.Status)
	}

	fmt.Printf("Response..: %+v\n", resp)
	var response models.ShortenUrlRedirect
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (a *API) ListClusters() (*map[string][]string, error) {
	reqPath := buildPath("clusters")
	resp, err := a.client.Do(http.MethodGet, reqPath, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list clusters: %s", resp.Status)
	}

	//TODO: create a meaning full response struct
	var response map[string][]string
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (a *API) Error(resp *http.Response) error {
	var errResp models.HttpResponse
	if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
		return err
	}
	return fmt.Errorf("%s %s", resp.Status, errResp.Message)
}
