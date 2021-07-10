package users

import (
	"github.com/labstack/echo/v4"
	"go-http-boilerplate/src/domain/users/model"
	"net/http"
)

type userController struct {
	us *UserService
}

func NewUserController() *userController {
	return &userController{
		us: NewUserService(),
	}
}

func (uc *userController) Signup(ctx echo.Context) error {
	dto := model.NewUserDTO()
	if err := ctx.Bind(dto); err != nil {
		return err
	}

	if err := ctx.Validate(dto); err != nil {
		return err
	}

	if err := uc.us.Signup(ctx, dto); err != nil {
		return err
	}

	return ctx.String(http.StatusCreated, "회원 가입 성공")
}

func (uc *userController) Signin(ctx echo.Context) error {
	dto := model.NewUserDTO()
	if err := ctx.Bind(dto); err != nil {
		return err
	}

	if err := ctx.Validate(dto); err != nil {
		return err
	}

	if err := uc.us.Signin(ctx, dto); err != nil {
		return err
	}

	return ctx.String(http.StatusOK, "로그인 성공")
}

func (uc *userController) SelfAuthenticate(ctx echo.Context) error {
	dto := model.NewUserDTO()
	if err := ctx.Bind(dto); err != nil {
		return err
	}

	if err := ctx.Validate(dto); err != nil {
		return err
	}

	if err := uc.us.SelfAuthenticate(dto); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (uc *userController) Logout(ctx echo.Context) error {
	if err := uc.us.Logout(ctx); err != nil {
		return err
	}

	return ctx.String(http.StatusOK, "로그아웃 성공")
}
