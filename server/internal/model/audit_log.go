package model

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID           uuid.UUID  `json:"id"`
	ActorID      *uuid.UUID `json:"actorId,omitempty"`
	Action       string     `json:"action"`
	TargetUserID *uuid.UUID `json:"targetUserId,omitempty"`
	ClientID     *uuid.UUID `json:"clientId,omitempty"`
	IPAddress    *string    `json:"ipAddress,omitempty"`
	UserAgent    *string    `json:"userAgent,omitempty"`
	CreatedAt    time.Time  `json:"createdAt"`
}

func (a *AuditLog) TableName() string {
	return "iam_audit_logs"
}
