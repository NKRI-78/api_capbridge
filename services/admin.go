package services

import (
	"errors"
	"math"
	"os"
	"strconv"
	"superapps/entities"
	helper "superapps/helpers"
)

func AdminListUser(page, limit string) (map[string]any, error) {
	url := os.Getenv("API_URL_DEV")

	var dataAdminListUser []entities.AdminListUserResponse
	var totalAdminListUser []entities.AdminListUser
	var adminListUser entities.AdminListUser

	pageinteger, _ := strconv.Atoi(page)
	limitinteger, _ := strconv.Atoi(limit)

	var offset = strconv.Itoa((pageinteger - 1) * limitinteger)

	errAllUser := dbDefault.Debug().Raw(`SELECT id FROM users`).Scan(&totalAdminListUser).Error

	if errAllUser != nil {
		helper.Logger("error", "In Server: "+errAllUser.Error())
	}

	var resultTotal = len(totalAdminListUser)

	var perPage = math.Ceil(float64(resultTotal) / float64(limitinteger))

	var prevPage int
	var nextPage int

	if pageinteger == 1 {
		prevPage = 1
	} else {
		prevPage = pageinteger - 1
	}

	nextPage = pageinteger + 1

	query := `SELECT u.uid AS id, u.email, p.fullname, u.phone, ur.name AS role, u.created_at, u.updated_at,
	 EXISTS (
		SELECT 1 
		FROM ktps k 
		INNER JOIN accounts a ON a.user_id = k.user_id 
		WHERE k.user_id = u.uid 
		AND (a.no IS NOT NULL AND a.no != "") 
		AND (k.nik IS NOT NULL AND k.nik != "")
	) AS verified
	FROM users u 
	INNER JOIN profiles p ON p.user_id = u.uid 
	INNER JOIN user_roles ur ON ur.id = u.role
	ORDER BY u.created_at DESC
	LIMIT ?, ?`

	rows, err := dbDefault.Debug().Raw(query, offset, limit).Rows()

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New(err.Error())
	}

	for rows.Next() {
		errAdminListUserRows := dbDefault.ScanRows(rows, &adminListUser)

		if errAdminListUserRows != nil {
			helper.Logger("error", "In Server: "+errAdminListUserRows.Error())
			return nil, errors.New(errAdminListUserRows.Error())
		}

		dataAdminListUser = append(dataAdminListUser, entities.AdminListUserResponse{
			Id:        adminListUser.Id,
			Avatar:    helper.DefaultIfEmpty(adminListUser.Avatar, "-"),
			Fullname:  adminListUser.Fullname,
			Email:     adminListUser.Email,
			Phone:     adminListUser.Phone,
			Role:      adminListUser.Role,
			Verified:  adminListUser.Verified,
			CreatedAt: adminListUser.CreatedAt,
			UpdatedAt: adminListUser.UpdatedAt,
		})
	}

	var nextUrl = strconv.Itoa(nextPage)
	var prevUrl = strconv.Itoa(prevPage)

	return map[string]any{
		"total":        resultTotal,
		"current_page": pageinteger,
		"per_page":     int(perPage),
		"prev_page":    prevPage,
		"next_page":    nextPage,
		"next_url":     url + "/api/v1/admin/list/user?page=" + nextUrl,
		"prev_url":     url + "/api/v1/admin/list/user?page=" + prevUrl,
		"data":         dataAdminListUser,
	}, nil
}

