package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type AuditLog struct {
	Id         uuid.UUID       `db:"id"`
	UserId     *uuid.UUID      `db:"user_id"`
	Action     string          `db:"action"`
	EntityType string          `db:"entity_type"`
	EntityId   uuid.UUID       `db:"entity_id"`
	Meta       json.RawMessage `db:"meta"`
	IpAddress  string          `db:"ip_address"`
	CreatedAt  time.Time       `db:"created_at"`
}
