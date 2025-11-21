package auth

import "auth-service/generated/api"

type RegisterInput struct {
	api.ApiV1PostRegisterRequestBody
}

type LoginInput struct {
	api.ApiV1PostLoginRequestBody
}
