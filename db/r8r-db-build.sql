CREATE TABLE users (
	user_id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL,
	email VARCHAR ( 255 ) UNIQUE NOT NULL,
	created_on TIMESTAMP NOT NULL,
        last_login TIMESTAMP 
);

CREATE TABLE posts (
	post_id serial PRIMARY KEY,
	created TIMESTAMP NOT NULL,
	created_by VARCHAR(50) NOT NULL,
	type VARCHAR ( 50 ) NOT NULL,
	title VARCHAR ( 100 ) UNIQUE NOT NULL,
    description VARCHAR (500)
);

CREATE TABLE post_images (
    post_id integer NOT NULL,
    imageurl text NOT NULL,
    CONSTRAINT "FK_postimages_post" FOREIGN KEY (post_id)
        REFERENCES public.posts (post_id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID
);

CREATE TABLE post_ratings (
    post_id integer NOT NULL,
    username VARCHAR(50) NOT NULL,
    rating smallint NOT NULL,
    CONSTRAINT "FK_postratings_post" FOREIGN KEY (post_id)
        REFERENCES public.posts (post_id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT "FK_postratings_user" FOREIGN KEY (username)
        REFERENCES public.users (username) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID
);

CREATE OR REPLACE VIEW avg_post_ratings
 AS
SELECT p.*, avg(pr.rating) as "rating"
FROM posts AS p
	LEFT JOIN post_ratings AS pr ON  p.post_id = pr.post_id
GROUP BY p.post_id;

ALTER TABLE avg_post_ratings
    OWNER TO postgres;

INSERT INTO users (
    username,
    password,
    email,
    created_on
) VALUES (
    'admin',
    'admin',
    'admin@admin.com',
    NOW()
);