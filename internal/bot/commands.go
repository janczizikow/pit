package bot

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/bwmarrin/discordgo"
	"github.com/janczizikow/pit/internal/repository"
	zlog "github.com/rs/zerolog/log"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "leaderboard",
		Description: "List top entries from pit leaderboard",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "class",
				Description: "Filter by specific class",
				Type:        discordgo.ApplicationCommandOptionString,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Barbarian",
						Value: "barbarian",
					},
					{
						Name:  "Druid",
						Value: "druid",
					},
					{
						Name:  "Necromancer",
						Value: "necromancer",
					},
					{
						Name:  "Rogue",
						Value: "rogue",
					},
					{
						Name:  "Sorcerer",
						Value: "sorcerer",
					},
				},
			},
			{
				Name:        "mode",
				Description: "Filter by specific mode (softcore/hardcore). Defaults to softcore",
				Type:        discordgo.ApplicationCommandOptionString,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Softcore",
						Value: "softcore",
					},
					{
						Name:  "Hardcore",
						Value: "hardcore",
					},
				},
			},
			{
				Name:        "season",
				Description: "Filter by season. Defaults to season 5",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Season 4",
						Value: 4,
					},
					{
						Name:  "Season 5",
						Value: 5,
					},
				},
			},
			{
				Name:        "page",
				Description: "Find specific page (each page = 10 listings)",
				Type:        discordgo.ApplicationCommandOptionInteger,
			},
		},
	},
}

type discordHandler struct {
	repo repository.SeasonSubmissionsRepository
}

func getString(opt *discordgo.ApplicationCommandInteractionDataOption, fallback string) string {
	if opt == nil {
		return fallback
	}

	return opt.StringValue()
}

func getInt(opt *discordgo.ApplicationCommandInteractionDataOption, fallback int) int {
	if opt == nil {
		return fallback
	}

	return int(opt.IntValue())
}

func (h *discordHandler) List(s *discordgo.Session, i *discordgo.InteractionCreate, opts optionMap) {
	builder := new(strings.Builder)
	class := getString(opts["class"], "")
	mode := getString(opts["mode"], "softcore")
	seasonId := getInt(opts["season"], 5)
	page := getInt(opts["page"], 1)
	pageSize := 10
	if page < 1 {
		builder.WriteString("⚠️ Error: page must be greater than zero")
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: builder.String(),
			},
		})
		return
	}

	offset := (page - 1) * pageSize

	params := repository.ListSubmissionsParams{
		Class:   class,
		Mode:    mode,
		OrderBy: "tier DESC, duration ASC",
		Limit:   pageSize,
		Offset:  offset,
	}
	submissions, _, err := h.repo.List(seasonId, params)
	if err != nil {
		zlog.Error().Err(err).Msg("failed to fetch submissions for bot")
		builder.WriteString("⚠️ Error: Something went wrong, please try again")
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: builder.String(),
			},
		})
		return
	}

	builder.WriteString("`")
	var table = [][]string{
		{"#", "player", "class", "tier", "time", "date"},
	}
	for j, sub := range submissions {
		table = append(
			table,
			[]string{
				strconv.FormatInt(int64(j+1+offset), 10),
				sub.Name,
				sub.Class,
				strconv.FormatInt(int64(sub.Tier), 10),
				(time.Duration(sub.Duration) * time.Second).String(),
				// "-",
				// "-",
				sub.CreatedAt.Format(time.DateOnly),
			},
		)
	}
	printTable(builder, table)
	builder.WriteString("`")

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: builder.String(),
		},
	})

	if err != nil {
		zlog.Error().Err(err).Msg("could not respond to interaction")
	}
}

// TODO: fix width of diff chars eg.: chinese characters
func printTable(w io.Writer, table [][]string) {
	// get number of columns from the first table row
	columnLengths := make([]int, len(table[0]))
	for _, line := range table {
		for i, val := range line {
			if utf8.RuneCountInString(val) > columnLengths[i] {
				columnLengths[i] = utf8.RuneCountInString(val)
			}
		}
	}

	var lineLength int
	for _, c := range columnLengths {
		lineLength += c + 3 // +3 for 3 additional characters before and after each field: "| %s "
	}
	lineLength += 1 // +1 for the last "|" in the line

	for i, line := range table {
		if i == 0 { // table header
			fmt.Fprintf(w, "+%s+\n", strings.Repeat("-", lineLength-2)) // lineLength-2 because of "+" as first and last character
		}
		for j, val := range line {
			fmt.Fprintf(w, "| %-*s ", columnLengths[j], val)
			if j == len(line)-1 {
				fmt.Fprintf(w, "|\n")
			}
		}
		if i == 0 || i == len(table)-1 { // table header or last line
			fmt.Fprintf(w, "+%s+\n", strings.Repeat("-", lineLength-2)) // lineLength-2 because of "+" as first and last character
		}
	}
}
