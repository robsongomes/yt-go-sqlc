CREATE TABLE posts (
  id        BIGSERIAL PRIMARY KEY,
  title     TEXT      NOT NULL,
  content   TEXT      NOT NULL,
  slug      TEXT      NOT NULL,
  author    TEXT      
);