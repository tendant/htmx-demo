-- name: FindUserByUsername :one
SELECT uuid, username, password
FROM idm_users
where username = $1;

-- name: FindTraining :many
SELECT uuid, name, created_at
FROM training
where name like '%demo%';
