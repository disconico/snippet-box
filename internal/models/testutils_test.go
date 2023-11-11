package models

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"testing"
)

func newTestDB(t *testing.T) *pgxpool.Pool {
	poolConfig, err := pgxpool.ParseConfig("postgres://test_user:password@localhost/test_snippetbox?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		t.Fatal(err)
	}

	script, err := os.ReadFile("./mocks/testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(context.Background(), string(script))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		script, err := os.ReadFile("./mocks/testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}

		_, err = db.Exec(context.Background(), string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close()
	})

	return db
}
