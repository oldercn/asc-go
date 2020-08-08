package asc

import (
	"context"
	"fmt"
)

// IconAssetType defines model for IconAssetType.
//
// https://developer.apple.com/documentation/appstoreconnectapi/iconassettype
type IconAssetType string

// List of IconAssetType
const (
	IconAssetTypeAppStore         IconAssetType = "APP_STORE"
	IconAssetTypeMessagesAppStore IconAssetType = "MESSAGES_APP_STORE"
	IconAssetTypeTVOSHomeScreen   IconAssetType = "TV_OS_HOME_SCREEN"
	IconAssetTypeTVOSTopShelf     IconAssetType = "TV_OS_TOP_SHELF"
	IconAssetTypeWatchAppStore    IconAssetType = "WATCH_APP_STORE"
)

// BuildIcon defines model for BuildIcon.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildicon
type BuildIcon struct {
	Attributes *struct {
		IconAsset *ImageAsset    `json:"iconAsset,omitempty"`
		IconType  *IconAssetType `json:"iconType,omitempty"`
	} `json:"attributes,omitempty"`
	ID    string        `json:"id"`
	Links ResourceLinks `json:"links"`
	Type  string        `json:"type"`
}

// BuildIconsResponse defines model for BuildIconsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/buildiconsresponse
type BuildIconsResponse struct {
	Data  []BuildIcon        `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// ListIconsQuery are query options for ListIcons
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_icons_for_a_build
type ListIconsQuery struct {
	FieldsBuildIcons []string `url:"fields[buildIcons],omitempty"`
	Limit            int      `url:"limit,omitempty"`
	Cursor           string   `url:"cursor,omitempty"`
}

// ListIconsForBuild lists all the icons for various platforms delivered with a build.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_icons_for_a_build
func (s *BuildsService) ListIconsForBuild(ctx context.Context, id string, params *ListIconsQuery) (*BuildIconsResponse, *Response, error) {
	url := fmt.Sprintf("builds/%s/icons", id)
	res := new(BuildIconsResponse)
	resp, err := s.client.get(ctx, url, params, res)
	return res, resp, err
}
