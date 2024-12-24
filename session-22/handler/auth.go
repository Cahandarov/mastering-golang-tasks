package handler

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"session-22/model"
	"session-22/service"
	"session-22/util"

	"net/http"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler() *AuthHandler {
	log.Info("NewAuthHandler generated")

	s := service.NewAuthService()

	app := &AuthHandler{
		s,
	}

	return app
}

func (a AuthHandler) login(w http.ResponseWriter, r *http.Request) {

	var authRequest model.AuthRequest

	if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
		log.Error("Error while decoding request body:", err)
		util.PrepareResponse(w, "", &model.InValidRequestError)
		return
	}

	resp, err := a.AuthService.Login(authRequest)

	if err != nil {
		util.PrepareResponse(w, "", err)
		return
	}

	util.PrepareResponse(w, resp, nil)

}
