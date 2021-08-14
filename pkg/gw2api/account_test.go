package gw2api_test

import (
	"os"
	"testing"

	"atomys.codes/gw2api-go/pkg/gw2api"
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
