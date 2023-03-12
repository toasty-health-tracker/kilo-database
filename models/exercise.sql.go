// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: exercise.sql

package models

import (
	"context"
	"database/sql"
)

const deleteExercise = `-- name: DeleteExercise :exec
DELETE FROM tracker.exercise
WHERE NAME = $1
`

func (q *Queries) DeleteExercise(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deleteExercise, name)
	return err
}

const getExercise = `-- name: GetExercise :many
SELECT a.name, b.body_part, b.level FROM tracker.exercise a
JOIN tracker.exercise_details b
ON a.name = b.exercise_name
WHERE NAME = $1 LIMIT 1
`

type GetExerciseRow struct {
	Name     string `json:"name"`
	BodyPart string `json:"bodyPart"`
	Level    string `json:"level"`
}

func (q *Queries) GetExercise(ctx context.Context, name string) ([]GetExerciseRow, error) {
	rows, err := q.db.QueryContext(ctx, getExercise, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetExerciseRow
	for rows.Next() {
		var i GetExerciseRow
		if err := rows.Scan(&i.Name, &i.BodyPart, &i.Level); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getExercises = `-- name: GetExercises :many
SELECT name FROM tracker.exercise
LIMIT $1
`

func (q *Queries) GetExercises(ctx context.Context, limit int32) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getExercises, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const submitExercise = `-- name: SubmitExercise :one
INSERT INTO tracker.exercise (
  NAME, TYPE, VARIATION
) VALUES (
  $1, $2, $3
)
ON CONFLICT (NAME) 
DO UPDATE SET 
  TYPE = $2,
  VARIATION = $3,
  UPDT_TS = CURRENT_TIMESTAMP
RETURNING name, type, variation, cret_ts, updt_ts
`

type SubmitExerciseParams struct {
	Name      string         `json:"name"`
	Type      sql.NullString `json:"type"`
	Variation sql.NullString `json:"variation"`
}

func (q *Queries) SubmitExercise(ctx context.Context, arg SubmitExerciseParams) (TrackerExercise, error) {
	row := q.db.QueryRowContext(ctx, submitExercise, arg.Name, arg.Type, arg.Variation)
	var i TrackerExercise
	err := row.Scan(
		&i.Name,
		&i.Type,
		&i.Variation,
		&i.CretTs,
		&i.UpdtTs,
	)
	return i, err
}

const submitExerciseDetails = `-- name: SubmitExerciseDetails :one
INSERT INTO tracker.exercise_details (
  EXERCISE_NAME, BODY_PART, LEVEL
) VALUES (
  $1, $2, $3
)
ON CONFLICT (EXERCISE_NAME, BODY_PART) 
DO UPDATE SET 
  LEVEL = $3,
  UPDT_TS = CURRENT_TIMESTAMP
RETURNING exercise_name, body_part, level, cret_ts, updt_ts
`

type SubmitExerciseDetailsParams struct {
	ExerciseName string `json:"exerciseName"`
	BodyPart     string `json:"bodyPart"`
	Level        string `json:"level"`
}

func (q *Queries) SubmitExerciseDetails(ctx context.Context, arg SubmitExerciseDetailsParams) (TrackerExerciseDetail, error) {
	row := q.db.QueryRowContext(ctx, submitExerciseDetails, arg.ExerciseName, arg.BodyPart, arg.Level)
	var i TrackerExerciseDetail
	err := row.Scan(
		&i.ExerciseName,
		&i.BodyPart,
		&i.Level,
		&i.CretTs,
		&i.UpdtTs,
	)
	return i, err
}

const submitExercisePerformed = `-- name: SubmitExercisePerformed :one
INSERT INTO tracker.exercise_performed (
  SET_ID, EXERCISE_NAME, REPS, WEIGHT, REPS_IN_RESERVE
) VALUES (
  $1, $2, $3, $4, $5
)
ON CONFLICT (WORKOUT_ID, GROUP_ID) 
DO UPDATE SET 
  UPDT_TS = CURRENT_TIMESTAMP
RETURNING id, set_id, exercise_name, reps, weight, reps_in_reserve, cret_ts, updt_ts
`

type SubmitExercisePerformedParams struct {
	SetID         int32          `json:"setID"`
	ExerciseName  string         `json:"exerciseName"`
	Reps          int16          `json:"reps"`
	Weight        int16          `json:"weight"`
	RepsInReserve sql.NullString `json:"repsInReserve"`
}

func (q *Queries) SubmitExercisePerformed(ctx context.Context, arg SubmitExercisePerformedParams) (TrackerExercisePerformed, error) {
	row := q.db.QueryRowContext(ctx, submitExercisePerformed,
		arg.SetID,
		arg.ExerciseName,
		arg.Reps,
		arg.Weight,
		arg.RepsInReserve,
	)
	var i TrackerExercisePerformed
	err := row.Scan(
		&i.ID,
		&i.SetID,
		&i.ExerciseName,
		&i.Reps,
		&i.Weight,
		&i.RepsInReserve,
		&i.CretTs,
		&i.UpdtTs,
	)
	return i, err
}
