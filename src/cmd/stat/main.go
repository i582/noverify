package main

import (
	"encoding/json"
	"flag"
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
	var markdown bool
	flag.BoolVar(&markdown, "m", false, "print with markdown table")
	flag.Parse()

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

	if markdown {
		fmt.Println("## The number of reports for tests.")
		fmt.Printf("Count: %d\n\n", len(reports))

		fmt.Println("Name | Count")
		fmt.Println("---- | :---:")

		for _, info := range reportsCount {
			fmt.Printf("%s | %d (%.2f%%)\n", info.CheckName, info.Count, percent(info.Count, len(reports)))
		}
	} else {
		table := uitable.New()

		table.AddRow("Name", "Count")

		for _, info := range reportsCount {
			table.AddRow(info.CheckName, fmt.Sprintf("%d (%.2f%%)", info.Count, percent(info.Count, len(reports))))
		}

		fmt.Println("The number of reports for tests.")
		fmt.Printf("Count: %d\n\n", len(reports))
		fmt.Println(table)
	}
}

func percent(x, y int) float64 {
	return float64(x) / float64(y) * 100
}
