package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"laundry-branch/model"
	"laundry-branch/service"
	"laundry-branch/utils"

	"github.com/gorilla/mux"
)

// Crear sucursal
func CreateBranch(w http.ResponseWriter, r *http.Request) {
	var branch model.Branch
	if err := json.NewDecoder(r.Body).Decode(&branch); err != nil {
		utils.SendError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	id, err := service.CreateBranch(&branch)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Error al crear sucursal")
		return
	}
	branch.ID = id
	utils.SendSuccess(w, http.StatusCreated, branch)
}

// Listar sucursales
func GetBranches(w http.ResponseWriter, r *http.Request) {
	branches, err := service.GetAllBranches()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Error al obtener sucursales")
		return
	}
	utils.SendSuccess(w, http.StatusOK, branches)
}

// Obtener sucursal por ID
func GetBranchByID(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	branch, err := service.GetBranchByID(id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Error al buscar sucursal")
		return
	}
	if branch == nil {
		utils.SendError(w, http.StatusNotFound, "Sucursal no encontrada")
		return
	}

	utils.SendSuccess(w, http.StatusOK, branch)
}

// Actualizar sucursal
func UpdateBranch(w http.ResponseWriter, r *http.Request) {
	var branch model.Branch
	if err := json.NewDecoder(r.Body).Decode(&branch); err != nil {
		utils.SendError(w, http.StatusBadRequest, "JSON inválido")
		return
	}

	idParam := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	branch.ID = id

	err = service.UpdateBranch(&branch)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Error al actualizar sucursal")
		return
	}

	utils.SendSuccess(w, http.StatusOK, branch)
}

// Eliminar sucursal
func DeleteBranch(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	err = service.DeleteBranch(id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Error al eliminar sucursal")
		return
	}

	utils.SendSuccess(w, http.StatusOK, map[string]string{"message": "Sucursal eliminada"})
}
