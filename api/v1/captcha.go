package v1

type CaptchaResponseData struct {
	Key  string `json:"key"`
	B64s string `json:"b64s"`
}
