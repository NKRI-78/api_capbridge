package services

import (
	"errors"
	"superapps/entities"
	helper "superapps/helpers"
	"superapps/middlewares"

	uuid "github.com/satori/go.uuid"
)

func Login(l *entities.Login) (entities.LoginResponse, error) {

	users := []entities.LoginScan{}
	roles := []entities.CheckRole{}

	queryUserExist := `SELECT uid AS id, enabled, password, verify, role, email FROM users WHERE email = ?`

	errUser := dbDefault.Debug().Raw(queryUserExist, l.Email).Scan(&users).Error

	if errUser != nil {
		helper.Logger("error", "In Server: "+errUser.Error())
		return entities.LoginResponse{}, errors.New(errUser.Error())
	}

	isUserExist := len(users)

	if isUserExist == 0 {
		helper.Logger("error", "In Server: USER_NOT_FOUND")
		return entities.LoginResponse{}, errors.New("USER_NOT_FOUND")
	}

	queryCheckRole := `SELECT id, name FROM user_roles WHERE id = ?`

	errCheckRole := dbDefault.Debug().Raw(queryCheckRole, users[0].Role).Scan(&roles).Error

	if errCheckRole != nil {
		helper.Logger("error", "In Server: "+errCheckRole.Error())
		return entities.LoginResponse{}, errors.New(errCheckRole.Error())
	}

	isCheckRoleExist := len(roles)

	if isCheckRoleExist == 0 {
		helper.Logger("error", "In Server: ROLE_NOT_FOUND")
		return entities.LoginResponse{}, errors.New("ROLE_NOT_FOUND")
	}

	passHashed := users[0].Password

	errVerify := helper.VerifyPassword(passHashed, l.Password)

	if errVerify != nil {
		helper.Logger("error", "In Server: CREDENTIALS_IS_INCORRECT")
		return entities.LoginResponse{}, errors.New("CREDENTIALS_IS_INCORRECT")
	}

	token, errToken := middlewares.CreateToken(users[0].Id)
	if errToken != nil {
		helper.Logger("error", "In Server: "+errToken.Error())
		return entities.LoginResponse{}, errToken
	}

	access := token["token"]

	return entities.LoginResponse{
		Id:      users[0].Id,
		Email:   users[0].Email,
		Enabled: users[0].Enabled,
		Verify:  users[0].Verify,
		Role:    roles[0].Name,
		Token:   access,
	}, nil
}

