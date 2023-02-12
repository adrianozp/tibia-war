package characters

import (
	"time"
)

type CharacterResponse struct {
	Characters  Characters  `json:"characters"`
	Information Information `json:"information"`
}

type CharGuild struct {
	Name string `json:"name"`
	Rank string `json:"rank"`
}

type Character struct {
	Name              string    `json:"name"`
	Sex               string    `json:"sex"`
	Title             string    `json:"title"`
	UnlockedTitles    int       `json:"unlocked_titles"`
	Vocation          string    `json:"vocation"`
	Level             int       `json:"level"`
	AchievementPoints int       `json:"achievement_points"`
	World             string    `json:"world"`
	FormerWorlds      []string  `json:"former_worlds"`
	Residence         string    `json:"residence"`
	Guild             CharGuild `json:"guild"`
	LastLogin         time.Time `json:"last_login"`
	AccountStatus     string    `json:"account_status"`
}

type Killers struct {
	Name   string `json:"name"`
	Player bool   `json:"player"`
	Traded bool   `json:"traded"`
	Summon string `json:"summon"`
}

type Deaths struct {
	Time    time.Time     `json:"time"`
	Level   int           `json:"level"`
	Killers []Killers     `json:"killers"`
	Assists []interface{} `json:"assists"`
	Reason  string        `json:"reason"`
}

type Characters struct {
	Character          Character          `json:"character"`
	Deaths             []Deaths           `json:"deaths"`
}


func (character *Characters) GetLastDeath() (death Deaths) {
	if len(character.Deaths) > 0 {
		death = character.Deaths[0]
	}
	return
}

func (character *Characters) GetDeathsSince(minutes float64, ignoreMonsters bool) (deaths []Deaths) {
	if len(character.Deaths) == 0 {
		return
	}
	for i := 0; i < len(character.Deaths); i++ {
		death := character.Deaths[i]
		if time.Since(death.Time).Minutes() < minutes && death.Killers[0].Player {
			deaths = append(deaths, character.Deaths[i])
		}
	}
	return
}

var shortVocation = map[string]string{
	"Royal Paladin": "RP",
	"Master Sorcerer" : "MS",
	"Elder Druid" : "ED",
	"Elite Knight" : "EK",
	"Paladin": "P",
	"Sorcerer" : "S",
	"Druid" : "D",
	"Knight" : "K",
}

func (character *Character) ShortVocation() (string) {
	return shortVocation[character.Vocation]
}
