// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: workout.sql

package models

import (
	"context"
	"database/sql"
	"time"
)

const deleteWorkout = `-- name: DeleteWorkout :exec
DELETE FROM tracker.workout
WHERE NAME = $1
`

func (q *Queries) DeleteWorkout(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deleteWorkout, name)
	return err
}

const deleteWorkoutPerformed = `-- name: DeleteWorkoutPerformed :exec
DELETE FROM tracker.workout_performed
WHERE SUBMITTED_ON = $1
`

func (q *Queries) DeleteWorkoutPerformed(ctx context.Context, submittedOn time.Time) error {
	_, err := q.db.ExecContext(ctx, deleteWorkoutPerformed, submittedOn)
	return err
}

const getWorkout = `-- name: GetWorkout :one
SELECT name, program_name, cret_ts, updt_ts FROM tracker.workout
WHERE NAME = $1 LIMIT 1
`

func (q *Queries) GetWorkout(ctx context.Context, name string) (TrackerWorkout, error) {
	row := q.db.QueryRowContext(ctx, getWorkout, name)
	var i TrackerWorkout
	err := row.Scan(
		&i.Name,
		&i.ProgramName,
		&i.CretTs,
		&i.UpdtTs,
	)
	return i, err
}

const getWorkoutNames = `-- name: GetWorkoutNames :many
SELECT NAME FROM tracker.workout
LIMIT $1
`

func (q *Queries) GetWorkoutNames(ctx context.Context, limit int32) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getWorkoutNames, limit)
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

const getWorkoutPerformed = `-- name: GetWorkoutPerformed :many
with workout as (
	select id, submitted_on, workout_name, rank() over(partition by submitted_on order by id desc) as rnk
	from tracker.workout_performed
)
select a.submitted_on, a.workout_name, b.group_id, b.set_number, c.exercise_name, c.reps, c.weight, c.reps_in_reserve 
from workout a
join tracker.set_performed b
	on a.id = b.workout_id
join tracker.exercise_performed c
	on b.id = c.set_id
WHERE a.submitted_on = $1
and a.rnk = 1
`

type GetWorkoutPerformedRow struct {
	SubmittedOn   time.Time      `json:"submittedOn"`
	WorkoutName   string         `json:"workoutName"`
	GroupID       int16          `json:"groupID"`
	SetNumber     int16          `json:"setNumber"`
	ExerciseName  string         `json:"exerciseName"`
	Reps          int16          `json:"reps"`
	Weight        int16          `json:"weight"`
	RepsInReserve sql.NullString `json:"repsInReserve"`
}

func (q *Queries) GetWorkoutPerformed(ctx context.Context) ([]GetWorkoutPerformedRow, error) {
	rows, err := q.db.QueryContext(ctx, getWorkoutPerformed)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetWorkoutPerformedRow
	for rows.Next() {
		var i GetWorkoutPerformedRow
		if err := rows.Scan(
			&i.SubmittedOn,
			&i.WorkoutName,
			&i.GroupID,
			&i.SetNumber,
			&i.ExerciseName,
			&i.Reps,
			&i.Weight,
			&i.RepsInReserve,
		); err != nil {
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

const submitWorkout = `-- name: SubmitWorkout :one
INSERT INTO tracker.workout (
  NAME, PROGRAM_NAME
) VALUES (
  $1, $2
)
ON CONFLICT (NAME) 
DO UPDATE SET 
  PROGRAM_NAME = $2,
  UPDT_TS = CURRENT_TIMESTAMP
RETURNING name, program_name, cret_ts, updt_ts
`

type SubmitWorkoutParams struct {
	Name        string `json:"name"`
	ProgramName string `json:"programName"`
}

func (q *Queries) SubmitWorkout(ctx context.Context, arg SubmitWorkoutParams) (TrackerWorkout, error) {
	row := q.db.QueryRowContext(ctx, submitWorkout, arg.Name, arg.ProgramName)
	var i TrackerWorkout
	err := row.Scan(
		&i.Name,
		&i.ProgramName,
		&i.CretTs,
		&i.UpdtTs,
	)
	return i, err
}

const submitWorkoutDetails = `-- name: SubmitWorkoutDetails :one
INSERT INTO tracker.workout_details (
  WORKOUT_NAME, GROUP_ID, EXERCISE_NAME, SETS, REPS, WEIGHT
) VALUES (
  $1, $2, $3, $4, $5, $6
)
ON CONFLICT (WORKOUT_NAME, GROUP_ID, EXERCISE_NAME) 
DO UPDATE SET 
  SETS = $4,
  REPS = $5,
  WEIGHT = $6,
  UPDT_TS = CURRENT_TIMESTAMP
RETURNING workout_name, group_id, exercise_name, sets, reps, weight, cret_ts, updt_ts
`

type SubmitWorkoutDetailsParams struct {
	WorkoutName  string        `json:"workoutName"`
	GroupID      int16         `json:"groupID"`
	ExerciseName string        `json:"exerciseName"`
	Sets         int16         `json:"sets"`
	Reps         int16         `json:"reps"`
	Weight       sql.NullInt16 `json:"weight"`
}

func (q *Queries) SubmitWorkoutDetails(ctx context.Context, arg SubmitWorkoutDetailsParams) (TrackerWorkoutDetail, error) {
	row := q.db.QueryRowContext(ctx, submitWorkoutDetails,
		arg.WorkoutName,
		arg.GroupID,
		arg.ExerciseName,
		arg.Sets,
		arg.Reps,
		arg.Weight,
	)
	var i TrackerWorkoutDetail
	err := row.Scan(
		&i.WorkoutName,
		&i.GroupID,
		&i.ExerciseName,
		&i.Sets,
		&i.Reps,
		&i.Weight,
		&i.CretTs,
		&i.UpdtTs,
	)
	return i, err
}

const submitWorkoutPerformed = `-- name: SubmitWorkoutPerformed :one
INSERT INTO tracker.workout_performed (
  SUBMITTED_ON, WORKOUT_NAME
) VALUES (
  $1, $2
)
RETURNING id, submitted_on, workout_name, cret_ts, updt_ts
`

type SubmitWorkoutPerformedParams struct {
	SubmittedOn time.Time `json:"submittedOn"`
	WorkoutName string    `json:"workoutName"`
}

func (q *Queries) SubmitWorkoutPerformed(ctx context.Context, arg SubmitWorkoutPerformedParams) (TrackerWorkoutPerformed, error) {
	row := q.db.QueryRowContext(ctx, submitWorkoutPerformed, arg.SubmittedOn, arg.WorkoutName)
	var i TrackerWorkoutPerformed
	err := row.Scan(
		&i.ID,
		&i.SubmittedOn,
		&i.WorkoutName,
		&i.CretTs,
		&i.UpdtTs,
	)
	return i, err
}
