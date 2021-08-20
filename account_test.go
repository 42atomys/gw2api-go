package gw2api_test

import (
	"os"
	"testing"

	"atomys.codes/gw2api-go"
)

func TestAccount(t *testing.T) {
	tests := []struct {
		name    string
		apiKey  string
		wantErr bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if tt.apiKey == "" {
				t.Skip("Requestor.Account() = Cannot test without api key")
			}

			var account gw2api.Account
			if got := r.Auth(tt.apiKey).Account(&account).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.Account() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountWithoutAuth(t *testing.T) {
	var account gw2api.Account
	err := gw2api.NewRequestor().Account(&account).Err()
	if err != gw2api.ErrRequireAuthentication {
		t.Errorf("Requestor.Account() = %v, want %v", err, gw2api.ErrRequireAuthentication)
	}
}

func TestAccountAchievements(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountAchievements() = Cannot test without api key")
			}

			var achievements []*gw2api.AccountAchievement
			if got := r.Auth(tt.apiKey).AccountAchievements(&achievements).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountAchievements() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountBank(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountBank() = Cannot test without api key")
			}

			var accountBank []*gw2api.InventoryItem
			if got := r.Auth(tt.apiKey).AccountBank(&accountBank).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountBank() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountBuildStorageIDs(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_INVALID"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountBuildStorageIDs() = Cannot test without api key")
			}

			var accountBuildStorage []int
			if got := r.Auth(tt.apiKey).AccountBuildStorageIDs(&accountBuildStorage).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountBuildStorageIDs() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountBuildStorages(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_INVALID"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountBuildStorages() = Cannot test without api key")
			}

			var accountBuildStorage []*gw2api.AccountBuildStorage
			if got := r.Auth(tt.apiKey).AccountBuildStorages(&accountBuildStorage, 1, 2).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountBuildStorages() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountBuildStorage(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_INVALID"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountBuildStorage() = Cannot test without api key")
			}

			var accountBuildStorage gw2api.AccountBuildStorage
			if got := r.Auth(tt.apiKey).AccountBuildStorage(&accountBuildStorage, 1).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountBuildStorage() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountDailyCrafting(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountDailyCrafting() = Cannot test without api key")
			}

			var accountDailyCraft []string
			if got := r.Auth(tt.apiKey).AccountDailyCrafting(&accountDailyCraft).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountDailyCrafting() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountDungeons(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountDungeons() = Cannot test without api key")
			}

			var accountDungeons []string
			if got := r.Auth(tt.apiKey).AccountDungeons(&accountDungeons).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountDungeons() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountDyes(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountDyes() = Cannot test without api key")
			}

			var accountDyes []int
			if got := r.Auth(tt.apiKey).AccountDyes(&accountDyes).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountDyes() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountEmotes(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountEmotes() = Cannot test without api key")
			}

			var result []string
			if got := r.Auth(tt.apiKey).AccountEmotes(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountEmotes() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountFinishers(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountFinishers() = Cannot test without api key")
			}

			var result []*gw2api.AccountFinisher
			if got := r.Auth(tt.apiKey).AccountFinishers(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountFinishers() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountGliders(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountGliders() = Cannot test without api key")
			}

			var result []int
			if got := r.Auth(tt.apiKey).AccountGliders(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountGliders() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountHome(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_INVALID"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountHome() = Cannot test without api key")
			}

			var result []string
			if got := r.Auth(tt.apiKey).AccountHome(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountHome() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountHomeCats(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountHomeCats() = Cannot test without api key")
			}

			var result []int
			if got := r.Auth(tt.apiKey).AccountHomeCats(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountHomeCats() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountHomeNodes(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountHomeNodes() = Cannot test without api key")
			}

			var result []string
			if got := r.Auth(tt.apiKey).AccountHomeNodes(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountHomeNodes() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountInventory(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountInventory() = Cannot test without api key")
			}

			var result []*gw2api.InventoryItem
			if got := r.Auth(tt.apiKey).AccountInventory(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountInventory() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountLegendaryArmory(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountLegendaryArmory() = Cannot test without api key")
			}

			var result []*gw2api.AccountLegendaryArmory
			if got := r.Auth(tt.apiKey).AccountLegendaryArmory(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountLegendaryArmory() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountLuck(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountLuck() = Cannot test without api key")
			}

			var result []*gw2api.AccountLuck
			if got := r.Auth(tt.apiKey).AccountLuck(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountLuck() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountMailCarriers(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountMailCarriers() = Cannot test without api key")
			}

			var result []int
			if got := r.Auth(tt.apiKey).AccountMailCarriers(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountMailCarriers() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountMapChests(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountMapChests() = Cannot test without api key")
			}

			var result []string
			if got := r.Auth(tt.apiKey).AccountMapChests(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountMapChests() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountMasteries(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountMasteries() = Cannot test without api key")
			}

			var result []*gw2api.AccountMastery
			if got := r.Auth(tt.apiKey).AccountMasteries(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountMasteries() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountMasteryPoints(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountMasteryPoints() = Cannot test without api key")
			}

			var result gw2api.AccountMasteryPoint
			if got := r.Auth(tt.apiKey).AccountMasteryPoints(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountMasteryPoints() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountMaterials(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountMaterials() = Cannot test without api key")
			}

			var result []*gw2api.AccountMaterial
			if got := r.Auth(tt.apiKey).AccountMaterials(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountMaterials() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountMinis(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountMinis() = Cannot test without api key")
			}

			var result []int
			if got := r.Auth(tt.apiKey).AccountMinis(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountMinis() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountMounts(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_INVALID"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountMounts() = Cannot test without api key")
			}

			var result []string
			if got := r.Auth(tt.apiKey).AccountMounts(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountMounts() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountMountsSkins(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountMountsSkins() = Cannot test without api key")
			}

			var result []int
			if got := r.Auth(tt.apiKey).AccountMountsSkins(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountMountsSkins() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountMountsTypes(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountMountsTypes() = Cannot test without api key")
			}

			var result []string
			if got := r.Auth(tt.apiKey).AccountMountsTypes(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountMountsTypes() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountNovelties(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountNovelties() = Cannot test without api key")
			}

			var result []int
			if got := r.Auth(tt.apiKey).AccountNovelties(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountNovelties() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountOutfils(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountOutfils() = Cannot test without api key")
			}

			var result []int
			if got := r.Auth(tt.apiKey).AccountOutfils(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountOutfils() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountPvpHeroes(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountPvpHeroes() = Cannot test without api key")
			}

			var result []int
			if got := r.Auth(tt.apiKey).AccountPvpHeroes(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountPvpHeroes() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountRaids(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountRaids() = Cannot test without api key")
			}

			var result []string
			if got := r.Auth(tt.apiKey).AccountRaids(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountRaids() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountReceipes(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountReceipes() = Cannot test without api key")
			}

			var result []int
			if got := r.Auth(tt.apiKey).AccountReceipes(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountReceipes() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountSkins(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountSkins() = Cannot test without api key")
			}

			var result []int
			if got := r.Auth(tt.apiKey).AccountSkins(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountSkins() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountTitles(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountTitles() = Cannot test without api key")
			}

			var result []int
			if got := r.Auth(tt.apiKey).AccountTitles(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountTitles() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountWallet(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountWallet() = Cannot test without api key")
			}

			var result []*gw2api.AccountCurrency
			if got := r.Auth(tt.apiKey).AccountWallet(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountWallet() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}

func TestAccountWorldBosses(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		withoutApi bool
		wantErr    bool
	}{
		{"with valid api_key", os.Getenv("API_KEY_VALID"), false, false},
		{"with invalid api_key", os.Getenv("API_KEY_MISSING_SCOPE"), false, true},
		{"without api_key", "", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gw2api.NewRequestor()

			if !tt.withoutApi && tt.apiKey == "" {
				t.Skip("Requestor.AccountWorldBosses() = Cannot test without api key")
			}

			var result []string
			if got := r.Auth(tt.apiKey).AccountWorldBosses(&result).Err(); (got != nil) != tt.wantErr {
				t.Errorf("Requestor.AccountWorldBosses() = %v, want error %v", got, tt.wantErr)
			}
		})
	}
}
