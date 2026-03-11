package repository

import (
	"gin-web-api-go/internal/model"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuditLogRepository struct {
	db *sqlx.DB
}

func NewAuditLogRepository(db *sqlx.DB) AuditLogRepository {
	return AuditLogRepository{
		db: db,
	}
}

func (r AuditLogRepository) GetAll(page int, limit int, user_id uuid.UUID, action string, to time.Time, from time.Time) ([]model.AuditLog, error) {
	var auditLogs []model.AuditLog
	page = (page - 1) * limit
	query := sq.Select("*").
		From("audit_logs").
		Limit(uint64(limit)).
		Offset(uint64(page)).
		PlaceholderFormat(sq.Dollar)

	if action != "" {
		query = query.Where(sq.Eq{"action": action})
	}
	if len(user_id) != 0 {
		query = query.Where(sq.Eq{"user_id": user_id})
	}
	if !from.IsZero() {
		query = query.Where(sq.GtOrEq{"created_at": from})
	}
	if !to.IsZero() {
		query = query.Where(sq.LtOrEq{"created_at": to})
	}

	sql, args, err := query.ToSql()
	err = r.db.Select(&auditLogs, sql, args...)

	return auditLogs, err
}
