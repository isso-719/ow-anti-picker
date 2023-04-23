package main

import (
	"fmt"
	"github.com/isso-719/ow-anti-picker"
	"github.com/isso-719/ow-anti-picker/owhero"
)

func main() {
	enemyHeroes := []owhero.HeroID{
		owhero.Winston,
		owhero.Tracer,
		owhero.Sombra,
		owhero.Kiriko,
		owhero.Lucio,
	}
	antiPickMap := owantipicker.GetAntiPickMap(enemyHeroes)
	fmt.Printf("AntiPickMap (HeroID: Important): %v\n", antiPickMap)

	antiPickList := antiPickMap.GetAntiPickMapToSortByImportantListWithHeroName()
	fmt.Printf("AntiPickList (HeroName): %v\n", antiPickList)
}
