CREATE TYPE ANIMAL AS ENUM ('cat', 'dog', 'bird', 'fish', 'reptile');

CREATE TABLE Users (
	username VARCHAR,
	password VARCHAR,
	name VARCHAR,
	animal ANIMAL,
	PRIMARY KEY (username)
);

CREATE TABLE Follows (
	username VARCHAR,
	follows VARCHAR,
	PRIMARY KEY (username, follows),
	FOREIGN KEY (username) REFERENCES Users(username),
	FOREIGN KEY (follows) REFERENCES Users(username)
);

CREATE TABLE Posts (
	username VARCHAR,
	post_id SERIAL,
	post_text VARCHAR,
	post_date TIMESTAMP,
	PRIMARY KEY (post_id),
	FOREIGN KEY (username) REFERENCES Users(username)
);
