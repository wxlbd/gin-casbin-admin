package captcha

type CaptchaResponse struct {
	CaptchaId    string `json:"captchaId"`
	CaptchaImage string `json:"captchaImage"`
}
