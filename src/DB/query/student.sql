-- name: CreateStudent :one
INSERT INTO Student (user_name, hashed_password, full_name, email, u_id, credit)
VALUES ( $1, $2, $3, $4, $5, $6 )
RETURNING *;

-- name: GetStudent :one
SELECT * from Student
WHERE s_id = $1
LIMIT 1;

-- name: GetStudentByUserName :one
SELECT * from Student
WHERE user_name = $1
LIMIT 1;

-- name: GetStudentsByUniversity :many
SELECT * from Student
WHERE u_id = $1
;

-- name: UpdateStudentCredit :one
UPDATE Student
SET credit = $2
WHERE s_id = $1
RETURNING *;

-- name: UpdateStudentProfile :one
UPDATE Student
SET full_name = $2, description = $3, avatar = $4
WHERE s_id = $1
RETURNING *;

-- name: UpdateStudentPassword :one
UPDATE Student
SET hashed_password = $2
WHERE s_id = $1
RETURNING *;

-- name: DeleteStudent :one
DELETE FROM Student
WHERE s_id = $1
RETURNING *;
