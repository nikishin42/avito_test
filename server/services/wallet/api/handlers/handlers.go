package handlers

import (
	"encoding/json"
	"net/http"
	"server/server/http/utils"
	"server/server/logger"
	"server/server/models"
	"server/server/services/wallet/usecases"
	"strconv"
)

type Handlers struct {
	u *usecases.WalletService
	l logger.Logger
}

func New(u *usecases.WalletService, l logger.Logger) *Handlers {
	return &Handlers{u, l}
}

func (h *Handlers) TopUpWallet(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.Write400("wrong id value: "+err.Error(), w)
		return
	}
	if id < 0 {
		utils.Write400("wrong id value", w)
		return
	}
	countStr := r.URL.Query().Get("count")
	count, err := strconv.ParseInt(countStr, 10, 64)
	if err != nil {
		utils.Write400("wrong count value: "+err.Error(), w)
		return
	}
	if count < 0 {
		utils.Write400("wrong count value", w)
		return
	}
	err = h.u.TopUpWallet(id, count)
	if err != nil {
		h.l.ErrorLog.Println(err.Error())
		utils.Write500(err.Error(), w)
		return
	}

	utils.Write200(nil, w)
}

func (h *Handlers) GetBalance(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.Write400("wrong id value: "+err.Error(), w)
		return
	}
	if id < 0 {
		utils.Write400("wrong id value", w)
		return
	}
	wallet, err := h.u.GetBalance(id)
	if err != nil {
		utils.Write400(err.Error(), w)
		return
	}
	response, err := json.Marshal(wallet)
	if err != nil {
		utils.Write500(err.Error(), w)
		return
	}
	utils.Write200(response, w)
}

func (h *Handlers) LockMoney(w http.ResponseWriter, r *http.Request) {
	var (
		err                   error
		idStr, serviceIDStr   string
		walletIDStr, priceStr string
		order                 models.Order
	)
	idStr = r.URL.Query().Get("order_id")
	order.ID, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.Write400("wrong order_id value: "+err.Error(), w)
		return
	}
	if order.ID < 0 {
		utils.Write400("wrong order_id value", w)
		return
	}
	serviceIDStr = r.URL.Query().Get("service_id")
	order.ServiceID, err = strconv.ParseInt(serviceIDStr, 10, 64)
	if err != nil {
		utils.Write400("wrong service_id value "+err.Error(), w)
		return
	}
	if order.ServiceID < 0 {
		utils.Write400("wrong service_id value", w)
	}
	walletIDStr = r.URL.Query().Get("wallet_id")
	order.WalletID, err = strconv.ParseInt(walletIDStr, 10, 64)
	if err != nil {
		utils.Write400("wrong wallet_id value "+err.Error(), w)
		return
	}
	if order.WalletID < 0 {
		utils.Write400("wrong wallet_id value", w)
		return
	}
	priceStr = r.URL.Query().Get("price")
	order.Price, err = strconv.ParseInt(priceStr, 10, 64)
	if err != nil {
		utils.Write400("wrong price value "+err.Error(), w)
		return
	}
	if order.Price < 0 {
		utils.Write400("wrong price value", w)
		return
	}
	err = h.u.LockMoney(order)
	if err != nil {
		utils.Write400(err.Error(), w)
		return
	}
	utils.Write200(nil, w)
}
