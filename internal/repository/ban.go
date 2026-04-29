package repository

import "context"

func (r *Repository) IsUserBanned(ctx context.Context, userID string) (bool, error) {
	var n int
	if err := r.db.GetContext(ctx, &n,
		`SELECT COUNT(*) FROM user_banned WHERE user_id = ? AND enabled = 1`,
		userID,
	); err != nil {
		return false, err
	}
	return n > 0, nil
}
