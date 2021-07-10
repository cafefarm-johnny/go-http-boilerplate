package users

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go-http-boilerplate/src/domain"
	"go-http-boilerplate/src/domain/users/model"
	"go-http-boilerplate/src/hash"
	"net/http"
)

type UserService struct {
	ur *userRepository
}

func NewUserService() *UserService {
	return &UserService{
		ur: NewUserRepository(),
	}
}

func (us *UserService) Signup(ctx echo.Context, dto *model.UserDTO) error {
	hashPwd, err := hash.ToBcrypt(dto.Password)
	if err != nil {
		ctx.Logger().Error("Signup: password encrypt: ", err)
		return domain.ErrInternalServerError
	}
	return us.ur.Save(dto.Username, hashPwd)
}

func (us *UserService) Signin(ctx echo.Context, dto *model.UserDTO) error {
	sess, err := session.Get(domain.SessionName, ctx)
	if err != nil {
		ctx.Logger().Error("Signin: get session: ", err)
		return domain.ErrInternalServerError
	}
	if hasSession(sess) {
		return domain.ErrDuplicatedLogin
	}

	u := us.ur.Find(dto.Username)
	if u == nil {
		return domain.ErrNotFoundUser
	}

	if !hash.CompareBcrypt(u.Password, []byte(dto.Password)) {
		return domain.ErrInvalidPassword
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   86400 * 7, // 7일
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}

	sess.Values[domain.SessionKey] = domain.NewSession(u.UUID, u.Username)
	if err = sess.Save(ctx.Request(), ctx.Response()); err != nil {
		ctx.Logger().Error("Signin: session create: ", err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (us *UserService) SelfAuthenticate(dto *model.UserDTO) error {
	u := us.ur.Find(dto.Username)
	if u == nil {
		return domain.ErrNotFoundUser
	}

	if !hash.CompareBcrypt(u.Password, []byte(dto.Password)) {
		return domain.ErrInvalidPassword
	}

	return nil
}

func (us *UserService) Logout(ctx echo.Context) error {
	sess, err := session.Get(domain.SessionName, ctx)
	if err != nil {
		ctx.Logger().Error(err)
		return domain.ErrForbidden
	}
	if !hasSession(sess) {
		return domain.ErrUnauthorized
	}

	sess.Values[domain.SessionKey] = ""
	if err := sess.Save(ctx.Request(), ctx.Response()); err != nil {
		ctx.Logger().Error(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func hasSession(sess *sessions.Session) bool {
	if sess.Values[domain.SessionKey] == nil {
		return false
	}
	return true
}
