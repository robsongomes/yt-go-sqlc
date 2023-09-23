-- name: ListPosts :many
SELECT * FROM posts
ORDER BY title DESC;

-- name: CreatePost :exec
INSERT INTO posts (title, content, slug, author)
VALUES ($1, $2, $3, $4);

-- name: CreatePostAndReturnPost :one
INSERT INTO posts (title, content, slug, author)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdatePostAuthor :exec
UPDATE posts SET author = $1; --WARNING: NO WHERE

-- name: UpdatePostAuthorById :exec
UPDATE posts SET author = $1 WHERE id = $2;

-- name: DeletePostById :exec
DELETE FROM posts WHERE id = $1;