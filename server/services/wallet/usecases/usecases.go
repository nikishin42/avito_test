package usecases

import (
	"errors"
	"fmt"
	"server/server/logger"
	"server/server/models"
	"server/server/repos"
	balance_db "server/server/repos/balance-db"
	in_memory "server/server/repos/in-memory"
)

type WalletService struct {
	balanceDB balance_db.BalanceDBRepo
	cache     in_memory.InMemoryCacheRepo
	l         logger.Logger
}

func New(r *repos.Repositories, l logger.Logger) *WalletService {
	return &WalletService{
		balanceDB: r.BalanceDB,
		cache:     r.Cache,
		l:         l,
	}
}

func (w *WalletService) checkWalletID(id int64) bool {
	return w.cache.CheckWalletID(id)
}

func (w *WalletService) TopUpWallet(id, count int64) error {
	ok := w.checkWalletID(id)
	if !ok {
		err := w.balanceDB.CreateWallet(id, count)
		if err != nil {
			return err
		}
		w.cache.UpdateWalletID(id)
		w.l.InfoLog.Println("created wallet, id =", id)
		return nil
	}
	err := w.balanceDB.TopUpWallet(id, count)
	if err != nil {
		return err
	}
	w.l.InfoLog.Println("top upped wallet, id =", id)
	return nil
}

func (w *WalletService) GetBalance(id int64) (models.Wallet, error) {
	ok := w.checkWalletID(id)
	if !ok {
		return models.Wallet{}, errors.New(fmt.Sprintf("wallet id %d not found", id))
	}
	return w.balanceDB.GetBalance(id)
}

func (w *WalletService) LockMoney(order models.Order) error {
	ok := w.checkWalletID(order.WalletID)
	if !ok {
		err := errors.New(fmt.Sprintf("wallet id %d not found", order.WalletID))
		w.l.ErrorLog.Println(err)
		return err
	}
	wallet, err := w.GetBalance(order.WalletID)
	if err != nil {
		w.l.ErrorLog.Println(err)
		return err
	}
	if wallet.Balance-order.Price < 0 {
		err = errors.New("not enough money in wallet")
		w.l.InfoLog.Println(err)
		return err
	}
	w.l.InfoLog.Printf("%#v", order)
	return w.balanceDB.LockMoney(order)
}
