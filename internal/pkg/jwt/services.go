package jwt

import (
	"github.com/ArseniySavin/auth-small-server/config"
	"github.com/ArseniySavin/auth-small-server/typos"
	catcher "github.com/ArseniySavin/catcher/pkg"
)

// Jwt -
type IJwtServices interface {
	Make(claims map[string]interface{}) (string, error)
	Parse(rawJwt string) (map[string]interface{}, error)
}

// DefaultJwtServices -
type DefaultJwtServices struct {
	secret string
}

// NewJwtServices -
func NewJwtServices(cfg *config.Config) IJwtServices {
	if cfg.JWT_SECRET == "" {
		catcher.LogFatal(typos.ErrEmptySecret)
	}

	return &DefaultJwtServices{
		secret: cfg.JWT_SECRET,
	}
}
