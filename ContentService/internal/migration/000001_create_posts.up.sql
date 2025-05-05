CREATE TABLE IF NOT EXISTS posts (
                                     id TEXT PRIMARY KEY,
                                     title TEXT NOT NULL,
                                     content TEXT NOT NULL,
                                     type TEXT NOT NULL,
                                     author_id TEXT NOT NULL,
                                     published BOOLEAN DEFAULT FALSE,
                                     likes_count INTEGER DEFAULT 0,
                                     dislikes_count INTEGER DEFAULT 0,
                                     comments_count INTEGER DEFAULT 0,
                                     created_at TIMESTAMP NOT NULL,
                                     updated_at TIMESTAMP NOT NULL
);
