package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/A-Victory/user-mig/user/models"
	"github.com/A-Victory/user-mig/user/service"
	"golang.org/x/crypto/bcrypt"
)

type handlers struct {
	service *service.Service
}

func NewHandlers(s *service.Service) *handlers {
	return &handlers{
		service: s,
	}
}

func (h *handlers) SignUp(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	new_user := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err.Error())
		json.NewEncoder(w).Encode("failed to create new user")
		return
	}

	hashPassowrd, err := genHashpassword(req.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		json.NewEncoder(w).Encode("failed to create new user")
		return
	}

	new_user.Email = req.Email
	new_user.FirstName = req.FirstName
	new_user.LastName = req.LastName
	new_user.Password = hashPassowrd

	user, err := h.service.GetUserByEmail(req.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		json.NewEncoder(w).Encode("failed to signin, an error occurred")
		return
	}

	if user.Email == req.Email {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("email already in use, signup with a different email")
		return
	}

	if err := h.service.CreateNewUser(new_user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		json.NewEncoder(w).Encode("failed to create new user")
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("user created successfully")

}

func (h *handlers) SignIn(w http.ResponseWriter, r *http.Request) {

	type login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	loginDetails := login{}

	if err := json.NewDecoder(r.Body).Decode(&loginDetails); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err.Error())
		json.NewEncoder(w).Encode("failed to create new user")
		return
	}

	user, err := h.service.GetUserByEmail(loginDetails.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		json.NewEncoder(w).Encode("failed to signin, an error occurred")
		return
	}

	if user == (models.User{}) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("no user associated with the email address")
		return
	}

	valid := comparePassword(loginDetails.Password, user.Password)
	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("incorrect password!!! Please try again")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func (h *handlers) ListUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.service.ListAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err.Error())
		json.NewEncoder(w).Encode("could not retrieve users, please try again")
		return
	}

	if users == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("no user entry yet!")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

func genHashpassword(password string) (string, error) {

	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashedPassword := string(b)

	return hashedPassword, nil
}

func comparePassword(inputPassword, dbPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(inputPassword))
	return err == nil
}
