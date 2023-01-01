package jwt

import (
	"time"

	"github.com/0xThomas3000/food_delivery/component/tokenprovider"
	"github.com/golang-jwt/jwt/v4"
)

type jwtProvider struct {
	secret string // Khóa bí mật(dc dùng để Generate ra dc Token / hoặc mới xác thực dc cái Token), mã hóa đối xứng (chỉ dùng 1 key)
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	// myClaims mặc định cũng có hàm Valid() của "RegisteredClaims" vì nó đã embed "RegisteredClaims"
	jwt.RegisteredClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	// generate the JWT
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Second * time.Duration(expiry))),
			IssuedAt:  jwt.NewNumericDate(time.Now().Local()),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret)) // Đổi key thành mảng []byte và truyền vào SignedString()
	if err != nil {
		return nil, err
	}

	// return the token
	return &tokenprovider.Token{
		Token:     myToken,
		Expiry:    expiry,
		CreatedAt: time.Now(),
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (*tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrInvalidToken
	}

	// validate the token
	if !res.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, tokenprovider.ErrInvalidToken
	}

	// return the token
	return &claims.Payload, nil
}

func (j *jwtProvider) String() string {
	return "JWT implement Provider"
}
