package service

import (
	"laundry-branch/model"
	"laundry-branch/repository"
)

func CreateBranch(branch *model.Branch) (int64, error) {
	return repository.CreateBranch(branch)
}

func GetAllBranches() ([]model.Branch, error) {
	return repository.GetAllBranches()
}

func GetBranchByID(id int64) (*model.Branch, error) {
	return repository.GetBranchByID(id)
}

func UpdateBranch(branch *model.Branch) error {
	return repository.UpdateBranch(branch)
}

func DeleteBranch(id int64) error {
	return repository.DeleteBranch(id)
}
