package domain

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// 400
var (
	ErrBadRequest = echo.NewHTTPError(
		http.StatusBadRequest,
		newDomainErr(400001, "잘못된 요청입니다."),
	)
)

// 401
var (
	ErrDuplicatedUser = echo.NewHTTPError(
		http.StatusUnauthorized,
		newDomainErr(401001, "이미 존재하는 회원입니다."),
	)
	ErrInvalidPassword = echo.NewHTTPError(
		http.StatusUnauthorized,
		newDomainErr(401002, "패스워드가 일치하지 않습니다."),
	)
	ErrUnauthorized = echo.NewHTTPError(
		http.StatusUnauthorized,
		newDomainErr(401003, "로그인 후 시도해주세요."),
	)
)

// 403
var (
	ErrDuplicatedLogin = echo.NewHTTPError(
		http.StatusForbidden,
		newDomainErr(403001, "중복 로그인 요청입니다."),
	)
	ErrForbidden = echo.NewHTTPError(
		http.StatusForbidden,
		newDomainErr(403002, "잘못된 접근입니다."),
	)
)

// 404
var (
	ErrNotFoundUser = echo.NewHTTPError(
		http.StatusNotFound,
		newDomainErr(404001, "존재하지 않는 회원입니다."),
	)
)

// 500
var (
	ErrInternalServerError = echo.NewHTTPError(
		http.StatusInternalServerError,
		newDomainErr(500001, "알 수 없는 에러가 발생하였습니다."),
	)
)

type domainErr struct {
	Code    int
	Message string
}

// error 인터페이스 덕타이핑
func (d *domainErr) Error() string {
	return fmt.Sprintf("{ code: %d, message: %s }", d.Code, d.Message)
}

func newDomainErr(code int, message string) *domainErr {
	return &domainErr{
		code,
		message,
	}
}
