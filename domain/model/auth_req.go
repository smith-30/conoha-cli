package model

type (
	AuthReq struct {
		Auth `json:"auth"`
	}

	Auth struct {
		PasswordCredentials `json:"passwordCredentials"`
		TenantID            string `json:"tenantId"`
	}

	PasswordCredentials struct {
		Password string `json:"password"`
		UserName string `json:"username"`
	}
)

func NewAuthReq(u, p, t string) *AuthReq {
	pc := PasswordCredentials{
		Password: p,
		UserName: u,
	}
	a := Auth{
		pc,
		t,
	}

	return &AuthReq{
		a,
	}
}
