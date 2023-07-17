-- name: CreateUser :one
INSERT INTO users (firstname,
   lastname,
   username,
   email,
   password,
   verification_code
    )
VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET username = $2, verified = $3, verification_code = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE from users
WHERE id = $1;


