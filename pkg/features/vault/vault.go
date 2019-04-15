package vault

import (
	"github.com/hashicorp/vault/api"
)

type authConfig struct {
	Token     string `yaml:"token"`
	VaultAddr string `yaml:"vault_addr"`
}

func loadConfig() (authConfig, err) {
	confi
	return nil
}

func test() {
	client, err := api.NewClient(&api.Config{
		Address: vaultAddress,
	})
}
