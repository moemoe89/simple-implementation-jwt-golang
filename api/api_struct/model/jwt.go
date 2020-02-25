//
//  Simple JWT
//
//  Copyright © 2016. All rights reserved.
//

package model

// JWTModel represent the jwt response API
type JWTModel struct {
	ThirdParty string `json:"third_party"`
	ExpiredTime float64 `json:"expired_time"`
}
