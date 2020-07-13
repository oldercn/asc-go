package testflight

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/builds"
	"github.com/aaronsky/asc-go/v1/asc/common"
)

// ExternalBetaState defines model for ExternalBetaState.
type ExternalBetaState string

// List of ExternalBetaState
const (
	ExternalBetaApproved                 ExternalBetaState = "BETA_APPROVED"
	ExternalBetaRejected                 ExternalBetaState = "BETA_REJECTED"
	ExternalBetaExpired                  ExternalBetaState = "EXPIRED"
	ExternalBetaInReview                 ExternalBetaState = "IN_BETA_REVIEW"
	ExternalBetaInTesting                ExternalBetaState = "IN_BETA_TESTING"
	ExternalBetaInExportComplianceReview ExternalBetaState = "IN_EXPORT_COMPLIANCE_REVIEW"
	ExternalBetaMissingExportCompliance  ExternalBetaState = "MISSING_EXPORT_COMPLIANCE"
	ExternalBetaProcessing               ExternalBetaState = "PROCESSING"
	ExternalBetaProcessingException      ExternalBetaState = "PROCESSING_EXCEPTION"
	ExternalBetaReadyForBetaSubmission   ExternalBetaState = "READY_FOR_BETA_SUBMISSION"
	ExternalBetaReadyForBetaTesting      ExternalBetaState = "READY_FOR_BETA_TESTING"
	ExternalBetaWaitingForBetaReview     ExternalBetaState = "WAITING_FOR_BETA_REVIEW"
)

// InternalBetaState defines model for InternalBetaState.
type InternalBetaState string

// List of InternalBetaState
const (
	InternalBetaExpired                  InternalBetaState = "EXPIRED"
	InternalBetaInTesting                InternalBetaState = "IN_BETA_TESTING"
	InternalBetaInExportComplianceReview InternalBetaState = "IN_EXPORT_COMPLIANCE_REVIEW"
	InternalBetaMissingExportCompliance  InternalBetaState = "MISSING_EXPORT_COMPLIANCE"
	InternalBetaProcessing               InternalBetaState = "PROCESSING"
	InternalBetaProcessingException      InternalBetaState = "PROCESSING_EXCEPTION"
	InternalBetaReadyForBetaTesting      InternalBetaState = "READY_FOR_BETA_TESTING"
)

// BuildBetaDetail defines model for BuildBetaDetail.
type BuildBetaDetail struct {
	Attributes *struct {
		AutoNotifyEnabled  *bool              `json:"autoNotifyEnabled,omitempty"`
		ExternalBuildState *ExternalBetaState `json:"externalBuildState,omitempty"`
		InternalBuildState *InternalBetaState `json:"internalBuildState,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string               `json:"id"`
	Links         common.ResourceLinks `json:"links"`
	Relationships *struct {
		Build *struct {
			Data *struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"data,omitempty"`
			Links *struct {
				Related *string `json:"related,omitempty"`
				Self    *string `json:"self,omitempty"`
			} `json:"links,omitempty"`
		} `json:"build,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// BuildBetaDetailUpdateRequest defines model for BuildBetaDetailUpdateRequest.
type BuildBetaDetailUpdateRequest struct {
	Data struct {
		Attributes *struct {
			AutoNotifyEnabled *bool `json:"autoNotifyEnabled,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// BuildBetaDetailResponse defines model for BuildBetaDetailResponse.
type BuildBetaDetailResponse struct {
	Data     BuildBetaDetail      `json:"data"`
	Included *[]builds.Build      `json:"included,omitempty"`
	Links    common.DocumentLinks `json:"links"`
}

// BuildBetaDetailsResponse defines model for BuildBetaDetailsResponse.
type BuildBetaDetailsResponse struct {
	Data     []BuildBetaDetail         `json:"data"`
	Included *[]builds.Build           `json:"included,omitempty"`
	Links    common.PagedDocumentLinks `json:"links"`
	Meta     *common.PagingInformation `json:"meta,omitempty"`
}

type ListBuildBetaDetailsQuery struct {
	Fields *struct {
		BuildBetaDetails *[]string `url:"buildBetaDetails,omitempty"`
		Builds           *[]string `url:"builds,omitempty"`
	} `url:"fields,omitempty"`
	Filter *struct {
		ID    *[]string `url:"id,omitempty"`
		Build *[]string `url:"build,omitempty"`
	} `url:"filter,omitempty"`
	Include *[]string `url:"include,omitempty"`
	Limit   *int      `url:"limit,omitempty"`
	Cursor  *string   `url:"cursor,omitempty"`
}

type GetBuildBetaDetailsQuery struct {
	Fields *struct {
		Builds           *[]string `url:"builds,omitempty"`
		BuildBetaDetails *[]string `url:"buildBetaDetails,omitempty"`
	} `url:"fields,omitempty"`
	Include *[]string `url:"include,omitempty"`
}

type GetBuildForBuildBetaDetailQuery struct {
	Fields *struct {
		Builds *[]string `url:"builds,omitempty"`
	} `url:"fields,omitempty"`
}

type GetBuildBetaDetailForBuildQuery struct {
	Fields *struct {
		BuildBetaDetails *[]string `url:"buildBetaDetails,omitempty"`
	} `url:"fields,omitempty"`
}

// ListBuildBetaDetails finds and lists build beta details for all builds.
func (s *Service) ListBuildBetaDetails(params *ListBuildBetaDetailsQuery) (*BuildBetaDetailsResponse, *common.Response, error) {
	res := new(BuildBetaDetailsResponse)
	resp, err := s.GetWithQuery("buildBetaDetails", params, res)
	return res, resp, err
}

// GetBuildBetaDetail gets a specific build beta details resource.
func (s *Service) GetBuildBetaDetail(id string, params *GetBuildBetaDetailsQuery) (*BuildBetaDetailResponse, *common.Response, error) {
	url := fmt.Sprintf("buildBetaDetails/%s", id)
	res := new(BuildBetaDetailResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBuildForBuildBetaDetail gets the build information for a specific build beta details resource.
func (s *Service) GetBuildForBuildBetaDetail(id string, params *GetBuildForBuildBetaDetailQuery) (*builds.BuildResponse, *common.Response, error) {
	url := fmt.Sprintf("buildBetaDetails/%s/build", id)
	res := new(builds.BuildResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// GetBuildBetaDetailForBuild gets the beta test details for a specific build.
func (s *Service) GetBuildBetaDetailForBuild(id string, params *GetBuildBetaDetailForBuildQuery) (*BuildBetaDetailResponse, *common.Response, error) {
	url := fmt.Sprintf("builds/%s/buildBetaDetail", id)
	res := new(BuildBetaDetailResponse)
	resp, err := s.GetWithQuery(url, params, res)
	return res, resp, err
}

// UpdateBuildBetaDetail updates beta test details for a specific build.
func (s *Service) UpdateBuildBetaDetail(id string, body *BuildBetaDetailUpdateRequest) (*BuildBetaDetailResponse, *common.Response, error) {
	url := fmt.Sprintf("buildBetaDetails/%s", id)
	res := new(BuildBetaDetailResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}
