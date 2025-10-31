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

func NullableStringToPgText(s *string) pgtype.Text {
	var desc pgtype.Text
	if s != nil {
		desc.String = *s
		desc.Valid = true
	} else {
		desc.Valid = false
	}
	return desc
}

func UUIDToPgUUID(uuid uuid.UUID) *pgtype.UUID {
	return &pgtype.UUID{
		Bytes: uuid,
		Valid: true,
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
