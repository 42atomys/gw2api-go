package gw2api_test

import (
	"testing"

	"atomys.codes/gw2api-go/pkg/gw2api"
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
			r := gw2api.NewRequestor()
			if got := r.World(&tt.args.world, tt.args.id).Err(); tt.wantErr == (got != nil) {
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
			r := gw2api.NewRequestor()
			if got := r.WorldIDs(tt.args.worldIDs).Err(); tt.wantErr == (got != nil) {
				t.Errorf("Requestor.World() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestRequestor_Worlds(t *testing.T) {
	type args struct {
		worlds []*gw2api.World
		ids    []int
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"all worlds", gw2api.NewRequestor(), args{make([]*gw2api.World, 0), nil}, false},
		{"no worlds given", gw2api.NewRequestor(), args{make([]*gw2api.World, 0), []int{}}, true},
		{"one or more worlds", gw2api.NewRequestor(), args{make([]*gw2api.World, 0), []int{2101}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()
			if got := r.Worlds(&tt.args.worlds, tt.args.ids...).Err(); tt.wantErr == (got != nil) {
				t.Errorf("Requestor.World() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}