func AdminListProject(page, limit string) (map[string]any, error) {
	url := os.Getenv("API_URL_DEV")

	var dataAdminListProject []entities.AdminListProjectResponse
	var totalAdminListProject []entities.AdminListProject
	var adminListProject entities.AdminListProject

	pageinteger, _ := strconv.Atoi(page)
	limitinteger, _ := strconv.Atoi(limit)

	var offset = strconv.Itoa((pageinteger - 1) * limitinteger)

	errAllUser := dbDefault.Debug().Raw(`SELECT id FROM projects`).Scan(&totalAdminListProject).Error

	if errAllUser != nil {
		helper.Logger("error", "In Server: "+errAllUser.Error())
	}

	var resultTotal = len(totalAdminListProject)

	var perPage = math.Ceil(float64(resultTotal) / float64(limitinteger))

	var prevPage int
	var nextPage int

	if pageinteger == 1 {
		prevPage = 1
	} else {
		prevPage = pageinteger - 1
	}

	nextPage = pageinteger + 1

	query := `SELECT p.uid AS id, p.title, p.goal, p.capital, p.roi, p.min_invest, p.min_invest,
	p.unit_price, p.unit_total, p.number_of_unit, p.periode, p.type_of_bond,
	p.nominal_value, p.time_periode, p.interest_rate, p.interest_payment_schedule, p.principal_payment_schedule,
	p.use_of_funds, p.collateral_guarantee, p.desc_job, p.is_apbn, p.is_approved,
	u.uid AS user_id, u.email AS user_email, u.phone AS user_phone
	FROM projects p
	INNER JOIN users u ON u.uid = p.user_id
	LIMIT ?, ?`

	rows, err := dbDefault.Debug().Raw(query, offset, limit).Rows()

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New(err.Error())
	}

	for rows.Next() {
		errAdminListUserRows := dbDefault.ScanRows(rows, &adminListProject)

		if errAdminListUserRows != nil {
			helper.Logger("error", "In Server: "+errAdminListUserRows.Error())
			return nil, errors.New(errAdminListUserRows.Error())
		}

		dataAdminListProject = append(dataAdminListProject, entities.AdminListProjectResponse{
			Id:                       adminListProject.Id,
			Title:                    helper.DefaultIfEmpty(adminListProject.Title, "-"),
			Goal:                     helper.DefaultIfEmpty(adminListProject.Goal, "-"),
			Capital:                  helper.DefaultIfEmpty(adminListProject.Capital, "-"),
			Roi:                      helper.DefaultIfEmpty(adminListProject.Roi, "-"),
			MinInvest:                helper.DefaultIfEmpty(adminListProject.MinInvest, "-"),
			UnitPrice:                helper.DefaultIfEmpty(adminListProject.UnitPrice, "-"),
			UnitTotal:                helper.DefaultIfEmpty(adminListProject.UnitTotal, "-"),
			NumberOfUnit:             helper.DefaultIfEmpty(adminListProject.NumberOfUnit, "-"),
			Periode:                  helper.DefaultIfEmpty(adminListProject.Periode, "-"),
			TypeOfBond:               helper.DefaultIfEmpty(adminListProject.TypeOfBond, "-"),
			NominalValue:             helper.DefaultIfEmpty(adminListProject.NominalValue, "-"),
			TimePeriode:              helper.DefaultIfEmpty(adminListProject.TimePeriode, "-"),
			InterestRate:             helper.DefaultIfEmpty(adminListProject.InterestRate, "-"),
			InterestPaymentSchedule:  helper.DefaultIfEmpty(adminListProject.InterestPaymentSchedule, "-"),
			PrincipalPaymentSchedule: helper.DefaultIfEmpty(adminListProject.PrincipalPaymentSchedule, "-"),
			UseOfFunds:               helper.DefaultIfEmpty(adminListProject.UseOfFunds, "-"),
			CollateralGuarantee:      helper.DefaultIfEmpty(adminListProject.CollateralGuarantee, "-"),
			DescJob:                  helper.DefaultIfEmpty(adminListProject.DescJob, "-"),
			IsApbn:                   adminListProject.IsApbn,
			IsApproved:               adminListProject.IsApproved,
			User: entities.AdminListProjectUser{
				Id:    adminListProject.UserId,
				Email: adminListProject.UserEmail,
				Phone: adminListProject.UserPhone,
			},
		})

	}

	var nextUrl = strconv.Itoa(nextPage)
	var prevUrl = strconv.Itoa(prevPage)

	return map[string]any{
		"total":        resultTotal,
		"current_page": pageinteger,
		"per_page":     int(perPage),
		"prev_page":    prevPage,
		"next_page":    nextPage,
		"next_url":     url + "/api/v1/admin/list/project?page=" + nextUrl,
		"prev_url":     url + "/api/v1/admin/list/project?page=" + prevUrl,
		"data":         dataAdminListProject,
	}, nil
}

func VerifyUser(avu *entities.AdminVerifyUser) (map[string]any, error) {

	query := `UPDATE users SET verify = ? WHERE uid = ?`

	err := dbDefault.Debug().Exec(query, 1, avu.UserId).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New(err.Error())
	}

	return map[string]any{}, nil
}

func VerifyProject(avp *entities.AdminVerifyProject) (map[string]any, error) {

	query := `UPDATE projects SET is_approved = ? WHERE uid = ?`

	err := dbDefault.Debug().Exec(query, 1, avp.Id).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New(err.Error())
	}

	return map[string]any{}, nil
}
