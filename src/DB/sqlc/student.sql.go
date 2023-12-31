// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: student.sql

package DB

import (
	"context"
	"database/sql"
)

const createStudent = `-- name: CreateStudent :one
INSERT INTO Student (user_name, hashed_password, full_name, email, u_id, credit)
VALUES ( $1, $2, $3, $4, $5, $6 )
RETURNING s_id, user_name, hashed_password, full_name, email, description, avatar, u_id, credit
`

type CreateStudentParams struct {
	UserName       string `json:"user_name"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	UID            int64  `json:"u_id"`
	Credit         int32  `json:"credit"`
}

func (q *Queries) CreateStudent(ctx context.Context, arg CreateStudentParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, createStudent,
		arg.UserName,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
		arg.UID,
		arg.Credit,
	)
	var i Student
	err := row.Scan(
		&i.SID,
		&i.UserName,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Description,
		&i.Avatar,
		&i.UID,
		&i.Credit,
	)
	return i, err
}

const deleteStudent = `-- name: DeleteStudent :one
DELETE FROM Student
WHERE s_id = $1
RETURNING s_id, user_name, hashed_password, full_name, email, description, avatar, u_id, credit
`

func (q *Queries) DeleteStudent(ctx context.Context, sID int64) (Student, error) {
	row := q.db.QueryRowContext(ctx, deleteStudent, sID)
	var i Student
	err := row.Scan(
		&i.SID,
		&i.UserName,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Description,
		&i.Avatar,
		&i.UID,
		&i.Credit,
	)
	return i, err
}

const getStudent = `-- name: GetStudent :one
SELECT s_id, user_name, hashed_password, full_name, email, description, avatar, u_id, credit from Student
WHERE s_id = $1
LIMIT 1
`

func (q *Queries) GetStudent(ctx context.Context, sID int64) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudent, sID)
	var i Student
	err := row.Scan(
		&i.SID,
		&i.UserName,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Description,
		&i.Avatar,
		&i.UID,
		&i.Credit,
	)
	return i, err
}

const getStudentByUserName = `-- name: GetStudentByUserName :one
SELECT s_id, user_name, hashed_password, full_name, email, description, avatar, u_id, credit from Student
WHERE user_name = $1
LIMIT 1
`

func (q *Queries) GetStudentByUserName(ctx context.Context, userName string) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudentByUserName, userName)
	var i Student
	err := row.Scan(
		&i.SID,
		&i.UserName,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Description,
		&i.Avatar,
		&i.UID,
		&i.Credit,
	)
	return i, err
}

const getStudentsByUniversity = `-- name: GetStudentsByUniversity :many
SELECT s_id, user_name, hashed_password, full_name, email, description, avatar, u_id, credit from Student
WHERE u_id = $1
`

func (q *Queries) GetStudentsByUniversity(ctx context.Context, uID int64) ([]Student, error) {
	rows, err := q.db.QueryContext(ctx, getStudentsByUniversity, uID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Student{}
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.SID,
			&i.UserName,
			&i.HashedPassword,
			&i.FullName,
			&i.Email,
			&i.Description,
			&i.Avatar,
			&i.UID,
			&i.Credit,
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

const updateStudentCredit = `-- name: UpdateStudentCredit :one
UPDATE Student
SET credit = $2
WHERE s_id = $1
RETURNING s_id, user_name, hashed_password, full_name, email, description, avatar, u_id, credit
`

type UpdateStudentCreditParams struct {
	SID    int64 `json:"s_id"`
	Credit int32 `json:"credit"`
}

func (q *Queries) UpdateStudentCredit(ctx context.Context, arg UpdateStudentCreditParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, updateStudentCredit, arg.SID, arg.Credit)
	var i Student
	err := row.Scan(
		&i.SID,
		&i.UserName,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Description,
		&i.Avatar,
		&i.UID,
		&i.Credit,
	)
	return i, err
}

const updateStudentPassword = `-- name: UpdateStudentPassword :one
UPDATE Student
SET hashed_password = $2
WHERE s_id = $1
RETURNING s_id, user_name, hashed_password, full_name, email, description, avatar, u_id, credit
`

type UpdateStudentPasswordParams struct {
	SID            int64  `json:"s_id"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) UpdateStudentPassword(ctx context.Context, arg UpdateStudentPasswordParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, updateStudentPassword, arg.SID, arg.HashedPassword)
	var i Student
	err := row.Scan(
		&i.SID,
		&i.UserName,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Description,
		&i.Avatar,
		&i.UID,
		&i.Credit,
	)
	return i, err
}

const updateStudentProfile = `-- name: UpdateStudentProfile :one
UPDATE Student
SET full_name = $2, description = $3, avatar = $4
WHERE s_id = $1
RETURNING s_id, user_name, hashed_password, full_name, email, description, avatar, u_id, credit
`

type UpdateStudentProfileParams struct {
	SID         int64          `json:"s_id"`
	FullName    string         `json:"full_name"`
	Description sql.NullString `json:"description"`
	Avatar      sql.NullString `json:"avatar"`
}

func (q *Queries) UpdateStudentProfile(ctx context.Context, arg UpdateStudentProfileParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, updateStudentProfile,
		arg.SID,
		arg.FullName,
		arg.Description,
		arg.Avatar,
	)
	var i Student
	err := row.Scan(
		&i.SID,
		&i.UserName,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.Description,
		&i.Avatar,
		&i.UID,
		&i.Credit,
	)
	return i, err
}
