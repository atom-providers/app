package app

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

const DefaultPrefix = "App"

func DefaultProvider() container.ProviderContainer {
	return container.ProviderContainer{
		Provider: Provide,
		Options: []opt.Option{
			opt.Prefix(DefaultPrefix),
		},
	}
}

// swagger:enum AppMode
// ENUM(development, release, test, product)
type AppMode string

type Config struct {
	Mode AppMode
	Cert *Cert
}
type Cert struct {
	CA   string
	Cert string
	Key  string
}
