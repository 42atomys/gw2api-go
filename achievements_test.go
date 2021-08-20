package gw2api_test

import (
	"testing"

	"atomys.codes/gw2api-go"
)

func TestRequestor_Achievement(t *testing.T) {
	type args struct {
		achievement gw2api.Achievement
		id          int
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"achievement with correct id", gw2api.NewRequestor(), args{gw2api.Achievement{}, 1}, false},
		{"achievement with invalid id", gw2api.NewRequestor(), args{gw2api.Achievement{}, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.requestor.Achievement(&tt.args.achievement, tt.args.id).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.Achievement() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestRequestor_AchievementIDs(t *testing.T) {
	type args struct {
		achievementIDs []int
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"AchievementIDs", gw2api.NewRequestor(), args{[]int{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.requestor.AchievementIDs(&tt.args.achievementIDs).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.Achievement() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestRequestor_Achievements(t *testing.T) {
	type args struct {
		ids []int
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"all achievements", gw2api.NewRequestor(), args{nil}, true},
		{"no achievements given", gw2api.NewRequestor(), args{[]int{}}, true},
		{"one or more achievements", gw2api.NewRequestor(), args{[]int{1}}, false},
		{"one or more invalid ids", gw2api.NewRequestor(), args{[]int{-1, -5}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var achievements = make([]*gw2api.Achievement, 0)
			if got := tt.requestor.Achievements(&achievements, tt.args.ids...).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.Achievement() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestRequestor_AchievementsCategory(t *testing.T) {
	type args struct {
		achievementsCategory gw2api.AchievementsCategory
		id                   int
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"achievementsCategory with correct id", gw2api.NewRequestor(), args{gw2api.AchievementsCategory{}, 1}, false},
		{"achievementsCategory with invalid id", gw2api.NewRequestor(), args{gw2api.AchievementsCategory{}, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.requestor.AchievementsCategory(&tt.args.achievementsCategory, tt.args.id).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AchievementsCategory() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestRequestor_AchievementsCategoryIDs(t *testing.T) {
	type args struct {
		achievementsCategoryIDs []int
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"AchievementsCategoryIDs", gw2api.NewRequestor(), args{[]int{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.requestor.AchievementsCategoryIDs(&tt.args.achievementsCategoryIDs).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AchievementsCategory() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestRequestor_AchievementsCategorys(t *testing.T) {
	type args struct {
		ids []int
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"all achievementsCategorys", gw2api.NewRequestor(), args{nil}, true},
		{"no achievementsCategorys given", gw2api.NewRequestor(), args{[]int{}}, true},
		{"one or more achievementsCategorys", gw2api.NewRequestor(), args{[]int{1}}, false},
		{"one or more invalid ids", gw2api.NewRequestor(), args{[]int{-1, -5}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var achievementsCategorys = make([]*gw2api.AchievementsCategory, 0)
			if got := tt.requestor.AchievementsCategories(&achievementsCategorys, tt.args.ids...).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AchievementsCategories() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestRequestor_AchievementsDaily(t *testing.T) {
	type args struct {
		achievementDaily gw2api.AchievementsDailyStructure
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"AchievementsDaily", gw2api.NewRequestor(), args{gw2api.AchievementsDailyStructure{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.requestor.AchievementsDaily(&tt.args.achievementDaily).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AchievementsDaily() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestRequestor_AchievementsDailyTomorrow(t *testing.T) {
	type args struct {
		achievementDaily gw2api.AchievementsDailyStructure
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"AchievementsDailyTomorrow", gw2api.NewRequestor(), args{gw2api.AchievementsDailyStructure{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.requestor.AchievementsDailyTomorrow(&tt.args.achievementDaily).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AchievementsDailyTomorrow() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestRequestor_AchievementsGroup(t *testing.T) {
	type args struct {
		achievementsGroup gw2api.AchievementsGroup
		id                string
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"achievementsGroup with correct id", gw2api.NewRequestor(), args{gw2api.AchievementsGroup{}, "A4ED8379-5B6B-4ECC-B6E1-70C350C902D2"}, false},
		{"achievementsGroup with invalid id", gw2api.NewRequestor(), args{gw2api.AchievementsGroup{}, "a"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.requestor.AchievementsGroup(&tt.args.achievementsGroup, tt.args.id).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AchievementsGroup() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestRequestor_AchievementsGroupIDs(t *testing.T) {
	type args struct {
		achievementsGroupIDs []string
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"AchievementsGroupIDs", gw2api.NewRequestor(), args{[]string{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.requestor.AchievementsGroupIDs(tt.args.achievementsGroupIDs).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AchievementsGroup() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestRequestor_AchievementsGroups(t *testing.T) {
	type args struct {
		ids []string
	}
	tests := []struct {
		name      string
		requestor *gw2api.Requestor
		args      args
		wantErr   bool
	}{
		{"all achievementsGroups", gw2api.NewRequestor(), args{nil}, true},
		{"no achievementsGroups given", gw2api.NewRequestor(), args{[]string{}}, true},
		{"one or more achievementsGroups", gw2api.NewRequestor(), args{[]string{"A4ED8379-5B6B-4ECC-B6E1-70C350C902D2"}}, false},
		{"one or more invalid ids", gw2api.NewRequestor(), args{[]string{"-1", "-5"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var achievementsGroups = make([]*gw2api.AchievementsGroup, 0)
			if got := tt.requestor.AchievementsGroups(&achievementsGroups, tt.args.ids...).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AchievementsGroup() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}
