package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/corriganrenard/sysDash/models"
	"github.com/julienschmidt/httprouter"
)

type (
	IpController struct {
		session *mgo.Session
	}
)

func NewIpController(s *mgo.Session) *IpController {
	return &IpController{s}
}

func (ipc IpController) GetIp(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub user
	ip := models.IpAddress{}

	// Fetch user
	if err := ipc.session.DB("ip_list").C("ips").FindId(oid).One(&ip); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	ipj, _ := json.Marshal(ip)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", ipj)
}

func (ipc IpController) CreateIp(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an user to be populated from the body
	ip := models.IpAddress{}
	// Populate the user data

	json.NewDecoder(r.Body).Decode(&ip)

	// Add an Id
	ip.Id = bson.NewObjectId()

	// Write the user to mongo
	ipc.session.DB("ip_list").C("ips").Insert(ip)

	// Marshal provided interface into JSON structure
	ipj, _ := json.Marshal(ip)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", ipj)
}

// RemoveUser removes an existing user resource
func (ipc IpController) RemoveIp(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := ipc.session.DB("ip_list").C("ips").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}
