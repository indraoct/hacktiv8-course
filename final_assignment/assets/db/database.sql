-- Create User table
CREATE TABLE  "user" (
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
CREATE TABLE  photo (
    id SERIAL PRIMARY KEY,
    caption TEXT,
    photo_url VARCHAR(255),
    user_id INTEGER REFERENCES "user"(id),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
    );

-- Create Comment table
CREATE TABLE comment (
     id SERIAL PRIMARY KEY,
     user_id INTEGER REFERENCES "user"(id),
     comment TEXT,
     photo_id INTEGER REFERENCES photo(id),
     message TEXT,
     created_at TIMESTAMPTZ,
     updated_at TIMESTAMPTZ
);

-- Create Social Media table
CREATE TABLE  socialmedia (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    social_media_url VARCHAR(255),
    user_id INTEGER REFERENCES "user"(id),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
    );
