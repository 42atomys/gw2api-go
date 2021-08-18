package gw2api_test

import (
	"testing"

	"atomys.codes/gw2api-go/pkg/gw2api"
)

func TestRequestor_WorldBoss(t *testing.T) {
	type args struct {
		worldBoss gw2api.WorldBoss
		id        string
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"worldBoss with correct id", gw2api.NewRequestor(), args{gw2api.WorldBoss{}, "the_shatterer"}, false},
		{"worldBoss with invalid id", gw2api.NewRequestor(), args{gw2api.WorldBoss{}, "a_fuckin_barrel"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.requestor.WorldBoss(&tt.args.worldBoss, tt.args.id).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.World() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestRequestor_WorldBosses(t *testing.T) {
	type args struct {
		worldBosses []string
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"WorldBosses", gw2api.NewRequestor(), args{[]string{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.requestor.WorldBosses(&tt.args.worldBosses).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.World() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}
