// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: mentor.sql

package DB

import (
	"context"
)

const createMentor = `-- name: CreateMentor :one
INSERT INTO Mentor (
    c_id,
    full_name,
    username,
    hashed_password,
    email,
    description,
    evaluation_count,
    score,
    balance,
    u_id,
    country,
    city,
    street
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
         )
RETURNING m_id, c_id, username, full_name, hashed_password, email, description, evaluation_count, score, balance, u_id, country, city, street
`

type CreateMentorParams struct {
	CID             int64  `json:"c_id"`
	FullName        string `json:"full_name"`
	Username        string `json:"username"`
	HashedPassword  string `json:"hashed_password"`
	Email           string `json:"email"`
	Description     string `json:"description"`
	EvaluationCount int32  `json:"evaluation_count"`
	Score           int32  `json:"score"`
	Balance         int32  `json:"balance"`
	UID             int64  `json:"u_id"`
	Country         string `json:"country"`
	City            string `json:"city"`
	Street          string `json:"street"`
}

func (q *Queries) CreateMentor(ctx context.Context, arg CreateMentorParams) (Mentor, error) {
	row := q.db.QueryRowContext(ctx, createMentor,
		arg.CID,
		arg.FullName,
		arg.Username,
		arg.HashedPassword,
		arg.Email,
		arg.Description,
		arg.EvaluationCount,
		arg.Score,
		arg.Balance,
		arg.UID,
		arg.Country,
		arg.City,
		arg.Street,
	)
	var i Mentor
	err := row.Scan(
		&i.MID,
		&i.CID,
		&i.Username,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.Description,
		&i.EvaluationCount,
		&i.Score,
		&i.Balance,
		&i.UID,
		&i.Country,
		&i.City,
		&i.Street,
	)
	return i, err
}

const getBalanceOfMentor = `-- name: GetBalanceOfMentor :one
SELECT balance from Mentor
WHERE m_id = $1
LIMIT 1
`

func (q *Queries) GetBalanceOfMentor(ctx context.Context, mID int64) (int32, error) {
	row := q.db.QueryRowContext(ctx, getBalanceOfMentor, mID)
	var balance int32
	err := row.Scan(&balance)
	return balance, err
}

const getMentor = `-- name: GetMentor :one
SELECT m_id, c_id, username, full_name, hashed_password, email, description, evaluation_count, score, balance, u_id, country, city, street from Mentor
WHERE m_id = $1
LIMIT 1
`

func (q *Queries) GetMentor(ctx context.Context, mID int64) (Mentor, error) {
	row := q.db.QueryRowContext(ctx, getMentor, mID)
	var i Mentor
	err := row.Scan(
		&i.MID,
		&i.CID,
		&i.Username,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.Description,
		&i.EvaluationCount,
		&i.Score,
		&i.Balance,
		&i.UID,
		&i.Country,
		&i.City,
		&i.Street,
	)
	return i, err
}

const getMentorByCourseID = `-- name: GetMentorByCourseID :one
SELECT m_id, c_id, username, full_name, hashed_password, email, description, evaluation_count, score, balance, u_id, country, city, street from Mentor
WHERE c_id = $1
LIMIT 1
`

func (q *Queries) GetMentorByCourseID(ctx context.Context, cID int64) (Mentor, error) {
	row := q.db.QueryRowContext(ctx, getMentorByCourseID, cID)
	var i Mentor
	err := row.Scan(
		&i.MID,
		&i.CID,
		&i.Username,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.Description,
		&i.EvaluationCount,
		&i.Score,
		&i.Balance,
		&i.UID,
		&i.Country,
		&i.City,
		&i.Street,
	)
	return i, err
}

const getMentorByUserName = `-- name: GetMentorByUserName :one
SELECT m_id, c_id, username, full_name, hashed_password, email, description, evaluation_count, score, balance, u_id, country, city, street from Mentor
WHERE username = $1
LIMIT 1
`

func (q *Queries) GetMentorByUserName(ctx context.Context, username string) (Mentor, error) {
	row := q.db.QueryRowContext(ctx, getMentorByUserName, username)
	var i Mentor
	err := row.Scan(
		&i.MID,
		&i.CID,
		&i.Username,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.Description,
		&i.EvaluationCount,
		&i.Score,
		&i.Balance,
		&i.UID,
		&i.Country,
		&i.City,
		&i.Street,
	)
	return i, err
}

const updateBalanceOfMentor = `-- name: UpdateBalanceOfMentor :one
UPDATE Mentor
SET balance = $2
WHERE m_id = $1
RETURNING m_id, c_id, username, full_name, hashed_password, email, description, evaluation_count, score, balance, u_id, country, city, street
`

type UpdateBalanceOfMentorParams struct {
	MID     int64 `json:"m_id"`
	Balance int32 `json:"balance"`
}

func (q *Queries) UpdateBalanceOfMentor(ctx context.Context, arg UpdateBalanceOfMentorParams) (Mentor, error) {
	row := q.db.QueryRowContext(ctx, updateBalanceOfMentor, arg.MID, arg.Balance)
	var i Mentor
	err := row.Scan(
		&i.MID,
		&i.CID,
		&i.Username,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.Description,
		&i.EvaluationCount,
		&i.Score,
		&i.Balance,
		&i.UID,
		&i.Country,
		&i.City,
		&i.Street,
	)
	return i, err
}

const updateScoreOfMentor = `-- name: UpdateScoreOfMentor :one
UPDATE Mentor
SET score = score + ($2), evaluation_count = evaluation_count+1
WHERE m_id = $1
RETURNING m_id, c_id, username, full_name, hashed_password, email, description, evaluation_count, score, balance, u_id, country, city, street
`

type UpdateScoreOfMentorParams struct {
	MID   int64 `json:"m_id"`
	Score int32 `json:"score"`
}

func (q *Queries) UpdateScoreOfMentor(ctx context.Context, arg UpdateScoreOfMentorParams) (Mentor, error) {
	row := q.db.QueryRowContext(ctx, updateScoreOfMentor, arg.MID, arg.Score)
	var i Mentor
	err := row.Scan(
		&i.MID,
		&i.CID,
		&i.Username,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.Description,
		&i.EvaluationCount,
		&i.Score,
		&i.Balance,
		&i.UID,
		&i.Country,
		&i.City,
		&i.Street,
	)
	return i, err
}
