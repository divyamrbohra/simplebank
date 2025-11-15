package db

import (
	"context"
	"log"
	"os"
	"simplebank/util"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestDummy(t *testing.T) {
	t.Log("🚀 Dummy test running")
}

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load Config:", err)
	}

	// Correctly using pgxpool.Connect to create a connection pool
	testDB, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Printf("❌ failed to connect to DB: %v", err)
		log.Fatal("cannot connect to db:", err)
		os.Exit(1)
	}

	testQueries = New(testDB)

	log.Println("✅ Connected to DB successfully")
	code := m.Run()
	testDB.Close()
	os.Exit(code)
}
