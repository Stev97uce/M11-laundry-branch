package repository

import (
	"database/sql"
	"log"

	"laundry-branch/database"
	"laundry-branch/model"
)

func CreateBranch(branch *model.Branch) (int64, error) {
	stmt, err := database.DB.Prepare(`
		INSERT INTO branches (name, address, city, phone_number, email, open_time, close_time)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(branch.Name, branch.Address, branch.City,
		branch.PhoneNumber, branch.Email, branch.OpenTime, branch.CloseTime)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func GetAllBranches() ([]model.Branch, error) {
	rows, err := database.DB.Query(`SELECT id, name, address, city, phone_number, email, open_time, close_time FROM branches`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var branches []model.Branch
	for rows.Next() {
		var b model.Branch
		err := rows.Scan(&b.ID, &b.Name, &b.Address, &b.City, &b.PhoneNumber, &b.Email, &b.OpenTime, &b.CloseTime)
		if err != nil {
			log.Println("Error escaneando fila:", err)
			continue
		}
		branches = append(branches, b)
	}
	return branches, nil
}

func GetBranchByID(id int64) (*model.Branch, error) {
	var b model.Branch
	err := database.DB.QueryRow(`
		SELECT id, name, address, city, phone_number, email, open_time, close_time
		FROM branches WHERE id = ?
	`, id).Scan(&b.ID, &b.Name, &b.Address, &b.City, &b.PhoneNumber, &b.Email, &b.OpenTime, &b.CloseTime)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &b, nil
}

func UpdateBranch(branch *model.Branch) error {
	_, err := database.DB.Exec(`
		UPDATE branches SET name=?, address=?, city=?, phone_number=?, email=?, open_time=?, close_time=?
		WHERE id=?
	`, branch.Name, branch.Address, branch.City, branch.PhoneNumber, branch.Email, branch.OpenTime, branch.CloseTime, branch.ID)
	return err
}

func DeleteBranch(id int64) error {
	_, err := database.DB.Exec(`DELETE FROM branches WHERE id = ?`, id)
	return err
}
