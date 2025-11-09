package repository

import (
	"context"
	"math"

	"github.com/pikachu0310/very-big-medal-pusher-data-server/openapi/models"
)

// GetCreditAllDistribution aggregates credit_all values from the latest save data table into digit-based ranges.
func (r *Repository) GetCreditAllDistribution(ctx context.Context) (*models.CreditAllDistributionResponse, error) {
	const totalUsersQuery = `
SELECT
  COUNT(*)
FROM v3_user_latest_save_data
WHERE hide_record = 0`

	var totalUsers int64
	if err := r.db.GetContext(ctx, &totalUsers, totalUsersQuery); err != nil {
		return nil, err
	}

	const distributionQuery = `
SELECT
  CASE
    WHEN credit_all IS NULL OR credit_all < 1000 THEN 3
    ELSE CAST(FLOOR(LOG10(credit_all)) + 1 AS SIGNED)
  END AS digits,
  COUNT(*) AS user_count
FROM v3_user_latest_save_data
WHERE hide_record = 0
GROUP BY digits
ORDER BY digits`

	var rows []struct {
		Digits    int   `db:"digits"`
		UserCount int64 `db:"user_count"`
	}
	if err := r.db.SelectContext(ctx, &rows, distributionQuery); err != nil {
		return nil, err
	}

	countsByDigits := make(map[int]int64, len(rows))
	maxDigits := 3
	for _, row := range rows {
		digits := row.Digits
		if digits < 3 {
			digits = 3
		}
		countsByDigits[digits] = row.UserCount
		if digits > maxDigits {
			maxDigits = digits
		}
	}

	if totalUsers == 0 {
		maxDigits = 3
	}

	distribution := make([]models.CreditAllDistributionBucket, 0, maxDigits-2)
	for digits := 3; digits <= maxDigits; digits++ {
		rangeMin, rangeMax := creditRangeForDigits(digits)
		distribution = append(distribution, models.CreditAllDistributionBucket{
			RangeMin: rangeMin,
			RangeMax: rangeMax,
			Users:    countsByDigits[digits],
		})
	}

	return &models.CreditAllDistributionResponse{
		Users:        totalUsers,
		Distribution: distribution,
	}, nil
}

func creditRangeForDigits(digits int) (int64, int64) {
	if digits <= 3 {
		return 0, 999
	}

	rangeMin, _ := pow10Bounded(digits - 1)
	rangeMaxCandidate, overflow := pow10Bounded(digits)
	if overflow {
		return rangeMin, math.MaxInt64
	}

	if rangeMaxCandidate == math.MaxInt64 {
		return rangeMin, math.MaxInt64
	}

	return rangeMin, rangeMaxCandidate - 1
}

func pow10Bounded(exp int) (int64, bool) {
	if exp <= 0 {
		return 1, false
	}

	result := int64(1)
	for i := 0; i < exp; i++ {
		if result > math.MaxInt64/10 {
			return math.MaxInt64, true
		}
		result *= 10
	}
	return result, false
}
