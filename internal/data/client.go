package data

import (
	"context"
	"database/sql"
	"testing"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-redis/redismock/v9"
	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"

	"glintfed.org/ent"
	"glintfed.org/ent/enttest"
)

type Client struct {
	DB  *sql.DB
	Ent *ent.Client

	RDB     *redis.Client
	RDBMock redismock.ClientMock
}

func NewClient(cfg Config) (c *Client, cleanup func(), err error) {
	c = &Client{}

	ctx := context.Background()

	if err = c.initSQLClient(ctx, cfg); err != nil {
		return
	}

	if err = c.initRedisClient(ctx, cfg); err != nil {
		return
	}

	return
}

func NewTestClient(t *testing.T) (c *Client, cleanup func(), err error) {
	c = &Client{
		Ent: enttest.Open(t,
			"sqlite3", "file:ent?mode=memory&_fk=1",
			enttest.WithOptions(ent.Log(t.Log)),
		),
	}

	c.RDB, c.RDBMock = redismock.NewClientMock()

	return
}

func (c *Client) initSQLClient(ctx context.Context, cfg Config) (err error) {
	if c.DB, err = otelsql.Open(
		cfg.Service.Database.SQL.Driver, cfg.Service.Database.SQL.DSN,
		otelsql.WithDBSystem(cfg.Service.Database.SQL.Driver),
		otelsql.WithDBName("glintfed"),
	); err != nil {
		return
	}

	c.Ent = ent.NewClient(ent.Driver(entsql.OpenDB("sqlite3", c.DB)))
	return c.Ent.Schema.Create(ctx)
}

func (c *Client) initRedisClient(_ context.Context, cfg Config) (err error) {
	c.RDB = redis.NewClient(&redis.Options{
		Addr:     cfg.Service.Database.Redis.Addr,
		Password: cfg.Service.Database.Redis.Password,
	})

	return redisotel.InstrumentTracing(c.RDB)
}
