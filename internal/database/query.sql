-- name: GetUserById :one
SELECT * FROM USERS
WHERE USER_ID = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM USERS;

-- name: CreateUser :one
INSERT INTO USERS (
    USER_NAME,
    USER_EMAIL,
    USER_PASSWORD
) VALUES (
$1, $2, $3
)
RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE USERS
  SET USER_PASSWORD = $2
WHERE USER_ID = $1;

-- name: DeleteUser :exec
DELETE FROM USERS
WHERE USER_ID = $1;

-- name: CreateJournalEntry :one
INSERT INTO
JOURNAL_ENTRIES (
    USER_ID,
    SYMPTOM_ID,
    EXERCISE_ID,
    JOURNALS_ID,
    JOURNAL_ENTRY_DATE
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: ListJournalEntries :many
SELECT * FROM JOURNAL_ENTRIES;

-- name: CreateJournal :one
INSERT INTO
JOURNALS (
    USER_ID,
    JOURNAL_NAME
) VALUES (
    $1, $2
) RETURNING *;

-- name: CreateExercise :one
INSERT INTO
EXERCISES (
    EXERCISE_NAME
) VALUES (
    $1
) RETURNING *;

-- name: ListExercises :many
SELECT * FROM EXERCISES;


