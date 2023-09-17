-- name: ListPosts :many
SELECT * FROM posts
ORDER BY title DESC;