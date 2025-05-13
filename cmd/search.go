package cmd

import (
	"fmt"
	"time"

	"github.com/prs-io/openobserve-go/models"

	"github.com/spf13/cobra"
)

var (
	//sinceTime    string
	since    time.Duration
	sqlQuery string
	size     int64 = 50 // default value: 50
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Perform a search query",
	Run: func(cmd *cobra.Command, args []string) {
		if since == 0 {
			cobra.CheckErr("--since flag is required")
		}
		if sqlQuery == "" {
			cobra.CheckErr("--sql flag is required")
		}

		startTime := time.Now().Add(-since).UnixMilli()
		endTime := time.Now().UnixMilli()
		searchQuery := models.NewSearchQuery(startTime, endTime, sqlQuery, nil)
		searchRequest := models.NewSearchRequest(*searchQuery)
		searchResponse, err := API.Search(orgID, searchRequest)
		if err != nil {
			cobra.CheckErr(err)
			return
		}
		for _, hit := range searchResponse.Hits {
			timestamp := time.UnixMicro(int64(hit["_timestamp"].(float64))).Format(time.RFC3339)
			fmt.Println(timestamp, hit["event_log_severitytext"].(string), hit["event_log_body"].(string))
		}
	},
}

func init() {
	//TODO: add flags for search command
	//searchCmd.Flags().StringVar(&sinceTime, "since-time", sinceTime, i18n.T("Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used."))
	searchCmd.Flags().DurationVar(&since, "since", since, "Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used.")
	searchCmd.Flags().StringVar(&sqlQuery, "sql", "", "SQL query string")
	searchCmd.Flags().Int64Var(&size, "size", size, "Number of results to return")
}
