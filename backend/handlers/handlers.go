package handlers

import (
	"backend/business"
	_ "embed"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	user business.UserBusiness
}

func New(u business.UserBusiness) user{
	return user{
		user : u,
	}
}

func (h user) EnterData(w http.ResponseWriter, r *http.Request) {
	fName := r.FormValue("F_name")
	lName := r.FormValue("L_name")
	mName := r.FormValue("M_name")
	response := h.user.EnterUserData(fName, lName, mName)
	json.NewEncoder(w).Encode(response)
}

func (h user) GetData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	response := h.user.GetUserData(params)
	json.NewEncoder(w).Encode(response)
}
func (h user) EditData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fName := r.FormValue("F_name")
	lName := r.FormValue("L_name")
	mName := r.FormValue("M_name")
	response := h.user.EditUserData(fName, lName, mName, params)
	json.NewEncoder(w).Encode(response)
}
func (h user) DeleteData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	response := h.user.DeleteUserData(params)
	json.NewEncoder(w).Encode(response)
}
