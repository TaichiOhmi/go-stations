package model

// A HealthzResponse expresses health check message.
// APIとしてこのメッセージが表示される。（取得される)
type HealthzResponse struct{
	Message string `json:"message"`
}
