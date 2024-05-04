-- Create User table
CREATE TABLE  "users" (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    title VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    age INTEGER,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
    );

-- Create Photo table
CREATE TABLE  photos (
    id SERIAL PRIMARY KEY,
    caption TEXT,
    photo_url VARCHAR(255),
    user_id INTEGER REFERENCES "users"(id),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
    );

-- Create Comment table
CREATE TABLE comments (
     id SERIAL PRIMARY KEY,
     user_id INTEGER REFERENCES "users"(id),
     comment TEXT,
     photo_id INTEGER REFERENCES photos(id),
     message TEXT,
     created_at TIMESTAMPTZ,
     updated_at TIMESTAMPTZ
);

-- Create Social Media table
CREATE TABLE  socialmedias (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    social_media_url VARCHAR(255),
    user_id INTEGER REFERENCES "users"(id),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
    );



-- Insert into User table
INSERT INTO "users" (username, title, email, password, age, created_at, updated_at) VALUES
                                                                                       ('user1', 'Title One', 'user1@example.com', 'password1', 25, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                                       ('user2', 'Title Two', 'user2@example.com', 'password2', 30, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                                       ('user3', 'Title Three', 'user3@example.com', 'password3', 35, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                                       ('user4', 'Title Four', 'user4@example.com', 'password4', 40, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                                       ('user5', 'Title Five', 'user5@example.com', 'password5', 45, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert into Photo table
INSERT INTO photos (caption, photo_url, user_id, created_at, updated_at) VALUES
                                                                            ('Caption 1', 'photo_url_1.jpg', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                            ('Caption 2', 'photo_url_2.jpg', 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                            ('Caption 3', 'photo_url_3.jpg', 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                            ('Caption 4', 'photo_url_4.jpg', 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                            ('Caption 5', 'photo_url_5.jpg', 5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert into Comment table
INSERT INTO comments (user_id, comment, photo_id, message, created_at, updated_at) VALUES
                                                                                      (1, 'Comment 1', 1, 'Message 1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                                      (2, 'Comment 2', 2, 'Message 2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                                      (3, 'Comment 3', 3, 'Message 3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                                      (4, 'Comment 4', 4, 'Message 4', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                                      (5, 'Comment 5', 5, 'Message 5', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert into Social Media table
INSERT INTO socialmedias (name, social_media_url, user_id, created_at, updated_at) VALUES
                                                                                      ('User One', 'https://twitter.com/user1', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                                      ('User Two', 'https://facebook.com/user2', 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                                      ('User Three', 'https://instagram.com/user3', 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                                      ('User Four', 'https://twitter.com/user4', 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
                                                                                      ('User Five', 'https://facebook.com/user5', 5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

