package models

type Statistics struct {
	Class             string   `db:"-" json:"class,omitempty"`
	TotalSubmissions  int      `db:"-" json:"total_submissions"`
	UniquePlayerCount int      `db:"-" json:"unique_player_count"`
	MaxTier           *int     `db:"-" json:"max_tier"`
	AverageTier       *int     `db:"-" json:"average_tier,omitempty"`
	PercentageTotal   *float64 `db:"-" json:"percentage_total,omitempty"`
	PercentageUnique  *float64 `db:"-" json:"percentage_unique,omitempty"`
}
