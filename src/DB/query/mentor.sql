-- name: CreateMentor :one
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
RETURNING *;

-- name: GetMentor :one
SELECT * from Mentor
WHERE m_id = $1
LIMIT 1;

-- name: GetMentorByUserName :one
SELECT * from Mentor
WHERE username = $1
LIMIT 1;

-- name: GetMentorByCourseID :one
SELECT * from Mentor
WHERE c_id = $1
LIMIT 1;

-- name: GetBalanceOfMentor :one
SELECT balance from Mentor
WHERE m_id = $1
LIMIT 1;

-- name: UpdateBalanceOfMentor :one
UPDATE Mentor
SET balance = $2
WHERE m_id = $1
RETURNING *;

-- name: UpdateScoreOfMentor :one
UPDATE Mentor
SET score = score + ($2), evaluation_count = evaluation_count+1
WHERE m_id = $1
RETURNING *;