-- name: FindUserByUsername :one
SELECT uuid, username, password
FROM idm_users
where username = $1;

-- name: Createuser
