package main

import (
	// Standard library packages

	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/corriganrenard/sysDash/controllers"
	"github.com/julienschmidt/httprouter"
)

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

//handler to serve static index.html
func HomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	http.ServeFile(w, r, "public/index.html")
	http.ServeFile(w, r, "public/index.css")
	http.ServeFile(w, r, "public/main.js")
}

func main() {
	// Instantiate a new router
	r := httprouter.New()

	//Get a IpController instance
	ipc := controllers.NewIpController(getSession())

	//serve the index page
	r.GET("/", HomeHandler)
	//get an ip record from the database using its bson _id
	r.GET("/ip/:id", ipc.GetIp)
	//create an IP and push to database
	r.POST("/ip", ipc.CreateIp)
	//delete an ip from the database using it's bson _id
	r.DELETE("/ip/:id", ipc.RemoveIp)

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)

	//ips := []string{"192.168.1.148", "192.168.10.116"}

	//result := pinger.Ping(ips)

	// for k, v := range result {
	// 	fmt.Printf("ip=%s success=%t\n", k, v)
	// }

}
