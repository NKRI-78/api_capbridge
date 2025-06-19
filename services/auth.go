package services

import (
	"errors"
	"superapps/entities"
	helper "superapps/helpers"
	middleware "superapps/middlewares"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(l *entities.Login) (map[string]any, error) {

	loginScan := entities.LoginScan{}
	users := []entities.LoginScan{}

	queryUserExist := `SELECT uid AS id, enabled, password, verify FROM users WHERE email = ?`

	err := dbDefault.Debug().Raw(queryUserExist, l.Email).Scan(&loginScan).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New(err.Error())
	}

	isUserExist := len(users)

	if isUserExist == 0 {
		return nil, errors.New("USER_NOT_FOUND")
	}

	passHashed := users[0].Password

	err = helper.VerifyPassword(passHashed, loginScan.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New("CREDENTIALS_IS_INCORRECT")
	}

	token, err := middleware.CreateToken(loginScan.Id)
	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	access := token["token"]

	return map[string]any{
		"data": access,
	}, nil
}

func Register(r *entities.Register) (map[string]any, error) {

	roles := []entities.CheckRole{}

	r.UserId = uuid.NewV4().String()

	queryCheckRole := `SELECT id, name FROM user_roles WHERE id = ?`

	errCheckRole := dbDefault.Debug().Raw(queryCheckRole, r.Role).Scan(&roles).Error

	if errCheckRole != nil {
		helper.Logger("error", "In Server: "+errCheckRole.Error())
		return nil, errors.New(errCheckRole.Error())
	}

	isCheckRoleExist := len(roles)

	if isCheckRoleExist == 0 {
		helper.Logger("error", "In Server: Role not found")
		return nil, errors.New("ROLE_NOT_FOUND")
	}

	queryInsertUser := `INSERT INTO users (uid, email, phone) VALUES (?, ?, ?)`

	errInsertUser := dbDefault.Debug().Exec(queryInsertUser, r.UserId, r.Email, r.Phone).Error

	if errInsertUser != nil {
		helper.Logger("error", "In Server: "+errInsertUser.Error())
		return nil, errors.New(errInsertUser.Error())
	}

	queryInsertProfile := `INSERT INTO profiles (user_id, fullname) VALUES (?, ?)`

	err := dbDefault.Debug().Exec(queryInsertProfile, r.UserId, r.Fullname).Error

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New(err.Error())
	}

	return map[string]any{
		"data": r,
	}, nil
}
