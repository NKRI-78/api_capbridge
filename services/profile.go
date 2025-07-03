package services

import (
	"errors"
	"superapps/entities"
	helper "superapps/helpers"
)

func GetProfile(gp *entities.GetProfile) (map[string]any, error) {
	// 1) Scan ke ProfileScan struct
	var user entities.ProfileScan

	queryUserExist := `
		SELECT user_id as id, fullname, avatar, gender, last_edu, status_marital, created_at, updated_at 
		FROM profiles WHERE user_id = ?
	`
	err := dbDefault.Debug().Raw(queryUserExist, gp.UserId).Scan(&user).Error
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	if user.Id == "" {
		return nil, errors.New("USER_NOT_FOUND")
	}

	// 2) Scan ke ProfileEmiten struct
	var emiten entities.ProfileEmiten
	queryEmiten := `
		SELECT company_address, company_name, company_data, company_nib, company_npwp, total_employees, capital_structure, financial_statements, 
		commissioner_name, commissioner_position, commissioner_ktp, commissioner_npwp,
		deed_of_incorporation, latest_amendment_deed, sk_kemenkumham 
		FROM companies WHERE user_id = ?
	`
	errEmiten := dbDefault.Debug().Raw(queryEmiten, gp.UserId).Scan(&emiten).Error
	if errEmiten != nil {
		helper.Logger("error", "In Server: "+errEmiten.Error())
		return nil, errEmiten
	}

	// 3) Gabungkan ke ProfileResponse
	response := entities.ProfileResponse{
		ProfileScan: user,
		Emiten:      emiten,
	}

	return map[string]any{
		"data": response,
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
