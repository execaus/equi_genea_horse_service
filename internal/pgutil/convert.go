package pgutil

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func StringPtrToPgText(s *string) *pgtype.Text {
	if s == nil {
		return &pgtype.Text{Valid: false}
	}
	return &pgtype.Text{
		String: *s,
		Valid:  true,
	}
}

func UUIDPtrToPgUUID(u *uuid.UUID) pgtype.UUID {
	if u == nil {
		return pgtype.UUID{Valid: false}
	}
	return pgtype.UUID{
		Bytes: *u,
		Valid: true,
	}
}
