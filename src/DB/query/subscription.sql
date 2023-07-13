-- name: CreateSubscription :one
INSERT INTO Subscription (
    s_id,
    plan_type,
    sub_expire_time,
    sub_start_time
) VALUES (
             $1, $2, $3, $4
         )
RETURNING *;

-- name: GetSubscriptionBySubID :one
SELECT * FROM Subscription
WHERE sub_id = $1
LIMIT 1;

-- name: GetSubscriptionBySID :one
SELECT * FROM Subscription
WHERE s_id = $1
LIMIT 1;

-- name: UpdateSubscriptionType :one
UPDATE Subscription
SET plan_type = $2
WHERE s_id = $1
RETURNING *;

-- name: UpdateSubscriptionTime :one
UPDATE Subscription
SET sub_expire_time = $2
WHERE s_id = $1
RETURNING *;

-- name: DeleteSubscriptionOfAStudent :exec
DELETE FROM Subscription
WHERE s_id = $1;
