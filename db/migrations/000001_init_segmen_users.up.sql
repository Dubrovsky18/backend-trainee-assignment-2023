CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP,
    segments   TEXT[]
);

CREATE TABLE segments
(
    slug TEXT PRIMARY KEY
);

CREATE TABLE user_segment_history
(
    user_id      INT,
    segment_slug TEXT,
    operation    TEXT,
    timestamp    TIMESTAMP
);