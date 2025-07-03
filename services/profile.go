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
		SELECT p.user_id as id, p.fullname, p.avatar, p.gender, p.last_edu, p.status_marital, p.created_at, p.updated_at,
		k.nik, k.place_and_datebirth, j.company_name, j.company_address, j.monthly_income, j.position,
		a.no, a.bank_name, bank_owner, bank_branch
		FROM profiles p
		LEFT JOIN ktps k ON k.user_id = p.user_id
		LEFT JOIN jobs j ON j.user_id = p.user_id
		LEFT JOIN accounts a ON a.user_id = p.user_id
		WHERE p.user_id = ?
	`
	err := dbDefault.Debug().Raw(queryUserExist, gp.UserId).Scan(&user).Error
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	if user.Id == "" {
		return nil, errors.New("USER_NOT_FOUND")
	}

	var emiten entities.ProfileEmiten
	queryEmiten := `
		SELECT company_address, company_name, company_data, company_nib, company_npwp, total_employees, capital_structure, financial_statements, 
		commissioner_name, commissioner_position, commissioner_ktp, commissioner_npwp, director_name, director_position, director_ktp, director_npwp, 
		deed_of_incorporation, latest_amendment_deed, sk_kemenkumham 
		FROM companies WHERE user_id = ?
	`
	errEmiten := dbDefault.Debug().Raw(queryEmiten, gp.UserId).Scan(&emiten).Error
	if errEmiten != nil {
		helper.Logger("error", "In Server: "+errEmiten.Error())
		return nil, errEmiten
	}

	response := entities.ProfileResponse{
		Profile: entities.Profile{
			Id:        user.Id,
			Fullname:  helper.DefaultIfEmpty(user.Fullname, "-"),
			Avatar:    helper.DefaultIfEmpty(user.Avatar, "-"),
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Emiten: emiten,
		Investor: entities.ProfileInvestor{
			StatusMarital: helper.DefaultIfEmpty(user.StatusMarital, "-"),
			Gender:        helper.DefaultIfEmpty(user.Gender, "-"),
			LastEdu:       helper.DefaultIfEmpty(user.LastEdu, "-"),
			Ktp:           helper.DefaultIfEmpty(user.Nik, "-"),
			AddressKtp:    helper.DefaultIfEmpty(user.PlaceAndDatebirth, "-"),
			Bank: entities.ProfileBank{
				No:     helper.DefaultIfEmpty(user.No, "-"),
				Name:   helper.DefaultIfEmpty(user.BankName, "-"),
				Owner:  helper.DefaultIfEmpty(user.BankOwner, "-"),
				Branch: helper.DefaultIfEmpty(user.BankBranch, "-"),
			},
			Job: entities.ProfileJob{
				CompanyName:    helper.DefaultIfEmpty(user.CompanyName, "-"),
				CompanyAddress: helper.DefaultIfEmpty(user.CompanyAddress, "-"),
				MonthlyIncome:  helper.DefaultIfEmpty(user.MonthlyIncome, "-"),
				Position:       helper.DefaultIfEmpty(user.Position, "-"),
			},
		},
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
