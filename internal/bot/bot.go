package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/janczizikow/pit/internal/repository"
	zlog "github.com/rs/zerolog/log"
)

type optionMap = map[string]*discordgo.ApplicationCommandInteractionDataOption

func parseOptions(options []*discordgo.ApplicationCommandInteractionDataOption) (om optionMap) {
	om = make(optionMap)
	for _, opt := range options {
		om[opt.Name] = opt
	}
	return om
}

func Start(db *pgxpool.Pool) (*discordgo.Session, error) {
	cfg, err := ReadConfig()
	if err != nil {
		return nil, err
	}
	session, err := discordgo.New("Bot " + cfg.Token)

	if err != nil {
		return nil, err
	}
	repo := repository.NewSeasonSubmissionsRepository(db)

	handler := &discordHandler{repo: repo}

	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		defer func() {
			if err := recover(); err != nil {
				zlog.Error().Msg("recovered from panic in bot handler")
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Error ⚠️ Something went wrong, please try again.",
					},
				})
			}
		}()
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		data := i.ApplicationCommandData()
		if data.Name != "leaderboard" {
			return
		}
		switch data.Name {
		case "leaderboard":
			handler.List(s, i, parseOptions(data.Options))
		}
	})

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		zlog.Info().Msg(fmt.Sprintf("Logged in as %s", r.User.String()))
	})

	_, err = session.ApplicationCommandBulkOverwrite(cfg.App, cfg.Guild, commands)
	if err != nil {
		zlog.Error().Err(err).Msg("could not register commands")
	}

	err = session.Open()

	if err != nil {
		return nil, err
	}

	return session, nil
}
