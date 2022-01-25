package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"noegotribes/internal/memstore"
	"os"

	"github.com/julienschmidt/httprouter"
)

type LoginAPIHandler struct {
	Store *memstore.MemStore
	*httprouter.Router
}

func NewLoginAPIHandler(mem *memstore.MemStore) *LoginAPIHandler {
	h := &LoginAPIHandler{
		Store:  mem,
		Router: httprouter.New(),
	}
	h.Router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	h.Router.POST("/api/login", h.PostLoginItem)

	return h
}

func (h *LoginAPIHandler) PostLoginItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var q memstore.Address
	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		return
	}
	qstring := fmt.Sprintf("%+v", q)

	f, err := os.OpenFile("address.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(qstring); err != nil {
		log.Println(err)
	}
}
