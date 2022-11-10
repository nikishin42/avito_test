package api

import (
	"github.com/gorilla/mux"
	"server/server/logger"
	"server/server/services/wallet/api/handlers"
	"server/server/services/wallet/usecases"
)

type WalletService struct {
	h *handlers.Handlers
	u *usecases.WalletService
	l logger.Logger
}

func New(u *usecases.WalletService, l logger.Logger) *WalletService {
	return &WalletService{
		h: handlers.New(u, l),
		u: u,
		l: l,
	}
}

func (s *WalletService) Router(r *mux.Router) {
	r.HandleFunc("/top_up", s.h.TopUpWallet).Methods("POST")
	r.HandleFunc("/get_balance", s.h.GetBalance).Methods("GET")
	r.HandleFunc("/lock_money", s.h.LockMoney).Methods("POST")

}

func (s *WalletService) Start() error {
	s.l.InfoLog.Println("Wallet service started")
	return nil
}

func (s *WalletService) Stop() error {
	s.l.InfoLog.Println("Wallet service stopped")
	return nil
}
