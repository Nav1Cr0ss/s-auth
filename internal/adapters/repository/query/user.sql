-- name: CreateUser :one
INSERT INTO "user" (phone, email, hashed_password)
VALUES ($1, $2, $3)
RETURNING "id";

-- name: GetUser :one
SELECT *
FROM "user"
WHERE "id" = $1
LIMIT 1;

-- name: GetAccountByLogin :one
SELECT *
FROM "user"
WHERE phone = $1 or email = $1
LIMIT 1;