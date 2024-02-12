CREATE TYPE ANIMAL AS ENUM ('cat', 'dog', 'bird', 'fish', 'reptile');

CREATE TABLE Users (
	username VARCHAR,
	password VARCHAR,
	name VARCHAR,
	animal ANIMAL,
	PRIMARY KEY (username)
);
