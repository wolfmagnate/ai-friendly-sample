-- name: GetUserByUserUID :one
SELECT id, created_at, updated_at, user_uid FROM users
WHERE user_uid = @user_uid
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  user_uid, created_at, updated_at
) VALUES (
  @user_uid, extract(epoch from now())::bigint, extract(epoch from now())::bigint
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = @userID;