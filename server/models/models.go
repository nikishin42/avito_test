package models

import "time"

type Wallet struct {
	ID      int64 `json:"wallet_id" db:"id"`
	Balance int64 `json:"balance" db:"balance"`
	Reserve int64 `json:"reserve" db:"reserve"`
}

type Order struct {
	ID        int64      `json:"order_id" db:"id"`
	WalletID  int64      `json:"wallet_id" db:"wallet_id"`
	ServiceID int64      `json:"service_id" db:"service_id"`
	Price     int64      `json:"price" db:"price"`
	StartDT   time.Time  `json:"start_dt" db:"start_dt"`
	CloseDT   *time.Time `json:"close_dt" db:"close_dt"`
	Canceled  bool       `json:"canceled" db:"canceled"`
}

type Service struct {
	ID   int64 `json:"service_id" db:"id"`
	Name int64 `json:"service_name" db:"name"`
}
