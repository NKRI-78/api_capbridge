package services

import (
	"database/sql"
	"errors"
	"superapps/entities"
	helper "superapps/helpers"

	uuid "github.com/satori/go.uuid"
)

func ProjectList() (map[string]any, error) {

	var project entities.ProjectListScan
	var dataProject = make([]entities.ProjectListResponse, 0)

	query := `SELECT uid AS id, title, goal, capital, roi, min_invest, unit_price, unit_total,
	number_of_unit, periode, type_of_bond, nominal_value, time_periode, interest_rate, interest_payment_schedule,
	principal_payment_schedule, use_of_funds, collateral_guarantee, desc_job, is_apbn, is_approved
	FROM projects`

	var rows *sql.Rows
	var err error

	rows, err = dbDefault.Debug().Raw(query).Rows()

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		errProjectRows := dbDefault.ScanRows(rows, &project)
		if errProjectRows != nil {
			helper.Logger("error", "In Server: "+errProjectRows.Error())
			return nil, errProjectRows
		}

		dataProjectMedia := make([]entities.ProjectMedia, 0)

		queryProjectMedia := `SELECT id, path FROM project_medias WHERE project_id = ?`

		errProjectMedia := dbDefault.Debug().
			Raw(queryProjectMedia, project.Id).
			Scan(&dataProjectMedia).Error

		if errProjectMedia != nil {
			helper.Logger("error", "In Server: "+errProjectMedia.Error())
			return nil, errProjectMedia
		}

		var dataProjectLoc entities.ProjectLocation

		queryProjectLoc := `SELECT id, url, name, lat, lng FROM project_locations WHERE project_id = ?`

		errProjectLoc := dbDefault.Debug().
			Raw(queryProjectLoc, project.Id).
			Scan(&dataProjectLoc).Error

		if errProjectLoc != nil {
			helper.Logger("error", "In Server: "+errProjectLoc.Error())
			return nil, errProjectLoc
		}

		var dataProjectDoc entities.ProjectDoc

		queryProjectDoc := `SELECT id, path FROM project_docs WHERE project_id = ?`

		errProjectDoc := dbDefault.Debug().
			Raw(queryProjectDoc, project.Id).
			Scan(&dataProjectDoc).Error

		if errProjectDoc != nil {
			helper.Logger("error", "In Server: "+errProjectDoc.Error())
			return nil, errProjectDoc
		}

		dataProject = append(dataProject, entities.ProjectListResponse{
			Id:      project.Id,
			Title:   project.Title,
			Goal:    project.Goal,
			Capital: project.Capital,
			Medias:  dataProjectMedia,
			Location: entities.ProjectLocation{
				Id:   dataProjectLoc.Id,
				Url:  helper.DefaultIfEmpty(dataProjectLoc.Url, "-"),
				Name: helper.DefaultIfEmpty(dataProjectLoc.Name, "-"),
				Lat:  helper.DefaultIfEmpty(dataProjectLoc.Lat, "-"),
				Lng:  helper.DefaultIfEmpty(dataProjectLoc.Lng, "-"),
			},
			Doc: entities.ProjectDoc{
				Id:   dataProjectDoc.Id,
				Path: helper.DefaultIfEmpty(dataProjectDoc.Path, "-"),
			},
			Roi:                      project.Roi,
			MinInvest:                project.MinInvest,
			UnitPrice:                project.UnitPrice,
			UnitTotal:                project.UnitTotal,
			NumberOfUnit:             project.NumberOfUnit,
			Periode:                  project.Periode,
			TypeOfBond:               project.TypeOfBond,
			NominalValue:             project.NominalValue,
			TimePeriode:              project.TimePeriode,
			InterestRate:             project.InterestRate,
			InterestPaymentSchedule:  project.InterestPaymentSchedule,
			PrincipalPaymentSchedule: project.PrincipalPaymentSchedule,
			UseOfFunds:               project.UseOfFunds,
			CollateralGuarantee:      project.CollateralGuarantee,
			DescJob:                  project.DescJob,
			IsApbn:                   project.IsApbn,
			IsApproved:               project.IsApproved,
			CreatedAt:                project.CreatedAt,
			UpdatedAt:                project.UpdatedAt,
		})
	}

	return map[string]any{
		"data": dataProject,
	}, nil
}

