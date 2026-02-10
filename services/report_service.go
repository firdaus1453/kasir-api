package services

import (
	"time"

	"kasir-api/models"
	"kasir-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetReportHariIni() (*models.SalesReport, error) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1)
	return s.repo.GetSalesReport(startOfDay, endOfDay)
}

func (s *ReportService) GetReportByDateRange(startDate, endDate time.Time) (*models.SalesReport, error) {
	// endDate should be inclusive, so add 1 day
	endDate = endDate.AddDate(0, 0, 1)
	return s.repo.GetSalesReport(startDate, endDate)
}
