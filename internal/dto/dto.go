package dto

type UserReq struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

type UserRes struct {
	ID uint `json:"id"`
}
