package owhero

type AntiPick struct {
	HeroID    HeroID
	AntiPicks []HeroID
}

// Anti Picks List Defind By Copyright (C) 2015 Lighthouse Studio, Inc. All Rights Reserved.
// Source: https://kamigame.jp/overwatch2/page/236571719085594935.html

var AntiPicks = []AntiPick{
	{DVa, []HeroID{Sombra, Zarya, Roadhog}},
	{Ashe, []HeroID{Widowmaker, DVa, Genji}},
	{Ana, []HeroID{DVa, Sombra, Genji, Kiriko}},
	{Widowmaker, []HeroID{Winston, Sombra, WreckingBall}},
	{Winston, []HeroID{Reaper, Bastion, Ana, Junkrat}},
	{Echo, []HeroID{Widowmaker, Ashe, Sojourn, Sombra}},
	{Orisa, []HeroID{Zarya, Reaper, Junkrat}},
	{Cassidy, []HeroID{Widowmaker, Winston, Tracer}},
	{Kiriko, []HeroID{Pharah, Sombra, Tracer}},
	{Genji, []HeroID{Brigitte, Moira, Pharah, Zarya}},
	{Zarya, []HeroID{Sombra, Zenyatta, Pharah, Mei}},
	{Sigma, []HeroID{Genji, Reaper, Reinhardt}},
	{Symmetra, []HeroID{Pharah, Widowmaker, Winston, DVa}},
	{JunkerQueen, []HeroID{Orisa, Zenyatta, Echo, Mei}},
	{Junkrat, []HeroID{Pharah, Echo, Widowmaker, Sojourn}},
	{Zenyatta, []HeroID{Winston, Widowmaker, Genji, Sombra}},
	{Sojourn, []HeroID{Zarya, Widowmaker, DVa, Tracer}},
	{Soldier76, []HeroID{Genji, Winston, Roadhog, Sigma}},
	{Sombra, []HeroID{Brigitte, Roadhog, Winston}},
	{Tracer, []HeroID{Winston, Pharah, Brigitte, Torbjorn}},
	{Doomfist, []HeroID{Reaper, Pharah, Echo, Ana}},
	{Hanzo, []HeroID{Widowmaker, Winston, DVa, Genji}},
	{Bastion, []HeroID{Ana, Roadhog, Reaper, Genji}},
	{Baptiste, []HeroID{DVa, Roadhog, Genji}},
	{Pharah, []HeroID{Widowmaker, Hanzo, Soldier76, Ashe}},
	{Brigitte, []HeroID{Pharah, Hanzo, Sojourn, Symmetra}},
	{Mercy, []HeroID{Winston, Roadhog, Sombra}},
	{Mei, []HeroID{Sombra, Pharah, Widowmaker}},
	{Moira, []HeroID{Widowmaker, Pharah, DVa}},
	{Reinhardt, []HeroID{Junkrat, Bastion, Reaper, Echo}},
	{Reaper, []HeroID{Widowmaker, Hanzo, Pharah, Echo}},
	{Lucio, []HeroID{Pharah, Winston, Roadhog, Torbjorn}},
	{WreckingBall, []HeroID{Sombra, Junkrat, Bastion}},
	{Roadhog, []HeroID{Ana, Reaper, Pharah, Zenyatta}},
}

func GetAntiPicksByHeroID(HeroID HeroID) []HeroID {
	for _, antiPick := range AntiPicks {
		if antiPick.HeroID == HeroID {
			return antiPick.AntiPicks
		}
	}
	return nil
}
