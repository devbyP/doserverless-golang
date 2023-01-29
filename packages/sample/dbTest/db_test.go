package main

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/require"
)

var (
	db  *pgx.Conn
	ctx context.Context
    testDb = "postgres://ab:E0o5YNLXndpJ@ep-blue-king-455349.ap-southeast-1.aws.neon.tech/neondb"
)

func TestMain(m *testing.M) {
	var err error
	ctx = context.Background()
	db, err = connectDatabase(ctx, testDb)
	if err != nil {
		return
	}
	m.Run()
}

func TestGetName(t *testing.T) {
	names, err := getNames(ctx, db)
	require.Empty(t, err)
	require.Len(t, names, 3)
}
