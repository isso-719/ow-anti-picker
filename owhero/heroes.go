package owhero

type HeroID int

const (
	Ana HeroID = iota
	Ashe
	Baptiste
	Bastion
	Brigitte
	Cassidy
	DVa
	Doomfist
	Echo
	Genji
	Hanzo
	JunkerQueen
	Junkrat
	Kiriko
	LifeWeaver
	Lucio
	Mei
	Mercy
	Moira
	Orisa
	Pharah
	Ramattra
	Reaper
	Reinhardt
	Roadhog
	Sigma
	Sojourn
	Soldier76
	Sombra
	Symmetra
	Torbjorn
	Tracer
	Widowmaker
	Winston
	WreckingBall
	Zarya
	Zenyatta
)

type Hero struct {
	ID   HeroID
	Name string
	Role Role
}

var HeroList = []Hero{
	{Ana, "Ana", Support},
	{Ashe, "Ashe", Damage},
	{Baptiste, "Baptiste", Support},
	{Bastion, "Bastion", Damage},
	{Brigitte, "Brigitte", Support},
	{Cassidy, "Cassidy", Damage},
	{DVa, "D.Va", Tank},
	{Doomfist, "Doomfist", Tank},
	{Echo, "Echo", Damage},
	{Genji, "Genji", Damage},
	{Hanzo, "Hanzo", Damage},
	{JunkerQueen, "Junker Queen", Tank},
	{Junkrat, "Junkrat", Damage},
	{Kiriko, "Kiriko", Support},
	{LifeWeaver, "Life Weaver", Support},
	{Lucio, "Lúcio", Support},
	{Mei, "Mei", Damage},
	{Mercy, "Mercy", Support},
	{Moira, "Moira", Support},
	{Orisa, "Orisa", Tank},
	{Pharah, "Pharah", Damage},
	{Ramattra, "Ramattra", Tank},
	{Reaper, "Reaper", Damage},
	{Reinhardt, "Reinhardt", Tank},
	{Roadhog, "Roadhog", Tank},
	{Sigma, "Sigma", Tank},
	{Sojourn, "Sojourn", Damage},
	{Soldier76, "Soldier: 76", Damage},
	{Sombra, "Sombra", Damage},
	{Symmetra, "Symmetra", Damage},
	{Torbjorn, "Torbjörn", Damage},
	{Tracer, "Tracer", Damage},
	{Widowmaker, "Widowmaker", Damage},
	{Winston, "Winston", Tank},
	{WreckingBall, "Wrecking Ball", Tank},
	{Zarya, "Zarya", Tank},
	{Zenyatta, "Zenyatta", Support},
}

func GetNameByHeroID(HeroID HeroID) string {
	for _, hero := range HeroList {
		if hero.ID == HeroID {
			return hero.Name
		}
	}
	return ""
}

func GetNameByHeroIDList(HeroIDList []HeroID) []string {
	var heroNameList []string
	for _, heroID := range HeroIDList {
		heroNameList = append(heroNameList, GetNameByHeroID(heroID))
	}
	return heroNameList
}
