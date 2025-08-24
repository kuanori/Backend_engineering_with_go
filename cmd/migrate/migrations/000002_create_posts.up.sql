CREATE TABLE IF NOT EXISTS posts(
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    title text NOT NULL,
    content text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

ALTER TABLE posts
ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id);
