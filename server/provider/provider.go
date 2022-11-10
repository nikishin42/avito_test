package provider

import (
	"database/sql"
	_ "github.com/lib/pq"
	"server/server/config"
	"server/server/logger"
	"server/server/provider/db"
	"server/server/repos/in-memory"
)

type Provider interface {
	GetBalanceDBConn() *sql.DB
	GetInMemoryCache() *in_memory.InMemoryCache
	Close()
}

type provider struct {
	balanceDB *sql.DB
	cache     *in_memory.InMemoryCache
	l         logger.Logger
}

func New(conf *config.Config, l logger.Logger) Provider {
	balanceDB, err := db.Connect(conf)
	if err != nil {
		l.ErrorLog.Fatalln(err)
	}
	l.InfoLog.Println("balance db connected")
	cache := in_memory.NewInMemoryCacheRepo(balanceDB, l)

	return &provider{
		balanceDB: balanceDB,
		cache:     cache,
		l:         l,
	}
}

func (p *provider) GetBalanceDBConn() *sql.DB {
	return p.balanceDB
}

func (p *provider) Close() {
	err := p.balanceDB.Close()
	if err != nil {
		p.l.ErrorLog.Println(err)
	} else {
		p.l.InfoLog.Println("balance db closed")
	}
}

func (p *provider) GetInMemoryCache() *in_memory.InMemoryCache {
	return p.cache
}
