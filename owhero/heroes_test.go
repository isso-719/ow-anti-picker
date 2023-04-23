package owhero

import (
	"reflect"
	"testing"
)

func TestGetNameByHeroID(t *testing.T) {
	type args struct {
		HeroID HeroID
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestGetNameByHeroID",
			args: args{
				HeroID: Reaper,
			},
			want: "Reaper",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNameByHeroID(tt.args.HeroID); got != tt.want {
				t.Errorf("GetNameByHeroID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNameByHeroIDList(t *testing.T) {
	type args struct {
		HeroIDList []HeroID
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetNameByHeroIDList",
			args: args{
				HeroIDList: []HeroID{
					Reaper,
					Bastion,
					Ana,
				},
			},
			want: []string{
				"Reaper",
				"Bastion",
				"Ana",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNameByHeroIDList(tt.args.HeroIDList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNameByHeroIDList() = %v, want %v", got, tt.want)
			}
		})
	}
}
