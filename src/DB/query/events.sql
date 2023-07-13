-- name: CreateEvent :one
INSERT INTO Events (
    c_id,
    start_time,
    end_time,
    title,
    color
) VALUES (
             $1, $2, $3, $4, $5
         )
RETURNING *;

-- name: GetEventsByCourseID :many
SELECT * FROM Events
WHERE c_id = $1
ORDER BY start_time;

-- name: GetEventByEventID :one
SELECT * FROM Events
WHERE e_id = $1
LIMIT 1;

-- name: UpdateEventTitle :one
UPDATE Events
SET title = $2
WHERE e_id = $1
RETURNING *;

-- name: UpdateEventColor :one
UPDATE Events
SET color = $2
WHERE e_id = $1
RETURNING *;

-- name: UpdateEventStartTime :one
UPDATE Events
SET start_time = $2
WHERE e_id = $1
RETURNING *;

-- name: UpdateEventEndTime :one
UPDATE Events
SET end_time = $2
WHERE e_id = $1
RETURNING *;

-- name: DeleteEventOfACourse :exec
DELETE FROM Events
WHERE e_id = $1;
