-- name: CreateUser :one
INSERT INTO users (
username,
hashed_password,
first_name,
last_name,
email,
contact_number,
address
) VALUES (
$1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;
