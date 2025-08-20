package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	t.Skip() // 確認したい時のみコメントアウト

	setEnv()
	tests := []struct {
		name   string
		params *Params
		isErr  bool
	}{
		{
			name: "success",
			params: &Params{
				Socket:   "tcp",
				Host:     os.Getenv("DB_HOST"),
				Port:     os.Getenv("DB_PORT"),
				Database: os.Getenv("DB_DATABASE"),
				Username: os.Getenv("DB_USERNAME"),
				Password: os.Getenv("DB_PASSWORD"),
			},
			isErr: false,
		},
		{
			name: "failed to connect mysql",
			params: &Params{
				Socket:   "tcp",
				Host:     "127.0.0.1",
				Port:     "80",
				Database: "clean-architecture",
				Username: "root",
				Password: "",
			},
			isErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewDatabaseClient(tt.params)
			if tt.isErr {
				assert.Error(t, err)
				assert.Nil(t, client)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, client)
		})
	}
}

func setEnv() {
	if os.Getenv("DB_HOST") == "" {
		os.Setenv("DB_HOST", "127.0.0.1")
	}
	if os.Getenv("DB_PORT") == "" {
		os.Setenv("DB_PORT", "3326")
	}
	if os.Getenv("DB_USERNAME") == "" {
		os.Setenv("DB_USERNAME", "root")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		os.Setenv("DB_PASSWORD", "password")
	}
	if os.Getenv("DB_DATABASE") == "" {
		os.Setenv("DB_DATABASE", "clean-architecture")
	}
}
