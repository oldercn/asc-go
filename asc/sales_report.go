package asc

import (
	"context"
)

type SalesReportQuery struct {
	FilterReportType    string `url:"filter[reportType]"`
	FilterReportSubType string `url:"filter[reportSubType]"`
	FilterVendorNumber  string `url:"filter[vendorNumber]"`
	FilterFrequency     string `url:"filter[frequency]"`
	FilterReportDate    string `url:"filter[reportDate]"`
}

type SalesReportService service

// https://api.appstoreconnect.apple.com/v1/salesReports
func (s *SalesReportService) DownloadSalesReport(ctx context.Context, params *SalesReportQuery) (*Response, error) {
	resp, err := s.client.download(ctx, "salesReports", params, withAccept("application/a-gzip"))

	return resp, err
}
