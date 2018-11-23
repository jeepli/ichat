package database

import (
	"github.com/jeepli/ichat/config"
	pg "gopkg.in/pg.v5"
)

// type DBAdapter interface {
// }

type DBHolder struct {
	*pg.DB
}

func NewDBHolder(c *config.DbConfig) *DBHolder {
	db := pg.Connect(&pg.Options{
		Addr:     c.Address,
		User:     c.User,
		Password: c.Password,
		Database: c.Database,
		// MaxRetries:            3,
		// RetryStatementTimeout: true,
		// DialTimeout:           5 * time.Second,
		// ReadTimeout:           5 * time.Second,
		// WriteTimeout:          5 * time.Second,
		// PoolSize:           10,
		// PoolTimeout:        5 * time.Second,
		// IdleTimeout:        5 * time.Second,
		// MaxAge:             5 * time.Second,
		// IdleCheckFrequency: 5 * time.Second,
		// DisableTransaction: false,
	})

	return &DBHolder{DB: db}
}
