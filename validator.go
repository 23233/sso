package sso

type CheckTicketReq struct {
	Ticket  string `json:"ticket" url:"ticket" form:"ticket" comment:"标识符" validate:"required,max=36,min=32"`
	NextUrl string `json:"next_url" url:"next_url" form:"next_url" comment:"跳转地址" validate:""`
}
