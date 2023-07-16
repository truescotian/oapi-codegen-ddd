package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/truescotian/oapi-codegen-example/internal/petstore/adapters"
	"github.com/truescotian/oapi-codegen-example/internal/petstore/domain/pet"
)

type application struct {
	petRepo pet.Repository
}

func NewApplication(ctx context.Context) application {
	db, err := NewDatabasePool(ctx, 1)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected")
	}

	pr := adapters.NewPostgresPetRepository(db)

	return application{petRepo: pr}
}

func NewDatabasePool(ctx context.Context, maxConns int) (*pgxpool.Pool, error) {
	if maxConns == 0 {
		maxConns = 1
	}

	connectionString := "postgres://postgres:postgres@localhost:5432/postgres"

	url := fmt.Sprintf(
		"%s?pool_max_conns=%d&pool_min_conns=%d",
		connectionString,
		maxConns,
		2,
	)

	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	// Setting the build statement cache to nil helps this work with pgbouncer
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	config.MaxConnLifetime = 1 * time.Hour
	config.MaxConnIdleTime = 30 * time.Second
	return pgxpool.NewWithConfig(ctx, config)
}
