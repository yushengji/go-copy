package gocp

import "go.uber.org/zap"

var log *zap.Logger

func init() {
	var err error
	log, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}
