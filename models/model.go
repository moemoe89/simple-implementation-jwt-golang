package models

type JWTModel struct {
	ThirdParty string `json:"third_party"`
	ExpiredTime float64 `json:"expired_time"`
}
