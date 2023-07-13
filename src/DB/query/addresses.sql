-- name: CreateAddresses :one
INSERT INTO Addresses (
    country, city, street
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetAddress :one
SELECT * from Addresses
WHERE country = $1 AND city = $2 AND street = $3
LIMIT 1;

-- name: GetAddressesByCountry :many
SELECT * from Addresses
WHERE country = $1
ORDER BY country;



