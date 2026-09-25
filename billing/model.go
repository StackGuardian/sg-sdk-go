package billing

// CreditBalance is one line of an organization's credit balance.
type CreditBalance struct {
	Total            *float64 `json:"total,omitempty" url:"-"`
	Balance          *float64 `json:"balance,omitempty" url:"-"`
	Consumed         *float64 `json:"consumed,omitempty" url:"-"`
	ExtraCreditsUsed *float64 `json:"extra_credits_used,omitempty" url:"-"`
}

type Balance struct {
	SgCredits       *CreditBalance `json:"sg_credits,omitempty" url:"-"`
	ActiveWorkflows *CreditBalance `json:"active_workflows,omitempty" url:"-"`
	CloudAccounts   *CreditBalance `json:"cloud_accounts,omitempty" url:"-"`
}

type BalanceResponse struct {
	Msg *Balance `json:"msg,omitempty" url:"-"`
}

type DashboardRequest struct {
	// Dashboard to embed: invoices (default) or usage
	Type *string `json:"-" url:"type,omitempty"`
}

type DashboardUrl struct {
	Url string `json:"url" url:"-"`
}

type DashboardResponse struct {
	Msg *DashboardUrl `json:"msg,omitempty" url:"-"`
}

type ListInvoicesRequest struct {
	// Page size between 1 and 100
	Limit *int `json:"-" url:"limit,omitempty"`
	// Cursor returned in the previous response's next_page
	NextPage *string `json:"-" url:"next_page,omitempty"`
}

type Invoice struct {
	Id                string   `json:"id" url:"-"`
	ExternalInvoiceId *string  `json:"external_invoice_id,omitempty" url:"-"`
	Status            *string  `json:"status,omitempty" url:"-"`
	Total             *float64 `json:"total,omitempty" url:"-"`
	Currency          *string  `json:"currency,omitempty" url:"-"`
	IssuedAt          *string  `json:"issued_at,omitempty" url:"-"`
	StartTimestamp    *string  `json:"start_timestamp,omitempty" url:"-"`
	EndTimestamp      *string  `json:"end_timestamp,omitempty" url:"-"`
}

type InvoiceList struct {
	Invoices []*Invoice `json:"invoices,omitempty" url:"-"`
	NextPage *string    `json:"next_page,omitempty" url:"-"`
}

type InvoicesResponse struct {
	Msg *InvoiceList `json:"msg,omitempty" url:"-"`
}

type Address struct {
	Line1      *string `json:"line1,omitempty" url:"-"`
	Line2      *string `json:"line2,omitempty" url:"-"`
	City       *string `json:"city,omitempty" url:"-"`
	State      *string `json:"state,omitempty" url:"-"`
	PostalCode *string `json:"postal_code,omitempty" url:"-"`
	Country    *string `json:"country,omitempty" url:"-"`
}

type PaymentMethod struct {
	Id        string  `json:"id" url:"-"`
	Brand     *string `json:"brand,omitempty" url:"-"`
	Last4     *string `json:"last4,omitempty" url:"-"`
	ExpMonth  *int    `json:"exp_month,omitempty" url:"-"`
	ExpYear   *int    `json:"exp_year,omitempty" url:"-"`
	IsDefault *bool   `json:"is_default,omitempty" url:"-"`
}

type BillingInformation struct {
	Name                 *string        `json:"name,omitempty" url:"-"`
	Email                *string        `json:"email,omitempty" url:"-"`
	Address              *Address       `json:"address,omitempty" url:"-"`
	DefaultPaymentMethod *PaymentMethod `json:"default_payment_method,omitempty" url:"-"`
	AzureSubscriptionId  *string        `json:"azure_subscription_id,omitempty" url:"-"`
}

type BillingDetails struct {
	PackageName         *string             `json:"package_name,omitempty" url:"-"`
	StartingAt          *string             `json:"starting_at,omitempty" url:"-"`
	EndingBefore        *string             `json:"ending_before,omitempty" url:"-"`
	IsFreeTrial         *bool               `json:"is_free_trial,omitempty" url:"-"`
	FreeTrialExpiryDate *string             `json:"free_trial_expiry_date,omitempty" url:"-"`
	BillingInformation  *BillingInformation `json:"billing_information,omitempty" url:"-"`
}

type BillingDetailsResponse struct {
	Msg *BillingDetails `json:"msg,omitempty" url:"-"`
}

// ChangePlanRequest selects the target package by id or by display name (for example "SG | Standard").
type ChangePlanRequest struct {
	TargetPackageId *string `json:"TargetPackageId,omitempty" url:"-"`
	TargetPackage   *string `json:"TargetPackage,omitempty" url:"-"`
}

type ChangePlanResult struct {
	Message       *string `json:"message,omitempty" url:"-"`
	NewContractId *string `json:"new_contract_id,omitempty" url:"-"`
}

type ChangePlanResponse struct {
	Msg *ChangePlanResult `json:"msg,omitempty" url:"-"`
}

// BillingProfileRequest carries the Stripe customer details. Address is passed to Stripe verbatim
// (line1, line2, city, state, postal_code, country).
type BillingProfileRequest struct {
	Email   *string                `json:"Email,omitempty" url:"-"`
	Name    *string                `json:"Name,omitempty" url:"-"`
	Address map[string]interface{} `json:"Address,omitempty" url:"-"`
}

type SetupBillingProfileResult struct {
	StripeCustomerId *string `json:"stripe_customer_id,omitempty" url:"-"`
	Message          *string `json:"message,omitempty" url:"-"`
}

type SetupBillingProfileResponse struct {
	Msg *SetupBillingProfileResult `json:"msg,omitempty" url:"-"`
}

type UpdateBillingProfileResult struct {
	Message       *string  `json:"message,omitempty" url:"-"`
	UpdatedFields []string `json:"updated_fields,omitempty" url:"-"`
}

type UpdateBillingProfileResponse struct {
	Msg *UpdateBillingProfileResult `json:"msg,omitempty" url:"-"`
}

type PaymentMethodList struct {
	PaymentMethods []*PaymentMethod `json:"payment_methods,omitempty" url:"-"`
}

type PaymentMethodsResponse struct {
	Msg *PaymentMethodList `json:"msg,omitempty" url:"-"`
}

// SetupIntent carries the Stripe SetupIntent client secret used by Stripe.js to add a card.
type SetupIntent struct {
	ClientSecret *string `json:"client_secret,omitempty" url:"-"`
}

type SetupIntentResponse struct {
	Msg *SetupIntent `json:"msg,omitempty" url:"-"`
}

type PaymentMethodRequest struct {
	PaymentMethodId string `json:"PaymentMethodId" url:"-"`
}

type BillingMessage struct {
	Message *string `json:"message,omitempty" url:"-"`
}

type BillingMessageResponse struct {
	Msg *BillingMessage `json:"msg,omitempty" url:"-"`
}
