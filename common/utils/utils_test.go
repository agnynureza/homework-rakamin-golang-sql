package utils

import (
	"testing"
)

func TestGenerateNewAccessToken(t *testing.T) {
	// os.Setenv("secret", "supersecret")
	// os.Setenv("minutesCount", "1800")

	// tests := []struct {
	// 	name    string
	// 	want    string
	// 	wantErr bool
	// }{
	// 	{
	// 		name:    "generate token jwt",
	// 		want:    "",
	// 		wantErr: false,
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		got, err := GenerateNewAccessToken()
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("GenerateNewAccessToken() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if got != tt.want {
	// 			t.Errorf("GenerateNewAccessToken() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
