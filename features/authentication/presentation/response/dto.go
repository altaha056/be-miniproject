package response

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
}

func ToAuthResponse(accessToken, refreshToken string) AuthResponse {
	return AuthResponse{
		AccessToken:  accessToken,
	}
}