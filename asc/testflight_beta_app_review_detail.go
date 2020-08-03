package asc

import (
	"fmt"
	"net/http"
)

// BetaAppReviewDetail defines model for BetaAppReviewDetail.
type BetaAppReviewDetail struct {
	Attributes *struct {
		ContactEmail        *string `json:"contactEmail,omitempty"`
		ContactFirstName    *string `json:"contactFirstName,omitempty"`
		ContactLastName     *string `json:"contactLastName,omitempty"`
		ContactPhone        *string `json:"contactPhone,omitempty"`
		DemoAccountName     *string `json:"demoAccountName,omitempty"`
		DemoAccountPassword *string `json:"demoAccountPassword,omitempty"`
		DemoAccountRequired *bool   `json:"demoAccountRequired,omitempty"`
		Notes               *string `json:"notes,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		App *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"app,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BetaAppReviewDetailUpdateRequest defines model for BetaAppReviewDetailUpdateRequest.
type BetaAppReviewDetailUpdateRequest struct {
	Attributes *BetaAppReviewDetailUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                      `json:"id"`
	Type       string                                      `json:"type"`
}

// BetaAppReviewDetailUpdateRequestAttributes are attributes for BetaAppReviewDetailUpdateRequest
type BetaAppReviewDetailUpdateRequestAttributes struct {
	ContactEmail        *string `json:"contactEmail,omitempty"`
	ContactFirstName    *string `json:"contactFirstName,omitempty"`
	ContactLastName     *string `json:"contactLastName,omitempty"`
	ContactPhone        *string `json:"contactPhone,omitempty"`
	DemoAccountName     *string `json:"demoAccountName,omitempty"`
	DemoAccountPassword *string `json:"demoAccountPassword,omitempty"`
	DemoAccountRequired *bool   `json:"demoAccountRequired,omitempty"`
	Notes               *string `json:"notes,omitempty"`
}

// BetaAppReviewDetailResponse defines model for BetaAppReviewDetailResponse.
type BetaAppReviewDetailResponse struct {
	Data     BetaAppReviewDetail `json:"data"`
	Included *[]App              `json:"included,omitempty"`
	Links    DocumentLinks       `json:"links"`
}

// BetaAppReviewDetailsResponse defines model for BetaAppReviewDetailsResponse.
type BetaAppReviewDetailsResponse struct {
	Data     []BetaAppReviewDetail `json:"data"`
	Included *[]App                `json:"included,omitempty"`
	Links    PagedDocumentLinks    `json:"links"`
	Meta     *PagingInformation    `json:"meta,omitempty"`
}

// ListBetaAppReviewDetailsQuery defines model for ListBetaAppReviewDetails
type ListBetaAppReviewDetailsQuery struct {
	FieldsApps                 *[]string `url:"fields[apps],omitempty"`
	FieldsBetaAppReviewDetails *[]string `url:"fields[betaAppReviewDetails],omitempty"`
	FilterApp                  *[]string `url:"filter[app],omitempty"`
	Include                    *[]string `url:"include,omitempty"`
	Limit                      *int      `url:"limit,omitempty"`
	Cursor                     *string   `url:"cursor,omitempty"`
}

// GetBetaAppReviewDetailQuery defines model for GetBetaAppReviewDetail
type GetBetaAppReviewDetailQuery struct {
	FieldsApps                 *[]string `url:"fields[apps],omitempty"`
	FieldsBetaAppReviewDetails *[]string `url:"fields[betaAppReviewDetails],omitempty"`
	Include                    *[]string `url:"include,omitempty"`
}

// GetAppForBetaAppReviewDetailQuery defines model for GetAppForBetaAppReviewDetail
type GetAppForBetaAppReviewDetailQuery struct {
	FieldsApps *[]string `url:"fields[apps],omitempty"`
}

// GetBetaAppReviewDetailsForAppQuery defines model for GetBetaAppReviewDetailsForApp
type GetBetaAppReviewDetailsForAppQuery struct {
	FieldsBetaAppReviewDetails *[]string `url:"fields[betaAppReviewDetails],omitempty"`
}

// ListBetaAppReviewDetails finds and lists beta app review details for all apps.
func (s *TestflightService) ListBetaAppReviewDetails(params *ListBetaAppReviewDetailsQuery) (*BetaAppReviewDetailsResponse, *http.Response, error) {
	res := new(BetaAppReviewDetailsResponse)
	resp, err := s.client.GetWithQuery("betaAppReviewDetails", params, res)
	return res, resp, err
}

// GetBetaAppReviewDetail gets beta app review details for a specific app.
func (s *TestflightService) GetBetaAppReviewDetail(id string, params *GetBetaAppReviewDetailQuery) (*BetaAppReviewDetailResponse, *http.Response, error) {
	url := fmt.Sprintf("betaAppReviewDetails/%s", id)
	res := new(BetaAppReviewDetailResponse)
	resp, err := s.client.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetAppForBetaAppReviewDetail gets the app information for a specific beta app review details resource.
func (s *TestflightService) GetAppForBetaAppReviewDetail(id string, params *GetAppForBetaAppReviewDetailQuery) (*AppResponse, *http.Response, error) {
	url := fmt.Sprintf("betaAppReviewDetails/%s/app", id)
	res := new(AppResponse)
	resp, err := s.client.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBetaAppReviewDetailsForApp gets the beta app review details for a specific app.
func (s *TestflightService) GetBetaAppReviewDetailsForApp(id string, params *GetBetaAppReviewDetailsForAppQuery) (*BetaAppReviewDetailResponse, *http.Response, error) {
	url := fmt.Sprintf("apps/%s/betaAppReviewDetail", id)
	res := new(BetaAppReviewDetailResponse)
	resp, err := s.client.GetWithQuery(url, params, res)
	return res, resp, err
}

// UpdateBetaAppReviewDetail updates the details for a specific app's beta app review.
func (s *TestflightService) UpdateBetaAppReviewDetail(id string, body *BetaAppReviewDetailUpdateRequest) (*BetaAppReviewDetailResponse, *http.Response, error) {
	url := fmt.Sprintf("betaAppReviewDetails/%s", id)
	res := new(BetaAppReviewDetailResponse)
	resp, err := s.client.Patch(url, body, res)
	return res, resp, err
}