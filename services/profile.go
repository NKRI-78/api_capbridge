package services

import (
	"errors"
	"superapps/entities"
	helper "superapps/helpers"
)

func GetProfile(gp *entities.GetProfile) (map[string]any, error) {

	profileScan := entities.ProfileScan{}
	users := []entities.ProfileScan{}

	queryUserExist := `SELECT user_id as id, fullname, avatar, created_at, updated_at FROM profiles WHERE user_id = ?`

	err := dbDefault.Debug().Raw(queryUserExist, gp.UserId).Scan(&profileScan).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New(err.Error())
	}

	isUserExist := len(users)

	if isUserExist == 0 {
		return nil, errors.New("USER_NOT_FOUND")
	}

	return map[string]any{
		"data": profileScan,
	}, nil
}

func UpdateProfile(up *entities.UpdateProfile) (map[string]any, error) {

	query := `UPDATE profiles SET fullname = ? WHERE user_id = ?`

	err := dbDefault.Debug().Exec(query, up.Fullname, up.UserId).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New(err.Error())
	}

	return map[string]any{
		"data": "",
	}, nil
}
