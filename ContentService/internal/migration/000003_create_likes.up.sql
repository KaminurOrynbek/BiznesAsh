CREATE TABLE likes (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    post_id UUID,
    comment_id UUID,
    is_like BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    UNIQUE (post_id, user_id),
    UNIQUE (comment_id, user_id)

);

