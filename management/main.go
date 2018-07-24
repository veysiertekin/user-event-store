package management

import (
	"github.com/damnever/cc"
)

var Configuration *cc.Config

func init() {
	var err error
	Configuration, err = cc.NewConfigFromFile("app-config.json")
	if err != nil {
		panic(err)
	}
}
