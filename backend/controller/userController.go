package controller

import (
	"api/userstruct"
	"net/http"
	"api/validation"
	"api/utils"
)



func CreateUser(w http.ResponseWriter, r *http.Request) {
    user := userstruct.User{}
	customErr := &validation.CustomError{}
    if customErr.Validate(w,r, &user) {
        return
    }
	utils.RespondWithJSON(w, http.StatusOK, user)
}


