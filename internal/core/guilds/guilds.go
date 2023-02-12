package guilds

import (
	"time"
)

type GuildResponse struct {
	Guilds      Guilds      `json:"guilds"`
	Information Information `json:"information"`
}

type Members struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Rank     string `json:"rank"`
	Vocation string `json:"vocation"`
	Level    int    `json:"level"`
	Joined   string `json:"joined"`
	Status   string `json:"status"`
}

type Guild struct {
	Name             string      `json:"name"`
	World            string      `json:"world"`
	LogoURL          string      `json:"logo_url"`
	Description      string      `json:"description"`
	Guildhalls       interface{} `json:"guildhalls"`
	Active           bool        `json:"active"`
	Founded          string      `json:"founded"`
	OpenApplications bool        `json:"open_applications"`
	Homepage         string      `json:"homepage"`
	InWar            bool        `json:"in_war"`
	DisbandDate      string      `json:"disband_date"`
	DisbandCondition string      `json:"disband_condition"`
	PlayersOnline    int         `json:"players_online"`
	PlayersOffline   int         `json:"players_offline"`
	MembersTotal     int         `json:"members_total"`
	MembersInvited   int         `json:"members_invited"`
	Members          []Members   `json:"members"`
	Invites          interface{} `json:"invites"`
}

type Guilds struct {
	Guild Guild `json:"guild"`
}

type Information struct {
	APIVersion int       `json:"api_version"`
	Timestamp  time.Time `json:"timestamp"`
}

func (guild *Guild) OnlineLevelRange(online bool, min int, max int) (int, []Members) {
	var players []Members
	for p := 0; p < len(guild.Members); p++ {
		player := guild.Members[p]
		if (!online || player.Status == "online") && player.Level > min && player.Level < max {
			players = append(players, player)
		}
	}
	return len(players), players
}