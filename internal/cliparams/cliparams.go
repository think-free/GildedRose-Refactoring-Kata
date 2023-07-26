package cliparams

import (
	"github.com/kelseyhightower/envconfig"
)

type ClientParameters struct {
	Days     int    `envconfig:"days" required:"false" default:"2"`
	LogLevel string `envconfig:"loglevel" required:"false" default:"info"`
}

func New() *ClientParameters {
	cp := &ClientParameters{}
	envconfig.MustProcess("", cp)
	cp.Days++
	return cp
}
