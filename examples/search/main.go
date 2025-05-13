package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/prs-io/openobserve-go/api"
	"github.com/prs-io/openobserve-go/client"
	"github.com/prs-io/openobserve-go/models"

	logparser "github.com/prs-io/plexus-logparser"
)

func main() {
	orgID := "default"
	fmt.Println("Org ID:", orgID)
	cfg := client.Config{
		BaseURL:  "https://openobserve.io",
		Username: os.Getenv("OPENOBSERVE_USERNAME"),
		Password: os.Getenv("OPENOBSERVE_PASSWORD"),
		Timeout:  10 * time.Second,
	}
	fmt.Printf("Config: %+v\n", cfg)
	openObserveClient := client.NewClient(cfg)

	a := api.NewAPI(openObserveClient)
	// boolVal := true
	now := time.Now()
	duration := time.Duration(900) * time.Second
	startTime := now.Add(-duration).UnixMilli()
	endTime := now.UnixMilli()
	// startTime = 1746805899809000
	// endTime = 1746806799809000

	//e.g.: Use logparser to get count of severity
	ch := make(chan logparser.LogEntry)
	parser := logparser.NewParser(ch, nil, nil, time.Second)

	//SQL Search
	searchQuery := models.NewSearchQuery(startTime, endTime, "select * from k8s", nil)
	searchRequest := models.NewSearchRequest(*searchQuery)
	searchResponse, err := a.Search(orgID, searchRequest)
	t := time.Now()
	for _, hit := range searchResponse.Hits {
		//fmt.Printf("HITs: %+v\n", hit["event_log_severitytext"])
		ch <- logparser.LogEntry{Timestamp: time.UnixMicro(int64(hit["_timestamp"].(float64))), Content: hit["event_log_body"].(string), Level: logparser.LevelFromString(hit["event_log_severitytext"].(string))}
		//fmt.Printf("LogEntry: %+v\n", le)
	}
	d := time.Since(t)
	defer parser.Stop()
	counters := parser.GetCounters()
	order(counters)

	grandTotal, total, max := 0, 0, 0
	for _, c := range counters {
		grandTotal += c.Messages
		if c.Sample == "" {
			continue
		}
		total += c.Messages
		if c.Messages > max {
			max = c.Messages
		}
	}
	// Print the results
	output(counters, 120, 100, d)

	// Search History
	// searchHistoryRequest := models.NewSearchHistoryRequest(startTime, endTime, "k8s", nil)
	// fmt.Printf("Search History Request: %+v\n", searchHistoryRequest)
	// searchResponse, err := a.SearchHistory(orgID, searchHistoryRequest)

	//SearchPartition
	// searchPartitionRequest := models.NewSearchPartitionRequest(startTime, endTime, "select * from k8s")
	// searchResponse, err := a.SearchPartition(orgID, searchPartitionRequest)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// fmt.Printf("Search Response:%+v", searchResponse)
}

func output(counters []logparser.LogCounter, screenWidth, maxLinesPerMessage int, duration time.Duration) {
	grandTotal, total, max := 0, 0, 0
	for _, c := range counters {
		grandTotal += c.Messages
		if c.Sample == "" {
			continue
		}
		total += c.Messages
		if c.Messages > max {
			max = c.Messages
		}
	}
	barWidth := 20
	lineWidth := screenWidth - barWidth
	messagesNumFmt := fmt.Sprintf("%%%dd", len(strconv.Itoa(max)))
	for _, c := range counters {
		if c.Sample == "" {
			continue
		}
		w := c.Messages * barWidth / max
		bar := strings.Repeat("▇", w+1) + strings.Repeat(" ", barWidth-w)
		prefix := colorize(c.Level, "%s "+messagesNumFmt+" (%2d%%) ", bar, c.Messages, int(float64(c.Messages*100)/float64(total)))
		sample := ""
		for i, line := range strings.Split(c.Sample, "\n") {
			if len(line) > lineWidth {
				line = line[:lineWidth] + "..."
			}
			sample += line + "\n" + strings.Repeat(" ", len(prefix))
			if i > maxLinesPerMessage {
				sample += "...\n"
				break
			}
		}
		sample = strings.TrimRight(sample, "\n ")
		fmt.Printf("%s%s\n", prefix, sample)
	}

	byLevel := map[logparser.Level]int{}
	for _, c := range counters {
		byLevel[c.Level] += c.Messages
	}
	fmt.Println()
	fmt.Printf("%d messages processed in %.3f seconds:\n", grandTotal, duration.Seconds())
	for l, c := range byLevel {
		fmt.Printf("  %s: %d\n", l, c)
	}
	fmt.Println()
}

func colorize(level logparser.Level, format string, a ...interface{}) string {
	c := "\033[37m" // grey
	switch level {
	case logparser.LevelCritical, logparser.LevelError:
		c = "\033[31m" // red
	case logparser.LevelWarning:
		c = "\033[33m" // yellow
	case logparser.LevelInfo:
		c = "\033[32m" // green
	}
	return fmt.Sprintf(c+format+"\033[0m", a...)
}

func order(counters []logparser.LogCounter) {
	sort.Slice(counters, func(i, j int) bool {
		ci, cj := counters[i], counters[j]
		if ci.Level == cj.Level {
			return ci.Messages > cj.Messages
		}
		return ci.Level < cj.Level
	})
}
