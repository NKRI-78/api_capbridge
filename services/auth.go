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

	queryUserExist := `SELECT uid AS id, enabled, password, verify, email FROM users WHERE email = ?`

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

	queryInsertUser := `INSERT INTO users (uid, email, phone, password) VALUES (?, ?, ?, ?)`

	errInsertUser := dbDefault.Debug().Exec(queryInsertUser, r.UserId, r.Email, r.Phone, hashedPassword).Error

	if errInsertUser != nil {
		helper.Logger("error", "In Server: "+errInsertUser.Error())
		return entities.RegisterResponse{}, errors.New(errInsertUser.Error())
	}

	queryInsertProfile := `INSERT INTO profiles (user_id, fullname) VALUES (?, ?)`

	errInsertProfile := dbDefault.Debug().Exec(queryInsertProfile, r.UserId, r.Fullname).Error

	if errInsertProfile != nil {
		helper.Logger("error", "In Server: "+errInsertProfile.Error())
		return entities.RegisterResponse{}, errInsertProfile
	}

	if r.Role != "4" {

		queryInsertAccount := `INSERT INTO accounts (user_id) VALUES (?)`

		errInsertAccount := dbDefault.Debug().Exec(queryInsertAccount, r.UserId).Error

		if errInsertAccount != nil {
			helper.Logger("error", "In Server: "+errInsertAccount.Error())
			return entities.RegisterResponse{}, errInsertAccount
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
		Token:   access,
	}, nil
}
