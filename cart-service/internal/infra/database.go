package infra

import (
	"cart-service/internal/config"
	"context"
	"database/sql"
	"log/slog"
	"time"

	"github.com/CROWNIX/go-utils/databases/sqlx"
	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	RDBMS sqlx.RDBMS
	Tx    sqlx.Tx
	Sq    squirrel.StatementBuilderType
}

func ProvideTx(db *DB) sqlx.Tx {
	return db.Tx
}

func NewMysql() (*DB, func()) {
	conf := config.GetConfig().DB
	db, err := sql.Open("mysql", conf.DSN)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxIdleTime(conf.ConnMaxIdleTime)
	db.SetConnMaxLifetime(conf.ConnMaxLifetime)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	rdbms := sqlx.NewRDBMS(db,
		sqlx.UseDebug(true),
	)

	slog.Info("initialization mysql successfully")
	return &DB{
			RDBMS: rdbms,
			Tx:    rdbms,
			Sq:    squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question),
		}, func() {
			slog.Info("starting close mysql")

			errClosed := db.Close()
			if errClosed != nil {
				slog.Error("close mysql failed", "error", errClosed)
			}

			slog.Info("close mysql successfully")
		}
}
