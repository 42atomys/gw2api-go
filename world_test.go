package gw2api_test

import (
	"testing"

	"atomys.codes/gw2api-go"
)

func TestRequestor_World(t *testing.T) {
	type args struct {
		world gw2api.World
		id    int
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"world with correct id", gw2api.NewRequestor(), args{gw2api.World{}, 2101}, false},
		{"world with invalid id", gw2api.NewRequestor(), args{gw2api.World{}, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.requestor.World(&tt.args.world, tt.args.id).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.World() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestRequestor_WorldIDs(t *testing.T) {
	type args struct {
		worldIDs []int
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"WorldIDs", gw2api.NewRequestor(), args{[]int{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.requestor.WorldIDs(tt.args.worldIDs).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.World() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestRequestor_Worlds(t *testing.T) {
	type args struct {
		ids []int
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"all worlds", gw2api.NewRequestor(), args{nil}, true},
		{"no worlds given", gw2api.NewRequestor(), args{[]int{}}, true},
		{"one or more worlds", gw2api.NewRequestor(), args{[]int{2101}}, false},
		{"one or more invalid ids", gw2api.NewRequestor(), args{[]int{-1, -5}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var worlds = make([]*gw2api.World, 0)
			if got := tt.requestor.Worlds(&worlds, tt.args.ids...).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.World() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}
