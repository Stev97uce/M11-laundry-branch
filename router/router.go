package router

import (
	"laundry-branch/controller"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/branches", controller.CreateBranch).Methods("POST")
	r.HandleFunc("/api/branches", controller.GetBranches).Methods("GET")
	r.HandleFunc("/api/branches/{id}", controller.GetBranchByID).Methods("GET")
	r.HandleFunc("/api/branches/{id}", controller.UpdateBranch).Methods("PUT")
	r.HandleFunc("/api/branches/{id}", controller.DeleteBranch).Methods("DELETE")

	return r
}
