package response

type User[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *T     `json:"data"`
}

type LogIn struct {
	UserID                  int               `json:"user_id"`
	UserName                string            `json:"user_name"`
	Lang                    string            `json:"lang"`
	CountryCode             string            `json:"country_code"`
	PhoneNumber             string            `json:"phone_number"`
	IsWhatsapp              bool              `json:"is_whatsapp"`
	Email                   string            `json:"email"`
	BindEmail               bool              `json:"bind_email"`
	BindMobile              bool              `json:"bind_mobile"`
	BindTotp                bool              `json:"bind_totp"`
	KycStatus               string            `json:"kyc_status"`
	LoginPasswordLevel      string            `json:"login_password_level"`
	LoginPasswordUpdateTime int               `json:"login_password_update_time"`
	OrderConfirmList        *OrderConfirmList `json:"order_confirm_list"`
	BindGoogleID            bool              `json:"bind_google_id"`
	BindAppleID             bool              `json:"bind_apple_id"`
	BindFacebookID          bool              `json:"bind_facebook_id"`
	BindTelegramID          bool              `json:"bind_telegram_id"`
	Token                   string            `json:"token"`
	ExpireTime              float64           `json:"expire_time"`
}

type OrderConfirmList struct {
	Limit      bool `json:"limit"`
	Market     bool `json:"market"`
	StopLimit  bool `json:"stop_limit"`
	StopMarket bool `json:"stop_market"`
}
