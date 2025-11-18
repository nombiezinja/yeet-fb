package config

type configVars struct {
	AppName          string
	Environment      string
	JwtSigningSecret string
	LogLevel         string
	Port             string
	DatabaseURL      string
}
