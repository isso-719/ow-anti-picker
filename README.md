# OverWatch Anti Picker

## What is this?

This library is used when you want to do anti-pick with OverWatch2.

## How to use?

### Install

```bash
go get github.com/isso-719/ow-anti-picker
```

### Usage

```go
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
	fmt.Printf("AntiPickMap (HeroID:Important): %v\n", antiPickMap)

	antiPickList := antiPickMap.GetAntiPickMapToSortByImportantListWithHeroName()
	fmt.Printf("AntiPickList (HeroName): %v\n", antiPickList)
}
```

### Result

```text
AntiPickMap (HeroID: Important): map[0:1 3:1 4:2 12:1 20:3 22:1 24:2 28:1 30:2 31:1 33:3]
AntiPickList (HeroName): [Winston Pharah Brigitte Roadhog Torbjörn Ana Junkrat Sombra Tracer Reaper Bastion]
```
