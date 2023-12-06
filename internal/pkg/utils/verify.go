package utils

var (
	IdVerify    = Rules{"ID": {NotEmpty()}}
	LoginVerify = Rules{
		"CaptchaId": {NotEmpty()},
		"Captcha":   {NotEmpty()},
		"Username":  {NotEmpty()},
		"Password":  {NotEmpty()},
	}
	RegisterVerify = Rules{
		"Username":    {NotEmpty()},
		"NickName":    {NotEmpty()},
		"Password":    {NotEmpty()},
		"AuthorityId": {NotEmpty()},
	}
	PageInfoVerify         = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"AuthorityId": {NotEmpty()}}
)
