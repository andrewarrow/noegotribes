package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"noegotribes/internal/memstore"

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
	var tx interface{}
	all, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(all, &tx)
	fmt.Println("Tx", tx)
}
