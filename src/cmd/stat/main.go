package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"github.com/gosuri/uitable"

	"github.com/VKCOM/noverify/src/linter"
)

type linterOutput struct {
	Reports []*linter.Report
	Errors  []string
}

func loadReportsFile(filename string) *linterOutput {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("read reports file: %v", err)
	}
	var output linterOutput
	if err := json.Unmarshal(data, &output); err != nil {
		log.Fatalf("unmarshal reports file: %v", err)
	}
	return &output
}

type ReportCount struct {
	CheckName string
	Count     int
}

func main() {
	reports := loadReportsFile("reports.json").Reports

	var reportsByType = map[string]int{}

	for _, report := range reports {
		reportsByType[report.CheckName]++
	}

	var reportsCount []ReportCount
	for checkName, count := range reportsByType {
		reportsCount = append(reportsCount, ReportCount{
			CheckName: checkName,
			Count:     count,
		})
	}
	sort.Slice(reportsCount, func(i, j int) bool {
		return reportsCount[i].Count > reportsCount[j].Count
	})

	table := uitable.New()

	table.AddRow("Name", "Count")

	for _, info := range reportsCount {
		table.AddRow(info.CheckName, info.Count)
	}

	fmt.Printf("Count: %d\n\n", len(reports))
	fmt.Println(table)
}
