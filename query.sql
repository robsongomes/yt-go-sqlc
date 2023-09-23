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