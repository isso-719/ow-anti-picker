package owantipicker

import (
	"github.com/isso-719/ow-anti-picker/owhero"
	"reflect"
	"testing"
)

func TestGetAntiPickMap(t *testing.T) {
	type args struct {
		enemyHeroList []owhero.HeroID
	}
	tests := []struct {
		name string
		args args
		want AntiPickMap
	}{
		{
			name: "TestGetAntiPickMap",
			args: args{
				enemyHeroList: []owhero.HeroID{
					owhero.Winston,
					owhero.Tracer,
					owhero.Sombra,
					owhero.Kiriko,
					owhero.Lucio,
				},
			},
			want: AntiPickMap{
				owhero.Reaper:   1,
				owhero.Bastion:  1,
				owhero.Ana:      1,
				owhero.Junkrat:  1,
				owhero.Winston:  3,
				owhero.Pharah:   3,
				owhero.Brigitte: 2,
				owhero.Torbjorn: 2,
				owhero.Roadhog:  2,
				owhero.Sombra:   1,
				owhero.Tracer:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAntiPickMap(tt.args.enemyHeroList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAntiPickMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAntiPickMap_GetAntiPickMapToSortList(t *testing.T) {
	tests := []struct {
		name string
		apm  AntiPickMap
		want AntiPickList
	}{
		{
			name: "TestAntiPickMap_GetAntiPickMapToSortList",
			apm: AntiPickMap{
				owhero.Reaper:  1,
				owhero.Bastion: 5,
				owhero.Ana:     4,
				owhero.Junkrat: 2,
				owhero.Winston: 3,
			},
			want: AntiPickList{
				owhero.Bastion,
				owhero.Ana,
				owhero.Winston,
				owhero.Junkrat,
				owhero.Reaper,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.apm.GetAntiPickMapToSortList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAntiPickMapToSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
