package configuration

import (
	"github.com/damnever/cc"
)

var Data *cc.Config

func init() {
	var err error
	Data, err = cc.NewConfigFromFile("app-dev.json")
	if err != nil {
		panic(err)
	}
}
