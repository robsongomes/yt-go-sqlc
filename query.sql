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

-- name: GetPostById :one
SELECT * FROM posts
WHERE id = $1;

-- name: GetPostsTitle :many
SELECT title FROM posts;

-- name: GetPostsAuthors :many
SELECT author FROM posts;

-- name: GetPostsInfo :many
SELECT title, author FROM posts;

-- name: GetPostsByIds :many
SELECT * FROM posts
WHERE id = ANY($1::int[]);

-- name: CountPosts :one
SELECT count(*) FROM posts;

-- name: CountPostsByAuthor :many
SELECT count(*), author from posts
WHERE author is not null
GROUP BY author;

-- name: SetPostViews :exec
-- INSERT INTO posts_views (post_id, views) VALUES ($1, 1)
-- ON CONFLICT(post_id) DO UPDATE SET views = EXCLUDED.views + 1;
UPDATE posts_views set views = views + 1 WHERE post_id = $1;

-- name: GetPostsViews :many
SELECT sqlc.embed(posts), sqlc.embed(posts_views) FROM posts
JOIN posts_views ON posts_views.post_id = posts.id;