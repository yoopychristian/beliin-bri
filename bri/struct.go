package bri

type GetTokenResponse struct {
	RefreshTokenExpiresIn string   `json:"refresh_token_expires_in"`
	APIProductList        string   `json:"api_product_list"`
	APIProductListJSON    []string `json:"api_product_list_json"`
	OrganizationName      string   `json:"organization_name"`
	DeveloperEmail        string   `json:"developer.email"`
	TokenType             string   `json:"token_type"`
	IssuedAt              string   `json:"issued_at"`
	ClientID              string   `json:"client_id"`
	AccessToken           string   `json:"access_token"`
	ApplicationName       string   `json:"application_name"`
	Scope                 string   `json:"scope"`
	ExpiresIn             string   `json:"expires_in"`
	RefreshCount          string   `json:"refresh_count"`
	Status                string   `json:"status"`
}

type Invoice struct {
	InstitutionCode string `json:"institutionCode"`
	BrivaNo         string `json:"brivaNo"`
	CustCode        string `json:"custCode"`
	Nama            string `json:"nama"`
	Amount          string `json:"amount"`
	Keterangan      string `json:"keterangan"`
	ExpiredDate     string `json:"expiredDate"`
}

type InvoiceResponse struct {
	Status              bool   `json:"status"`
	ResponseDescription string `json:"responseDescription"`
	ResponseCode        string `json:"responseCode"`
	Data                struct {
		InstitutionCode string `json:"institutionCode"`
		BrivaNo         string `json:"brivaNo"`
		CustCode        string `json:"custCode"`
		Nama            string `json:"nama"`
		Amount          string `json:"amount"`
		Keterangan      string `json:"keterangan"`
		ExpiredDate     string `json:"expiredDate"`
	} `json:"data"`
}

type GetStatusResponse struct {
	Status              bool   `json:"status"`
	ResponseDescription string `json:"responseDescription"`
	ResponseCode        string `json:"responseCode"`
	Data                struct {
		StatusBayar string `json:"statusBayar"`
	} `json:"data"`
}
