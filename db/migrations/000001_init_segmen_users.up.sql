CREATE TABLE users
(
    id         BIGINt PRIMARY KEY,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE segments
(
    slug TEXT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE relation_user_slugs
(
    id BIGINT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    user_id bigint references users,
    name_slug text references segments
)


CREATE TABLE user_segment_history
(
    user_id      INT,
    segment_slug TEXT,
    operation    TEXT,
    timestamp    TIMESTAMP
);