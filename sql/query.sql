-- name: AddFile :exec
INSERT INTO file_status (path, last_event, size)
VALUES (?, ?, ?);

-- name: GetFileByPath :one
SELECT *
FROM file_status
WHERE path = ? LIMIT 1;

-- name: ListAllFiles :many
SELECT *
FROM file_status
ORDER BY path;

-- name: UpdateFile :exec
UPDATE file_status
SET last_event = ?,
    size       = ?
WHERE path = ?;

-- name: DeleteFile :exec
DELETE
FROM file_status
WHERE path = ?;

-- name: UpsertFile :exec
INSERT INTO file_status (path, last_event, size)
VALUES (?, ?, ?) ON CONFLICT(path) DO
UPDATE SET last_event = excluded.last_event, size = excluded.size;
