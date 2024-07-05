CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    chat_id BIGINT NOT NULL,
    username VARCHAR(255),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    role_id INT NOT NULL,
    is_subscribed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE channels (
    id SERIAL PRIMARY KEY,
    chat_id BIGINT NOT NULL,
    channel_name VARCHAR(255) NOT NULL,
    is_admin BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    image_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE broadcasts (
    id SERIAL PRIMARY KEY,
    post_id INT NOT NULL,
    status_id INT NOT NULL
);
CREATE TABLE broadcast_recipients (
    id SERIAL PRIMARY KEY,
    broadcast_id INT NOT NULL,
    user_id INT NOT NULL,
    sent_at TIMESTAMP
);
CREATE TABLE broadcast_channels (
    id SERIAL PRIMARY KEY,
    post_id INT NOT NULL,
    channel_id INT NOT NULL,
    publish_time TIMESTAMP NOT NULL,
    expire_time TIMESTAMP NOT NULL
);
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);
CREATE TABLE broadcast_statuses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);
ALTER TABLE users
ADD FOREIGN KEY (role_id) REFERENCES roles(id);
ALTER TABLE broadcasts
ADD FOREIGN KEY (status_id) REFERENCES broadcast_statuses(id);
ALTER TABLE broadcast_recipients
ADD FOREIGN KEY (broadcast_id) REFERENCES broadcasts(id);
ALTER TABLE broadcast_recipients
ADD FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE broadcast_channels
ADD FOREIGN KEY (post_id) REFERENCES posts(id);
ALTER TABLE broadcast_channels
ADD FOREIGN KEY (channel_id) REFERENCES channels(id);
CREATE INDEX ON users (role_id);
CREATE INDEX ON broadcasts (status_id);
CREATE INDEX ON broadcast_recipients (broadcast_id);
CREATE INDEX ON broadcast_recipients (user_id);
CREATE INDEX ON broadcast_channels (post_id);
CREATE INDEX ON broadcast_channels (channel_id);