package owantipicker

import (
	"github.com/isso-719/ow-anti-picker/owhero"
	"sort"
)

type AntiPickMap map[owhero.HeroID]antiPickRevealCount
type antiPickRevealCount int

// GetAntiPickMap : generates a map of anti-picks and their reveal count
// Input: EnemyHeroList (5 Heroes)
// Output: AntiPickList: map[HeroID]Count
func GetAntiPickMap(enemyHeroList []owhero.HeroID) AntiPickMap {
	var antiPickMap = make(map[owhero.HeroID]antiPickRevealCount)

	for _, enemyHeroID := range enemyHeroList {
		antiPicks := owhero.GetAntiPicksByHeroID(enemyHeroID)

		for _, antiPick := range antiPicks {
			if antiPickMap[antiPick] == 0 {
				antiPickMap[antiPick] = 1
			} else {
				antiPickMap[antiPick]++
			}
		}
	}
	return antiPickMap
}

type AntiPickList []owhero.HeroID

// GetAntiPickMapToSortByImportantList : converts AntiPickMap to array of HeroIDs sorted by reveal count (important hero first)
func (apm AntiPickMap) GetAntiPickMapToSortByImportantList() AntiPickList {
	var antiPickHeroIDs []owhero.HeroID
	for k := range apm {
		antiPickHeroIDs = append(antiPickHeroIDs, k)
	}
	sort.Slice(antiPickHeroIDs, func(i, j int) bool {
		return apm[antiPickHeroIDs[i]] > apm[antiPickHeroIDs[j]]
	})

	return antiPickHeroIDs
}

type AntiPickListWithHeroName []string

func (apm AntiPickMap) GetAntiPickMapToSortByImportantListWithHeroName() AntiPickListWithHeroName {
	var antiPickHeroIDs []string

	antiPickMapToSortByImportantList := apm.GetAntiPickMapToSortByImportantList()

	for _, antiPickHeroID := range antiPickMapToSortByImportantList {
		antiPickHeroIDs = append(antiPickHeroIDs, owhero.GetNameByHeroID(antiPickHeroID))
	}

	return antiPickHeroIDs
}
