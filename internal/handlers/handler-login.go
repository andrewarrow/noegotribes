package handlers

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	all, _ := ioutil.ReadAll(r.Body)
	allString := string(all)

	sDec, _ := b64.StdEncoding.DecodeString(allString[1 : len(allString)-1])
	fmt.Println(string(sDec))
	fmt.Println()

	var q memstore.Address
	reader := bytes.NewReader(sDec)
	err := json.NewDecoder(reader).Decode(&q)
	if err != nil {
		fmt.Println(err)
		return
	}
	qstring := fmt.Sprintf("%+v", q)

	f, err := os.OpenFile("address.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(qstring + "\n"); err != nil {
		log.Println(err)
	}
}
