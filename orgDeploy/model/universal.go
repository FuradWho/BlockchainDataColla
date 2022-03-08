package model

type UserInfo struct {
	Id             string `form:"id"`
	UserName       string `form:"user_name"`
	InvitationCode string `form:"invitation_code"`
	CompanyId      string `form:"company_id"`
	Iphone         string `form:"iphone"`
	Password       string `form:"password"`
}
