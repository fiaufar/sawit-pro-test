/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

/** This is test table. Remove this table and replace with your own tables. */
CREATE TABLE "user" (
	id serial PRIMARY KEY,
	fullname VARCHAR ( 60 ) NOT NULL,
	phone_number VARCHAR ( 14 ) UNIQUE NOT NULL,
  	successful_logged_in int NOT NULL DEFAULT 0
);

CREATE TABLE user_credential (
	user_id serial PRIMARY KEY,
  	salt VARCHAR(50) NOT NULL,
	"password" TEXT NOT NULL
);