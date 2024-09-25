package bot

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/janczizikow/pit/internal/models"
	"github.com/janczizikow/pit/internal/repository"
	zlog "github.com/rs/zerolog/log"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "list",
		Description: "List top entries from pit leaderboard",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "class",
				Description: "Filter by specific class",
				Type:        discordgo.ApplicationCommandOptionString,
			},
			{
				Name:        "mode",
				Description: "Filter by specific mode (softcore/hardcore). Defaults to softcore",
				Type:        discordgo.ApplicationCommandOptionString,
			},
		},
	},
}

type discordHandler struct {
	repo repository.SeasonSubmissionsRepository
}

func (h *discordHandler) List(s *discordgo.Session, i *discordgo.InteractionCreate, opts optionMap) {
	builder := new(strings.Builder)

	// builder.WriteString(opts["message"].StringValue())
	params := repository.ListSubmissionsParams{
		Class:   "",
		Mode:    "softcore",
		OrderBy: "tier DESC, duration ASC",
		Limit:   5,
		Offset:  0,
	}
	submissions, _, err := h.repo.List(5, params)
	if err != nil {
		return
	}

	builder.WriteString("`")
	cols := []string{"player", "class", "tier", "time", "video", "build", "date"}
	j := 0
	PrintTable(builder, "Ranking", cols, func() (bool, [][]string) {
		if len(submissions)-1 < j {
			return false, nil
		}
		s := submissions[j]
		j++
		return true, submissionToRow(s)
	})
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

func submissionToRow(sub *models.Submission) [][]string {
	rows := make([][]string, 0, 1)
	rows = append(rows, []string{
		sub.Name,
		sub.Class,
		strconv.FormatInt(int64(sub.Tier), 10),
		(time.Duration(sub.Duration) * time.Second).String(),
		"-",
		"-",
		sub.CreatedAt.Format(time.DateOnly),
	})
	return rows
}

func PrintTable(
	w io.Writer,
	caption string,
	cols []string,
	getNextRows func() (bool, [][]string),
) {
	nbCols := len(cols)
	if nbCols == 0 {
		return
	}
	if len(caption) > 0 {
		fmt.Fprintf(w, "%s:\n\n", caption)
	}
	tw := tabwriter.NewWriter(w, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	colSep := "\t"
	header := append([]string{"#"}, cols...)
	fmt.Fprintf(tw, "%s%s\n", strings.Join(header, colSep), colSep)
	var sb strings.Builder
	i := 1
	hasMore, rows := getNextRows()
	for hasMore {
		for iRow, row := range rows {
			nbRowCols := len(row)
			for j := 0; j < nbCols; j++ {
				if j < nbRowCols {
					sb.WriteString(row[j])
				}
				sb.WriteString("\t")
			}
			iRowStr := ""
			if iRow == 0 {
				iRowStr = strconv.Itoa(i)
			}
			fmt.Fprintf(tw, "%s%s%s\n", iRowStr, colSep, sb.String())
			sb.Reset()
		}
		i++
		hasMore, rows = getNextRows()
	}
	_ = tw.Flush()
}
