package in_memory

import (
	"database/sql"
	cache "github.com/ondi/go-ttl-cache"
	"server/server/logger"
	"time"
)

type InMemoryCacheRepo interface {
	CheckWalletID(ID int64) bool
	UpdateWalletID(ID int64)
	FillingWalletIDs()
}

type InMemoryCache struct {
	balanceDB *sql.DB
	walletIDs *cache.SyncCache_t[int64, struct{}]
	l         logger.Logger
}

func NewInMemoryCacheRepo(bdbr *sql.DB, l logger.Logger) *InMemoryCache {
	walletIDs := cache.NewSync(256*1024, time.Hour, cache.Drop[int64, struct{}])

	inMemoryCache := &InMemoryCache{
		balanceDB: bdbr,
		walletIDs: walletIDs,
		l:         l}

	go inMemoryCache.FillingWalletIDs()

	return inMemoryCache
}

func (c *InMemoryCache) FillingWalletIDs() {
LOOP:
	for {
		rows, err := c.balanceDB.Query("SELECT id FROM wallets")
		if err != nil {
			c.l.ErrorLog.Println("walletIDs cache not inited:", err)
			time.Sleep(5 * time.Minute)
			continue
		}
		for rows.Next() {
			var id int64
			err = rows.Scan(&id)
			if err != nil {
				c.l.ErrorLog.Println("walletIDs cache not inited:", err)
				time.Sleep(5 * time.Minute)
				continue LOOP
			}
			c.walletIDs.Push(
				time.Now(),
				id,
				func() struct{} { return struct{}{} },
				func(e *struct{}) {},
			)
		}
		time.Sleep(30 * time.Minute)
	}
}

func (c *InMemoryCache) CheckWalletID(id int64) bool {
	_, ok := c.walletIDs.Find(time.Now(), id)
	return ok
}

func (c *InMemoryCache) UpdateWalletID(id int64) {
	c.walletIDs.Push(
		time.Now(),
		id,
		func() struct{} { return struct{}{} },
		func(e *struct{}) {},
	)
}
