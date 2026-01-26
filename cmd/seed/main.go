package main

import (
	"flag"
	"native-setup/config"
	"native-setup/internal/infra/database"
	"native-setup/pkg/env"
	"strconv"

	"github.com/sirupsen/logrus"
)

//
//
// Execute Seeds example:
// go run ./cmd/seed --count=10 
//
//

func main() {
	env.Load()
	dbLogger := database.NewLogger()

	env, err := config.NewEnvs()
	if err != nil {
		logrus.Warnf("Seed failed: %v", err)
	}

	db, err := database.GetStandalone(env, dbLogger)
	if err != nil {
		logrus.Warnf("Seed failed: %v", err)
	}

	countFlag := flag.String("count", "1", "specify the count")
	flag.Parse()

	count, err := strconv.Atoi(*countFlag)
	if err != nil {
		logrus.Warnf("Invalid count: %v", err)
	}

	if err := InitSeeds(db, Options{Count: count}); err != nil {
		logrus.Warnf("Seed failed: %v", err)
	}
}
