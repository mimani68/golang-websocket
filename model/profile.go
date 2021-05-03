package model

type Profile struct {
	Id           string `json:"id,omitempty"`
	Nickname     string `json:"nickname,omitempty"`
	LoginProfile string `json:"loginProfile,omitempty"`
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	DeviceId     string `json:"deviceId,omitempty"`
}
