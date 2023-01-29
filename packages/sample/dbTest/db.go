package main

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func connectDatabase(ctx context.Context, url string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}
	return conn, nil
}

type name struct {
	id   int
	name string
}

func getNames(ctx context.Context, db *pgx.Conn) ([]*name, error) {
    sql := "SELECT id, name FROM names;"
    rows, err := db.Query(ctx, sql)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    names := make([]*name, 0)
    for rows.Next() {
        n := &name{}
        if err := rows.Scan(&n.id, &n.name); err != nil {
            return nil, err
        }
        names = append(names, n)
    }
    return names, nil
}
