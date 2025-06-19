package services

import (
	"database/sql"
	"errors"
	"superapps/entities"
	helper "superapps/helpers"
)

func ProjectList() (map[string]any, error) {

	var project entities.ProjectListScan
	var dataProject = make([]entities.ProjectListResponse, 0)

	query := `SELECT uid AS id, title, goal, capital, roi, min_invest, unit_price, unit_total,
	number_of_unit, period 
	FROM projects`

	var rows *sql.Rows
	var err error

	rows, err = dbPPOB.Debug().Raw(query).Rows()

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		errProjectRows := dbPPOB.ScanRows(rows, &project)

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

	return map[string]any{}, nil
}

func ProjectUpdate(pu *entities.ProjectUpdate) (map[string]any, error) {

	return map[string]any{}, nil
}

func ProjectDelete(pd *entities.ProjectDelete) (map[string]any, error) {
	return map[string]any{}, nil
}
