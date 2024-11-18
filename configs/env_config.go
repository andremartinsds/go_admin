package configs

import (
	"github.com/lestrrat-go/jwx/v2/jwt"
	"strconv"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
)

type Conf struct {
	JWTSecret    string
	JWTExpiresIn int
	TokenAuth    *jwtauth.JWTAuth
}

var Config *Conf

func LoadEnvConfig() error {
	var cfg *Conf

	viper.SetConfigName(".env")
	// viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../..")
	err := viper.ReadInConfig()
	expiresIn, _ := strconv.Atoi(viper.GetString("JWT_EXPIRES"))
	cfg = &Conf{
		JWTSecret:    viper.GetString("JWT_SECRET"),
		JWTExpiresIn: expiresIn,
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil, jwt.WithAcceptableSkew(24*time.Hour*100))

	if err != nil {
		return err
	}
	Config = cfg
	return nil
}
