package owhero

import (
	"reflect"
	"testing"
)

func TestGetAntiPicksByHeroID(t *testing.T) {
	type args struct {
		HeroID HeroID
	}
	tests := []struct {
		name string
		args args
		want []HeroID
	}{
		{
			name: "TestGetAntiPicksByHeroID",
			args: args{
				HeroID: Reaper,
			},
			want: []HeroID{
				Widowmaker,
				Hanzo,
				Pharah,
				Echo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAntiPicksByHeroID(tt.args.HeroID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAntiPicksByHeroID() = %v, want %v", got, tt.want)
			}
		})
	}
}
