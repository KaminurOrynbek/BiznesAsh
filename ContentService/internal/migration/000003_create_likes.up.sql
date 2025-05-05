CREATE TABLE IF NOT EXISTS likes (
                                     id TEXT PRIMARY KEY,
                                     post_id TEXT NOT NULL,
                                     user_id TEXT NOT NULL,
                                     is_like BOOLEAN NOT NULL,
                                     created_at TIMESTAMP NOT NULL,
                                     FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
    );
