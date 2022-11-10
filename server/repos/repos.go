package repos

import (
	"server/server/logger"
	"server/server/provider"
	"server/server/repos/balance-db"
	"server/server/repos/in-memory"
)

type Repositories struct {
	BalanceDB balance_db.BalanceDBRepo
	Cache     in_memory.InMemoryCacheRepo
}

func New(p provider.Provider, l logger.Logger) *Repositories {
	balanceDB := balance_db.NewBalanceDBRepo(p.GetBalanceDBConn(), l)
	cache := p.GetInMemoryCache()
	return &Repositories{
		BalanceDB: balanceDB,
		Cache:     cache,
	}
}
