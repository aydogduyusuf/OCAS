-- name: CreateUniversity :one
INSERT INTO University(
    university_name,
    abbreviation,
    email_extension,
    country,
    city,
    street
)
VALUES ( $1, $2, $3, $4, $5, $6 )
RETURNING *;

-- name: GetUniversity :one
SELECT * FROM University
WHERE u_id = $1
LIMIT 1;

-- name: GetUniversityByName :one
SELECT * FROM University
WHERE university_name = $1
LIMIT 1;

-- name: GetAllUniversities :many
SELECT * FROM University;
