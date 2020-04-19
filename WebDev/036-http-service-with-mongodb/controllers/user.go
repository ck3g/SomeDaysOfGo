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

const docName = "some-days-of-go-playground"

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
	id := p.ByName("id")

	// Verify id is ObhectId hex representation, otherwise return status not found
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// ObjectIdHex returns an ObjectId rom the provided hex representation
	oid := bson.ObjectIdHex(id)

	user := models.User{}

	err := uc.session.DB(docName).C("users").FindId(oid).One(&user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
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

	user.ID = bson.NewObjectId()
	uc.session.DB(docName).C("users").Insert(user)

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
