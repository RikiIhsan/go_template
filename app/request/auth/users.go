package auth

type (
	//data for a signup user request
	UserSignUp struct {
		Username string `json:"username" validate:"required,max=10,alphanum"`
		Password string `json:"password" validate:"required,min=8,max=18"`
		Name     string `json:"name" validate:"required,max=255"`
		Email    string `json:"email" validate:"required,email"`
	}
	//data for a signin user request
	UserSignIn struct {
		Username string `json:"username" validate:"required,max=10,alphanum"`
		Password string `json:"password" validate:"required,min=8,max=18"`
	}
	//data for a register a thridparty application
	RegisterKey struct {
		Detail string `json:"detail" validate:"required,max=255,min=10"`
	}
)
