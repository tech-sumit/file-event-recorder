-- file_status to store file statuses
CREATE TABLE file_status
(
    path       TEXT PRIMARY KEY,
    last_event TEXT    NOT NULL,
    size       INTEGER NOT NULL
);
