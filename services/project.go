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
	number_of_unit, period 
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
			return nil, errors.New(errProjectRows.Error())
		}

		dataProject = append(dataProject, entities.ProjectListResponse{
			Id:           project.Id,
			Title:        project.Title,
			Goal:         project.Goal,
			Capital:      project.Capital,
			Medias:       []entities.ProjectMedia{},
			Location:     entities.ProjectLocation{},
			Doc:          entities.ProjectDoc{},
			Roi:          project.Roi,
			MinInvest:    project.MinInvest,
			UnitPrice:    project.UnitPrice,
			UnitTotal:    project.UnitTotal,
			NumberOfUnit: project.NumberOfUnit,
			Periode:      project.Periode,
			CreatedAt:    project.CreatedAt,
			UpdatedAt:    project.UpdatedAt,
		})
	}

	return map[string]any{
		"data": dataProject,
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

func ProjectStoreMedia(pm *entities.ProjectMedia) (map[string]any, error) {

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

func ProjectStoreLocation(pl *entities.ProjectLocation) (map[string]any, error) {

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
