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
	Tier     int    `json:"Tier"`
	Duration string `json:"Time"`
	Name     string `json:"Player"`
	Class    string `json:"Class"`
	Mode     string `json:"Mode"`
	Video    string `json:"Video"`
	Build    string `json:"Build"`
}
type SubmissionsParsed []*SubmissionParsed

func main() {
	content, err := os.ReadFile("seeds/season4.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var s SubmissionsParsed
	err = json.Unmarshal(content, &s)
	if err != nil {
		log.Fatal("Error when parsing file: ", err)
	}

	file, err := os.Create("seeds/season4.csv")
	if err != nil {
		log.Fatal("Error when creating CSV file: ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"name", "class", "tier", "mode", "video", "build", "duration", "verified", "season_id"}
	writer.Write(headers)
	data := [][]string{}
	for _, v := range s {
		var mode = "softcore"
		if v.Mode == "HC" {
			mode = "hardcore"
		}
		dur := strings.TrimSpace(strings.ReplaceAll(v.Duration, ":00", ""))
		duration, err := time.ParseDuration(
			fmt.Sprintf("%ss",
				strings.ReplaceAll(dur, ":", "m"),
			),
		)
		if err != nil {
			log.Fatal("Error when parsing duration: ", err)
		}
		data = append(data, []string{
			strings.TrimSpace(v.Name),
			strings.ToLower(v.Class),
			strconv.FormatInt(int64(v.Tier), 10),
			mode,
			v.Video,
			strings.TrimSpace(v.Build),
			strconv.FormatInt(int64(duration.Seconds()), 10),
			strconv.FormatBool(true),
			"4",
		})
	}
	writer.WriteAll(data)
}
