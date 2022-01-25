package state

import (
	"errors"
	"fmt"
	"net/http"
	"noegotribes/internal/memstore"

	"golang.org/x/sync/singleflight"
)

type LoginAPI struct {
	Login []memstore.Taco
	g     singleflight.Group
}

func (c *LoginAPI) PostLoginItem() error {

	url := "/api/login"
	res, err := Post(url, "application/json", []byte{1, 2, 3})
	if err != nil {
		err = errors.New(fmt.Sprintf("Error PostLoginItem() %v", err))
		return err
	}
	if res.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("PostLoginItem %s returned status code %v", url, res.StatusCode))
		return err
	}
	return nil
}

func LoadLoginAPI() *LoginAPI {
	return &LoginAPI{}
}

type LoginAPISetter interface {
	LoginAPISet(v *LoginAPI)
}

type LoginAPIRef struct {
	*LoginAPI
}

func (r *LoginAPIRef) LoginAPISet(v *LoginAPI) {
	r.LoginAPI = v
}