func ProjectDetail(id string) (map[string]any, error) {
	var project entities.ProjectListScan

	query := `SELECT uid AS id, title, goal, capital, roi, min_invest, unit_price, unit_total,
		number_of_unit, periode, type_of_bond, nominal_value, time_periode, interest_rate, 
		interest_payment_schedule, principal_payment_schedule, use_of_funds, collateral_guarantee, 
		desc_job, is_apbn, is_approved, user_id, created_at, updated_at
		FROM projects WHERE uid = ?`

	err := dbDefault.Debug().Raw(query, id).Scan(&project).Error
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	// Fetch project medias
	dataProjectMedia := make([]entities.ProjectMedia, 0)
	queryProjectMedia := `SELECT id, path FROM project_medias WHERE project_id = ?`
	errProjectMedia := dbDefault.Debug().Raw(queryProjectMedia, project.Id).Scan(&dataProjectMedia).Error
	if errProjectMedia != nil {
		helper.Logger("error", "In Server: "+errProjectMedia.Error())
		return nil, errProjectMedia
	}

	// Fetch project location
	var dataProjectLoc entities.ProjectLocation
	queryProjectLoc := `SELECT id, url, name, lat, lng FROM project_locations WHERE project_id = ?`
	errProjectLoc := dbDefault.Debug().Raw(queryProjectLoc, project.Id).Scan(&dataProjectLoc).Error
	if errProjectLoc != nil {
		helper.Logger("error", "In Server: "+errProjectLoc.Error())
		return nil, errProjectLoc
	}

	// Fetch project docs
	var dataProjectDoc entities.ProjectDoc
	queryProjectDoc := `SELECT id, path FROM project_docs WHERE project_id = ?`
	errProjectDoc := dbDefault.Debug().Raw(queryProjectDoc, project.Id).Scan(&dataProjectDoc).Error
	if errProjectDoc != nil {
		helper.Logger("error", "In Server: "+errProjectDoc.Error())
		return nil, errProjectDoc
	}

	// Fetch project company
	var dataProjectCompany entities.ProjectCompany
	queryProjectCompany := `SELECT company_name FROM companies WHERE user_id = ?`
	errProjectCompany := dbDefault.Debug().Raw(queryProjectCompany, project.UserId).Scan(&dataProjectCompany).Error
	if errProjectCompany != nil {
		helper.Logger("error", "In Server: "+errProjectCompany.Error())
		return nil, errProjectCompany
	}

	// Compose detail response, SAME with ProjectListResponse
	projectDetail := entities.ProjectListResponse{
		Id:      project.Id,
		Title:   project.Title,
		Goal:    project.Goal,
		Capital: project.Capital,
		Medias:  dataProjectMedia,
		Location: entities.ProjectLocation{
			Id:   dataProjectLoc.Id,
			Url:  helper.DefaultIfEmpty(dataProjectLoc.Url, "-"),
			Name: helper.DefaultIfEmpty(dataProjectLoc.Name, "-"),
			Lat:  helper.DefaultIfEmpty(dataProjectLoc.Lat, "-"),
			Lng:  helper.DefaultIfEmpty(dataProjectLoc.Lng, "-"),
		},
		Doc: entities.ProjectDoc{
			Id:   dataProjectDoc.Id,
			Path: helper.DefaultIfEmpty(dataProjectDoc.Path, "-"),
		},
		Roi:                      project.Roi,
		MinInvest:                project.MinInvest,
		UnitPrice:                project.UnitPrice,
		UnitTotal:                project.UnitTotal,
		NumberOfUnit:             project.NumberOfUnit,
		Periode:                  project.Periode,
		TypeOfBond:               project.TypeOfBond,
		NominalValue:             project.NominalValue,
		TimePeriode:              project.TimePeriode,
		InterestRate:             project.InterestRate,
		InterestPaymentSchedule:  project.InterestPaymentSchedule,
		PrincipalPaymentSchedule: project.PrincipalPaymentSchedule,
		UseOfFunds:               project.UseOfFunds,
		CollateralGuarantee:      project.CollateralGuarantee,
		DescJob:                  project.DescJob,
		IsApbn:                   project.IsApbn,
		IsApproved:               project.IsApproved,
		Company: entities.Company{
			dataProjectCompany.CompanyName,
		},
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
	}

	return map[string]any{
		"data": projectDetail,
	}, nil
}

