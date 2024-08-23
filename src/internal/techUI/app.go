package techUI

import (
	"SmartShopper/internal/config"
	"SmartShopper/internal/techUI/requesters"
)

func Run(configDir string) {
	cfg, err := config.Init(configDir)
	if err != nil {
		return
	}

	requester := requesters.NewRequester(
		cfg.Auth.JWT.AccessTokenTTL,
		cfg.Auth.JWT.RefreshTokenTTL,
		cfg.Port,
	)

	requester.Run()
}
