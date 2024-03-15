package dto

type AuthenticationTokensDto struct {
	Id          int    `json:"id"`
	AccessToken string `json:"accessToken"`
}
