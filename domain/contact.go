package domain

type Contact struct {
	Id                   int64
	PhoneNumber          string
	EmergencyPhoneNumber string
	Email                string
}

type ContactRequest struct {
	PhoneNumber          string
	EmergencyPhoneNumber string
	Email                string
}
