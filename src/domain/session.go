package domain

import "encoding/gob"

const (
	SessionName = "my-session"
	SessionKey  = "client"
)

// echo/session은 내부적으로 gorilla/sessions를 사용하고 있다.
// gorilla/sessions는 encoding/gob 패키지를 통해 직렬화 처리를 수행하는데
// 복잡한(기본 데이터 타입이 아닌) 데이터 타입을 세션에 저장하려면
// gob에 해당 구조체 정보(필드 정보들)를 등록해야 직렬화를 올바르게 처리할 수 있다.
// 추가로 쿠키 기반의 세션은 큰 객체를 저장할수 없는데, 쿠키는 4KB의 제한을 가지기 때문이다.
func init() {
	gob.Register(&session{})
}

type session struct {
	UUID     string
	Username string
}

func NewSession(uuid, username string) *session {
	return &session{
		UUID:     uuid,
		Username: username,
	}
}
