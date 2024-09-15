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
	Timestamp time.Time `json:"Timestamp"`
	Date      time.Time `json:"Date"`
	Tier      int       `json:"Tier"`
	Duration  string    `json:"Time Used"`
	Name      string    `json:"Player"`
	Class     string    `json:"Class"`
	Mode      string    `json:"Mode"`
	Video     string    `json:"Run Video"`
	Build     string    `json:"Build Planner"`
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

	headers := []string{"name", "class", "tier", "mode", "build", "video", "duration", "created_at", "updated_at"}
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
		build := v.Build
		if build == "" {
			build = `""`
		}
		data = append(data, []string{
			v.Name,
			strings.ToLower(v.Class),
			strconv.FormatInt(int64(v.Tier), 10),
			mode,
			build,
			v.Video,
			strconv.FormatInt(int64(duration.Seconds()), 10),
			v.Date.Format(time.RFC3339),
			v.Date.Format(time.RFC3339),
		})
	}
	writer.WriteAll(data)
}
