package auth

import (
	"encoding/json"
	"net/http"

	"github.com/auth_service/pkg/handlerutil"
)

func registerUser(service *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service.logger.Infof("%s %s", r.Method, r.URL)

		if r.Method != "POST" {
			service.logger.Error(notFounded)
			resError := handlerutil.NewError(notFounded)
			w.WriteHeader(http.StatusNotFound)
			w.Write(resError)
			return
		}

		r.ParseForm()
		publicID := r.Form.Get("public_id")
		payload := r.Form.Get("payload")

		w.Header().Set("Content-Type", "application/json")

		if len(publicID) == 0 || len(payload) == 0 {
			service.logger.Error(missingParams)
			resError := handlerutil.NewError(missingParams)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(resError)
			return
		}

		user, s_err := service.createNewUser(publicID, payload)
		if s_err != nil {
			service.logger.Error(s_err.Messgae)
			resError := handlerutil.NewError(s_err.Messgae)
			w.WriteHeader(s_err.Httpcode)
			w.Write(resError)
			return
		}

		resdata := struct {
			ID       int    `json:"id"`
			PublicID string `json:"public_id"`
		}{user.ID, user.PublicID}

		jsondata, err := json.Marshal(resdata)
		if err != nil {
			service.logger.Error(err)
			resError := handlerutil.NewError(internalError)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(resError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsondata)
	}
}

func authUser(service *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service.logger.Infof("%s %s", r.Method, r.URL)

		if r.Method != "POST" {
			service.logger.Error(notFounded)
			resError := handlerutil.NewError(notFounded)
			w.WriteHeader(http.StatusNotFound)
			w.Write(resError)
			return
		}

		r.ParseForm()
		publicID := r.Form.Get("public_id")
		payload := r.Form.Get("payload")

		w.Header().Set("Content-Type", "application/json")

		if len(publicID) == 0 || len(payload) == 0 {
			service.logger.Error(missingParams)
			resError := handlerutil.NewError(missingParams)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(resError)
			return
		}

		user, s_err := service.AuthUser(publicID, payload)
		if s_err != nil {
			service.logger.Error(s_err.Messgae)
			resError := handlerutil.NewError(s_err.Messgae)
			w.WriteHeader(s_err.Httpcode)
			w.Write(resError)
			return
		}

		resdata := struct {
			ID       int    `json:"id"`
			PublicID string `json:"public_id"`
		}{user.ID, user.PublicID}

		jsondata, err := json.Marshal(resdata)
		if err != nil {
			service.logger.Error(err)
			resError := handlerutil.NewError(internalError)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(resError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsondata)
	}
}