func ProjectStore(ps *entities.ProjectStore) (map[string]any, error) {

	ps.Id = uuid.NewV4().String()

	queryInsertProject := `INSERT INTO projects (id, title, goal, capital, roi, min_invest, unit_price, unit_total, number_of_unit, periode) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	errInsertProject := dbDefault.Debug().Exec(queryInsertProject, ps.Id, ps.Title,
		ps.Goal, ps.Capital, ps.Roi, ps.MinInvest, ps.UnitPrice, ps.UnitTotal, ps.NumberOfUnit,
	).Error

	if errInsertProject != nil {
		helper.Logger("error", "In Server: "+errInsertProject.Error())
		return nil, errors.New(errInsertProject.Error())
	}

	queryInsertProjectMedia := `INSERT INTO project_medias (id, project_id, path) VALUES (?, ?, ?)`

	for _, v := range ps.Medias {
		errExerciseMedia := dbDefault.Debug().Exec(queryInsertProjectMedia, ps.Id, v.Path).Error

		if errExerciseMedia != nil {
			helper.Logger("error", "In Server: "+errExerciseMedia.Error())
			return nil, errors.New(errExerciseMedia.Error())
		}
	}

	queryInsertProjectLocation := `INSERT INTO project_locations (project_id, name, url, lat, lng) VALUES (?, ?, ?, ?, ?)`

	errInsertProjectLocation := dbDefault.Debug().Exec(queryInsertProjectLocation, ps.Id,
		ps.Location.Name, ps.Location.Url, ps.Location.Lat, ps.Location.Lng).Error

	if errInsertProjectLocation != nil {
		helper.Logger("error", "In Server: "+errInsertProjectLocation.Error())
		return nil, errors.New(errInsertProjectLocation.Error())
	}

	return map[string]any{}, nil
}

func ProjectUpdate(pu *entities.ProjectUpdate) (map[string]any, error) {
	return map[string]any{}, nil
}

func ProjectDelete(pd *entities.ProjectDelete) (map[string]any, error) {
	return map[string]any{}, nil
}

func ProjectStoreMedia(pm *entities.ProjectStoreMedia) (map[string]any, error) {

	queryInsertProjectMedia := `INSERT INTO project_medias (id, project_id, path) VALUES (?, ?, ?)`

	errInsertProjectMedia := dbDefault.Debug().Exec(queryInsertProjectMedia, pm.Id, pm.ProjectId, pm.Path).Error

	if errInsertProjectMedia != nil {
		helper.Logger("error", "In Server: "+errInsertProjectMedia.Error())
		return nil, errors.New(errInsertProjectMedia.Error())
	}

	return map[string]any{
		"data": pm,
	}, nil
}

func ProjectStoreLocation(pl *entities.ProjectStoreLocation) (map[string]any, error) {

	queryInsertProjectLocation := `INSERT INTO project_locations (project_id, name, url, lat, lng) VALUES (?, ?, ?, ?, ?)`

	errInsertProjectLocation := dbDefault.Debug().Exec(queryInsertProjectLocation, pl.ProjectId, pl.Name, pl.Url, pl.Lat, pl.Lng).Error

	if errInsertProjectLocation != nil {
		helper.Logger("error", "In Server: "+errInsertProjectLocation.Error())
		return nil, errors.New(errInsertProjectLocation.Error())
	}

	return map[string]any{
		"data": pl,
	}, nil
}
