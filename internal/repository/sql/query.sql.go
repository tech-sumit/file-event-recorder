// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package sql

import (
	"context"
)

const addFile = `-- name: AddFile :exec
INSERT INTO file_status (path, last_event, size)
VALUES (?, ?, ?)
`

type AddFileParams struct {
	Path      string
	LastEvent string
	Size      int64
}

func (q *Queries) AddFile(ctx context.Context, arg AddFileParams) error {
	_, err := q.db.ExecContext(ctx, addFile, arg.Path, arg.LastEvent, arg.Size)
	return err
}

const deleteFile = `-- name: DeleteFile :exec
DELETE
FROM file_status
WHERE path = ?
`

func (q *Queries) DeleteFile(ctx context.Context, path string) error {
	_, err := q.db.ExecContext(ctx, deleteFile, path)
	return err
}

const getFileByPath = `-- name: GetFileByPath :one
SELECT path, last_event, size
FROM file_status
WHERE path = ? LIMIT 1
`

func (q *Queries) GetFileByPath(ctx context.Context, path string) (FileStatus, error) {
	row := q.db.QueryRowContext(ctx, getFileByPath, path)
	var i FileStatus
	err := row.Scan(&i.Path, &i.LastEvent, &i.Size)
	return i, err
}

const listAllFiles = `-- name: ListAllFiles :many
SELECT path, last_event, size
FROM file_status
ORDER BY path
`

func (q *Queries) ListAllFiles(ctx context.Context) ([]FileStatus, error) {
	rows, err := q.db.QueryContext(ctx, listAllFiles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FileStatus
	for rows.Next() {
		var i FileStatus
		if err := rows.Scan(&i.Path, &i.LastEvent, &i.Size); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFile = `-- name: UpdateFile :exec
UPDATE file_status
SET last_event = ?,
    size       = ?
WHERE path = ?
`

type UpdateFileParams struct {
	LastEvent string
	Size      int64
	Path      string
}

func (q *Queries) UpdateFile(ctx context.Context, arg UpdateFileParams) error {
	_, err := q.db.ExecContext(ctx, updateFile, arg.LastEvent, arg.Size, arg.Path)
	return err
}

const upsertFile = `-- name: UpsertFile :exec
INSERT INTO file_status (path, last_event, size)
VALUES (?, ?, ?) ON CONFLICT(path) DO
UPDATE SET last_event = excluded.last_event, size = excluded.size
`

type UpsertFileParams struct {
	Path      string
	LastEvent string
	Size      int64
}

func (q *Queries) UpsertFile(ctx context.Context, arg UpsertFileParams) error {
	_, err := q.db.ExecContext(ctx, upsertFile, arg.Path, arg.LastEvent, arg.Size)
	return err
}
