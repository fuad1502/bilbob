CREATE TYPE ANIMAL AS ENUM ('cat', 'dog', 'bird', 'fish', 'reptile');
CREATE TYPE FOLLOWS_STATE AS ENUM ('requested', 'follows');

CREATE TABLE Users (
	username VARCHAR NOT NULL,
	password VARCHAR NOT NULL,
	name VARCHAR NOT NULL,
	animal ANIMAL NOT NULL,
	profile_pic_path VARCHAR,
	PRIMARY KEY (username)
);

CREATE TABLE Followings (
	username VARCHAR NOT NULL,
	follows VARCHAR NOT NULL,
	state FOLLOWS_STATE NOT NULL,
	PRIMARY KEY (username, follows),
	FOREIGN KEY (username) REFERENCES Users(username),
	FOREIGN KEY (follows) REFERENCES Users(username)
);

CREATE TABLE Posts (
	username VARCHAR NOT NULL,
	post_id SERIAL NOT NULL,
	post_text VARCHAR NOT NULL,
	post_date TIMESTAMP NOT NULL,
	PRIMARY KEY (post_id),
	FOREIGN KEY (username) REFERENCES Users(username)
);
