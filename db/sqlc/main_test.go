package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/gdygd/simplebank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("Cannot connectd db..", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())

}
