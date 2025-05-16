package http

// RememberCipherLoginDataRequest defines request for vault provider.
type RememberCipherLoginDataRequest struct {
	Login    string `json:"login"`
	Meta     string `json:"meta"`
	Password string `json:"password"`
	Uri      string `json:"uri"`
}

// RememberCipherCardDataRequest defines request for vault provider.
type RememberCipherCardDataRequest struct {
	Brand          string `json:"brand"`
	CardHolderName string `json:"cardHolderName"`
	Code           string `json:"code"`
	ExpMonth       string `json:"expMonth"`
	ExpYear        string `json:"expYear"`
	Meta           string `json:"meta"`
	Number         string `json:"number"`
}

// RememberCipherCustomBinaryDataRequest  defines request for vault provider.
type RememberCipherCustomBinaryDataRequest struct {
	Key   string `json:"key"`
	Meta  string `json:"meta"`
	Value string `json:"value"`
}

// RememberCipherCustomDataRequest defines  defines request for vault provider.
type RememberCipherCustomDataRequest struct {
	Key   string `json:"key"`
	Meta  string `json:"meta"`
	Value string `json:"value"`
}
