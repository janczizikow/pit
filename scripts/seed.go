package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type SubmissionParsed struct {
	Date     string `json:"Date"`
	Tier     string `json:"Tier"`
	Duration string `json:"Time Used"`
	Name     string `json:"Player"`
	Class    string `json:"Class"`
	Mode     string `json:"Mode"`
	Video    string `json:"Run Video"`
	Build    string `json:"Build Planner"`
	Verified string `json:"Validate"`
}
type SubmissionsParsed []*SubmissionParsed

func main() {
	content, err := os.ReadFile("scripts/data.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var s SubmissionsParsed
	err = json.Unmarshal(content, &s)
	if err != nil {
		log.Fatal("Error when parsing file: ", err)
	}

	file, err := os.Create("data.csv")
	if err != nil {
		log.Fatal("Error when creating CSV file: ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"name", "class", "tier", "mode", "build", "video", "duration", "verified", "created_at", "updated_at"}
	writer.Write(headers)
	data := [][]string{}
	for _, v := range s {
		var mode = "softcore"
		if v.Mode == "HC" {
			mode = "hardcore"
		}
		var duration time.Duration
		if strings.Contains(v.Duration, "sec") {
			duration, err = time.ParseDuration(
				strings.ReplaceAll(strings.TrimSpace(v.Duration), " sec", "s"),
			)
			if err != nil {
				log.Fatal("Error when parsing duration: ", err)
			}
		} else {
			duration, err = time.ParseDuration(
				fmt.Sprintf("%ss",
					strings.ReplaceAll(strings.TrimSpace(v.Duration), ":", "m"),
				),
			)
			if err != nil {
				log.Fatal("Error when parsing duration: ", err)
			}
		}
		date := strings.Split(v.Date, "/")
		if len(date[0]) == 1 {
			date[0] = fmt.Sprintf("0%s", date[0])
		}
		created, err := time.Parse(time.DateOnly, fmt.Sprintf("%s-%s-%s", date[2], date[0], date[1]))
		if err != nil {
			log.Fatal("Error when parsing date: ", err)
		}
		data = append(data, []string{
			v.Name,
			strings.ToLower(v.Class),
			v.Tier,
			mode,
			v.Build,
			v.Video,
			strconv.FormatInt(int64(duration.Seconds()), 10),
			strconv.FormatBool(v.Verified == "y"),
			created.Format(time.RFC3339),
			created.Format(time.RFC3339),
		})
	}
	writer.WriteAll(data)
}
