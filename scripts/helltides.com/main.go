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
	Tier     string `json:"tier"`
	Duration string `json:"time"`
	Name     string `json:"name"`
	Class    string `json:"class"`
	Mode     string `json:"mode"`
	Video    string `json:"video"`
	Build    string `json:"build"`
}
type SubmissionsParsed []*SubmissionParsed

func main() {
	content, err := os.ReadFile("seeds/helltides.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var s SubmissionsParsed
	err = json.Unmarshal(content, &s)
	if err != nil {
		log.Fatal("Error when parsing file: ", err)
	}

	file, err := os.Create("seeds/helltides.csv")
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
		duration, err := time.ParseDuration(
			fmt.Sprintf("%ss",
				strings.ReplaceAll(strings.TrimSpace(v.Duration), ":", "m"),
			),
		)
		if err != nil {
			log.Fatal("Error when parsing duration: ", err)
		}
		data = append(data, []string{
			strings.TrimSpace(v.Name),
			strings.ToLower(v.Class),
			v.Tier,
			v.Mode,
			v.Video,
			strings.TrimSpace(v.Build),
			strconv.FormatInt(int64(duration.Seconds()), 10),
			strconv.FormatBool(true),
			"5",
		})
	}
	writer.WriteAll(data)
}
