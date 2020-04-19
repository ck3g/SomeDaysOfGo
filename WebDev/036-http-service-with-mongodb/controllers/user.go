package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ck3g/SomeDaysOfGo/WebDev/036-http-service-with-mongodb/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserController defines a container for the controller
type UserController struct {
	session *mgo.Session
}

// NewUserController initializes the controller
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser defines and endpoint to read info about the user
func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	user := models.User{
		Name:   "John Doe",
		Gender: "male",
		Age:    32,
		ID:     bson.ObjectId(p.ByName("id")),
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userJSON)
}

// CreateUser defines an endpoint to create new users
func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user := models.User{}

	json.NewDecoder(req.Body).Decode(&user)

	user.ID = "503"

	userJSON, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", userJSON)
}

// DeleteUser defines an endpoint to delete the user
func (uc UserController) DeleteUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Write code to delete user\n")
}
