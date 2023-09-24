CREATE TABLE posts (
  id        BIGSERIAL PRIMARY KEY,
  title     TEXT      NOT NULL,
  content   TEXT      NOT NULL,
  slug      TEXT      NOT NULL,
  author    TEXT      
);

CREATE TABLE posts_views (
  post_id    BIGINT NOT NULL UNIQUE,
  views      INT NOT NULL DEFAULT 0,
  CONSTRAINT fk_posts FOREIGN KEY(post_id)
  REFERENCES posts(id)
);