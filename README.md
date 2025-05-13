# openobserve-go

A Go client library for OpenObserve, supporting logs, metrics, traces, and user management.

> **Note:** This library is still in development. Use with caution as APIs and functionality may change.

## OpenObserve REST API

- [API Documentation](https://openobserve.ai/docs/api/)
- [Swagger UI](https://<openobserve server>/swagger/)

## Install

To install the library, run:

```bash
go get github.com/prs-io/openobserve-go
```

## Usage Example

Here’s how you can use the client to perform a search:

```go
package main

import (
    "fmt"
    "log"
    "time"

   	"github.com/prs-io/openobserve-go/api"
	"github.com/prs-io/openobserve-go/client"
	"github.com/prs-io/openobserve-go/models"
)

func main() {
    //Config
    orgID := "default"
	fmt.Println("Org ID:", orgID)
	cfg := client.Config{
		BaseURL:  "https://openobserve.io",
		Username: os.Getenv("OPENOBSERVE_USERNAME"),
		Password: os.Getenv("OPENOBSERVE_PASSWORD"),
		Timeout:  10 * time.Second,
	}
	fmt.Printf("Config: %+v\n", cfg)

    //Client
	openObserveClient := client.NewClient(cfg)

    //API
    a := api.NewAPI(openObserveClient)

    	Search History
	searchHistoryRequest := models.NewSearchHistoryRequest(startTime, endTime, "k8s", nil)
	fmt.Printf("Search History Request: %+v\n", searchHistoryRequest)
	searchResponse, err := a.SearchHistory(orgID, searchHistoryRequest)
    if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// fmt.Printf("Search Response:%+v", searchResponse)
}
```

More examples can be found [here](examples/).
