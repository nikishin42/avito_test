package balance_db

import (
	"database/sql"
	"server/server/logger"
	"server/server/models"
	"time"
)

type BalanceDBRepo interface {
	TopUpWallet(id, count int64) error
	CreateWallet(id, count int64) error
	GetBalance(id int64) (models.Wallet, error)
	LockMoney(order models.Order) error
	TransactionValidation(order models.Order) error
}

type balanceDB struct {
	db *sql.DB
	l  logger.Logger
}

func NewBalanceDBRepo(db *sql.DB, l logger.Logger) BalanceDBRepo {
	return &balanceDB{
		db: db,
		l:  l,
	}
}

func (b *balanceDB) TopUpWallet(id, count int64) error {
	_, err := b.db.Exec("update wallets set balance = balance+$1 where id = $2", count, id)
	return err
}

func (b *balanceDB) CreateWallet(id, count int64) error {
	_, err := b.db.Exec("insert into wallets values ($1, $2, default)", id, count)
	return err
}

func (b *balanceDB) GetBalance(id int64) (models.Wallet, error) {
	row := b.db.QueryRow("select * from wallets where id = $1", id)
	wallet := models.Wallet{}
	err := row.Scan(&wallet.ID, &wallet.Balance, &wallet.Reserve)
	return wallet, err
}

func (b *balanceDB) LockMoney(order models.Order) error {
	_, err := b.db.Exec("insert into orders (id, wallet_id, service_id, price, start_dt, close_dt) values ($1, $2, $3, $4, $5, $6)",
		order.ID, order.WalletID, order.ServiceID, order.Price, time.Now(), order.CloseDT)
	if err != nil {
		b.l.InfoLog.Println("saving new order has been failed:", err)
		return nil
	}
	b.l.InfoLog.Println("new order saved")

	_, err = b.db.Exec("update wallets set balance=balance-$1, reserved=reserved+$1 where id=$2", order.Price, order.WalletID)
	if err == nil {
		b.l.InfoLog.Println("money has been locked")
		return nil
	}
	b.l.ErrorLog.Println("money hasn't been locked:", err)

	_, err = b.db.Exec("delete from orders where id=$1", order.ID)
	if err == nil {
		b.l.InfoLog.Println("new order has been deleted")
		return nil
	}
	b.l.ErrorLog.Println("new order has not deleted")
	return err
}

func (b *balanceDB) TransactionValidation(order models.Order) error {
	//_, err := b.db.Exec("insert into orders (id, wallet_id, service_id, price, start_dt, close_dt) values ($1, $2, $3, $4, $5, $6)",
	//	order.ID, order.WalletID, order.ServiceID, order.Price, time.Now(), order.CloseDT)
	//if err != nil {
	//	b.l.InfoLog.Println("saving new order has been failed:", err)
	//	return nil
	//}
	//b.l.InfoLog.Println("new order saved")
	//
	//_, err = b.db.Exec("update wallets set balance=balance-$1, reserved=reserved+$1 where id=$2", order.Price, order.WalletID)
	//if err == nil {
	//	b.l.InfoLog.Println("money has been locked")
	//	return nil
	//}
	//b.l.ErrorLog.Println("money hasn't been locked:", err)
	//
	//_, err = b.db.Exec("delete from orders where id=$1", order.ID)
	//if err == nil {
	//	b.l.InfoLog.Println("new order has been deleted")
	//	return nil
	//}
	//b.l.ErrorLog.Println("new order has not deleted")
	//return err
	return nil
}
