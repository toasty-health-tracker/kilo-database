// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: body_parts.sql

package models

import (
	"context"
)

const submitBodyPart = `-- name: SubmitBodyPart :one
INSERT INTO tracker.body_parts (
  NAME, REGION, UPPER_OR_LOWER
) VALUES (
  $1, $2, $3
)
ON CONFLICT (NAME) 
DO UPDATE SET 
  REGION = $2,
  UPPER_OR_LOWER = $3,
  UPDT_TS = CURRENT_TIMESTAMP
RETURNING name, region, upper_or_lower, cret_ts, updt_ts
`

type SubmitBodyPartParams struct {
	Name         string `json:"name"`
	Region       string `json:"region"`
	UpperOrLower string `json:"upperOrLower"`
}

func (q *Queries) SubmitBodyPart(ctx context.Context, arg SubmitBodyPartParams) (TrackerBodyPart, error) {
	row := q.db.QueryRowContext(ctx, submitBodyPart, arg.Name, arg.Region, arg.UpperOrLower)
	var i TrackerBodyPart
	err := row.Scan(
		&i.Name,
		&i.Region,
		&i.UpperOrLower,
		&i.CretTs,
		&i.UpdtTs,
	)
	return i, err
}