func Register(r *entities.Register) (entities.RegisterResponse, error) {

	hashedPassword, errHasshed := helper.Hash(r.Password)
	if errHasshed != nil {
		helper.Logger("error", "In Server: "+errHasshed.Error())
		return entities.RegisterResponse{}, errors.New(errHasshed.Error())
	}

	users := []entities.LoginScan{}
	roles := []entities.CheckRole{}

	r.UserId = uuid.NewV4().String()

	queryCheckRole := `SELECT id, name FROM user_roles WHERE id = ?`

	errCheckRole := dbDefault.Debug().Raw(queryCheckRole, r.Role).Scan(&roles).Error

	if errCheckRole != nil {
		helper.Logger("error", "In Server: "+errCheckRole.Error())
		return entities.RegisterResponse{}, errors.New(errCheckRole.Error())
	}

	isCheckRoleExist := len(roles)

	if isCheckRoleExist == 0 {
		helper.Logger("error", "In Server: ROLE_NOT_FOUND")
		return entities.RegisterResponse{}, errors.New("ROLE_NOT_FOUND")
	}

	queryUserExist := `SELECT uid AS id, enabled, password, verify, email FROM users WHERE email = ?`

	errUserExist := dbDefault.Debug().Raw(queryUserExist, r.Email).Scan(&users).Error

	if errUserExist != nil {
		helper.Logger("error", "In Server: "+errUserExist.Error())
		return entities.RegisterResponse{}, errors.New(errUserExist.Error())
	}

	isUserExist := len(users)

	if isUserExist == 1 {
		helper.Logger("error", "In Server: USER_ALREADY_EXIST")
		return entities.RegisterResponse{}, errors.New("USER_ALREADY_EXIST")
	}

	queryInsertUser := `INSERT INTO users (uid, email, phone, password, role) VALUES (?, ?, ?, ?, ?)`

	errInsertUser := dbDefault.Debug().Exec(queryInsertUser, r.UserId, r.Email, r.Phone, hashedPassword, r.Role).Error

	if errInsertUser != nil {
		helper.Logger("error", "In Server: "+errInsertUser.Error())
		return entities.RegisterResponse{}, errors.New(errInsertUser.Error())
	}

	if r.Role == "1" {

		queryInsertProfile := `INSERT INTO profiles (user_id, fullname, gender, last_edu, status_marital) VALUES (?, ?, ?, ?, ?)`

		errInsertProfile := dbDefault.Debug().Exec(queryInsertProfile, r.UserId, r.Fullname, r.Investor.Gender, r.Investor.LastEdu, r.Investor.StatusMarital).Error

		if errInsertProfile != nil {
			helper.Logger("error", "In Server: "+errInsertProfile.Error())
			return entities.RegisterResponse{}, errInsertProfile
		}

		queryInsertAccount := `INSERT INTO accounts (user_id, no, bank_name, bank_branch, bank_owner) VALUES (?, ?, ?, ?, ?)`

		errInsertAccount := dbDefault.Debug().Exec(queryInsertAccount,
			r.UserId, r.Investor.Bank.No, r.Investor.Bank.Name, r.Investor.Bank.Branch, r.Investor.Bank.Owner,
		).Error

		if errInsertAccount != nil {
			helper.Logger("error", "In Server: "+errInsertAccount.Error())
			return entities.RegisterResponse{}, errInsertAccount
		}

		queryInsertKtp := `INSERT INTO ktps (user_id, nik, place_and_datebirth) VALUES (?, ?, ?)`

		errInsertKtp := dbDefault.Debug().Exec(queryInsertKtp, r.UserId, r.Investor.Ktp, r.Investor.AddressKtp).Error

		if errInsertKtp != nil {
			helper.Logger("error", "In Server: "+errInsertKtp.Error())
			return entities.RegisterResponse{}, errInsertKtp
		}

		queryInsertJob := `INSERT INTO jobs (company_name, company_address, monthly_income, position, user_id) 
		VALUES (?, ?, ?, ?, ?)`

		errInsertJob := dbDefault.Debug().Exec(queryInsertJob, r.Investor.Job.CompanyName, r.Investor.Job.CompanyAddress, r.Investor.Job.MonthlyIncome, r.Investor.Job.Position, r.UserId).Error

		if errInsertJob != nil {
			helper.Logger("error", "In Server: "+errInsertJob.Error())
			return entities.RegisterResponse{}, errInsertJob
		}
	}

	if r.Role == "2" {

		queryInsertProfile := `INSERT INTO profiles (user_id, fullname) VALUES (?, ?)`

		errInsertProfile := dbDefault.Debug().Exec(queryInsertProfile, r.UserId, r.Fullname).Error

		if errInsertProfile != nil {
			helper.Logger("error", "In Server: "+errInsertProfile.Error())
			return entities.RegisterResponse{}, errInsertProfile
		}

		queryInsertCompany := `INSERT INTO companies (
			company_data, company_name, company_nib, 
			deed_of_incorporation, latest_amendment_deed, sk_kemenkumham, company_address, company_npwp,
			total_employees, capital_structure, financial_statements, commissioner_name, commissioner_position,
			commissioner_ktp, commissioner_npwp, director_name, director_position, director_ktp, director_npwp, user_id
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

		errInsertCompany := dbDefault.Debug().Exec(queryInsertCompany,
			r.Emiten.CompanyData, r.Emiten.CompanyName, r.Emiten.CompanyNib, r.Emiten.DeedOfIncorporation, r.Emiten.LatestAmendmentDeed, r.Emiten.SkKemenkumham, r.Emiten.CompanyAddress, r.Emiten.CompanyNpwp,
			r.Emiten.TotalEmployees, r.Emiten.CapitalStructure, r.Emiten.FinancialStatements, r.Emiten.CommisionerName, r.Emiten.CommisionerPosition, r.Emiten.CommisionerKtp, r.Emiten.CommisionerNpwp,
			r.Emiten.DirectorName, r.Emiten.DirectorPosition, r.Emiten.DirectorKtp, r.Emiten.DirectorNpwp, r.UserId,
		).Error

		if errInsertCompany != nil {
			helper.Logger("error", "In Server: "+errInsertCompany.Error())
			return entities.RegisterResponse{}, errInsertCompany
		}

		queryInsertBond := `INSERT INTO projects 
		(uid, user_id, title, type_of_bond, nominal_value, time_periode, interest_rate, interest_payment_schedule, principal_payment_schedule, use_of_funds, collateral_guarantee, desc_job, is_apbn)
	 	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

		isApbn := "0"
		if r.Emiten.InfoBond.IsApbn {
			isApbn = "1"
		}

		projectId := uuid.NewV4().String()

		errInsertBond := dbDefault.Debug().Exec(queryInsertBond,
			projectId,
			r.UserId,
			r.Emiten.InfoBond.Title,
			r.Emiten.InfoBond.TypeOfBond,
			r.Emiten.InfoBond.NominalValue,
			r.Emiten.InfoBond.TimePeriode,
			r.Emiten.InfoBond.InterestRate,
			r.Emiten.InfoBond.InterestPaymentSchedule,
			r.Emiten.InfoBond.PrincipalPaymentSchedule,
			r.Emiten.InfoBond.UseOfFunds,
			r.Emiten.InfoBond.CollateralGuarantee,
			r.Emiten.InfoBond.DescJob,
			isApbn,
		).Error

		if errInsertBond != nil {
			helper.Logger("error", "In Server: "+errInsertBond.Error())
			return entities.RegisterResponse{}, errInsertBond
		}

		queryInsertBondMedia := `INSERT INTO project_medias (project_id, path) VALUES (?, ?)`

		errInsertBondMedia := dbDefault.Debug().Exec(queryInsertBondMedia, projectId, r.Emiten.InfoBond.Img).Error

		if errInsertBondMedia != nil {
			helper.Logger("error", "In Server: "+errInsertBondMedia.Error())
			return entities.RegisterResponse{}, errInsertBondMedia
		}

		queryInsertProjectLoc := `INSERT INTO project_locations (project_id, name, url, lat, lng) VALUES (?, ?, ?, ?, ?)`

		errInsertProjectLoc := dbDefault.Debug().Exec(queryInsertProjectLoc,
			projectId, r.Emiten.InfoBond.Location.Name, r.Emiten.InfoBond.Location.Url,
			r.Emiten.InfoBond.Location.Lat, r.Emiten.InfoBond.Location.Lng,
		).Error

		if errInsertProjectLoc != nil {
			helper.Logger("error", "In Server: "+errInsertProjectLoc.Error())
			return entities.RegisterResponse{}, errInsertProjectLoc
		}

		queryInsertProjectDoc := `INSERT INTO project_docs (project_id, path) VALUES (?, ?)`

		errInsertProjectDoc := dbDefault.Debug().Exec(queryInsertProjectDoc,
			projectId, r.Emiten.InfoBond.Doc,
		).Error

		if errInsertProjectDoc != nil {
			helper.Logger("error", "In Server: "+errInsertProjectDoc.Error())
			return entities.RegisterResponse{}, errInsertProjectDoc
		}

	}

	token, errToken := middlewares.CreateToken(r.UserId)
	if errToken != nil {
		helper.Logger("error", "In Server: "+errToken.Error())
		return entities.RegisterResponse{}, errToken
	}

	access := token["token"]

	return entities.RegisterResponse{
		Id:      r.UserId,
		Email:   r.Email,
		Enabled: false,
		Verify:  false,
		Role:    roles[0].Name,
		Token:   access,
	}, nil
}
