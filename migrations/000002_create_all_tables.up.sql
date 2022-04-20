CREATE TABLE IF NOT EXISTS countries (
    code int PRIMARY KEY,
    name varchar,
    continent_name varchar
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    full_name varchar,
    phone_number varchar,
    username varchar,
    country_code int,
    created_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    user_id varchar,
    title varchar,
    time varchar,
    place varchar,
    latitude float8,
    longitude float8,
    created_at timestamp
);

CREATE TABLE IF NOT EXISTS participants (
    id SERIAL PRIMARY KEY,
    event_id int,
    user_id varchar
);

CREATE TABLE IF NOT EXISTS notifications (
    id SERIAL PRIMARY KEY,
    event_id int,
    created_at timestamp
);

ALTER TABLE users ADD FOREIGN KEY (country_code) REFERENCES countries (code);

ALTER TABLE events ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE participants ADD FOREIGN KEY (event_id) REFERENCES events (id);

ALTER TABLE participants ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE notifications ADD FOREIGN KEY (event_id) REFERENCES events (id);
