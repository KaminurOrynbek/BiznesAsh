CREATE TABLE IF NOT EXISTS notifications (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    message TEXT NOT NULL,
    post_id TEXT,
    comment_id TEXT,
    type TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    is_read BOOLEAN NOT NULL DEFAULT false
);
