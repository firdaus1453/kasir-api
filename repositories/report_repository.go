package repositories

import (
	"database/sql"
	"time"

	"kasir-api/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetSalesReport(startDate, endDate time.Time) (*models.SalesReport, error) {
	report := &models.SalesReport{}

	// Get total revenue and total transactions
	summaryQuery := `
		SELECT COALESCE(SUM(total_amount), 0), COUNT(*)
		FROM transactions
		WHERE created_at >= $1 AND created_at < $2`

	err := repo.db.QueryRow(summaryQuery, startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// Get best selling product
	bestSellerQuery := `
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as total_qty
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE t.created_at >= $1 AND t.created_at < $2
		GROUP BY p.id, p.name
		ORDER BY total_qty DESC
		LIMIT 1`

	err = repo.db.QueryRow(bestSellerQuery, startDate, endDate).Scan(&report.ProdukTerlaris.Nama, &report.ProdukTerlaris.QtyTerjual)
	if err == sql.ErrNoRows {
		report.ProdukTerlaris = models.ProdukTerlaris{Nama: "-", QtyTerjual: 0}
	} else if err != nil {
		return nil, err
	}

	return report, nil
}
