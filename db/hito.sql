CREATE TABLE IF NOT EXISTS users (
	id serial PRIMARY KEY,
	name VARCHAR (20) NOT NULL,
	username VARCHAR (50) UNIQUE NOT NULL,
	password VARCHAR (60) NOT NULL,
  admin BOOLEAN NOT NULL,
	last_login_at TIMESTAMP (6) NOT NULL,
	created_at TIMESTAMP (6) NOT NULL,
  updated_at TIMESTAMP (6) NOT NULL
);

CREATE TABLE IF NOT EXISTS landmarks (
  id serial PRIMARY KEY,
  name VARCHAR (255) UNIQUE NOT NULL,
  native_name VARCHAR (255) UNIQUE NOT NULL,
  type VARCHAR (255) NOT NULL,
  description TEXT NOT NULL,
  continent VARCHAR (255) NOT NULL,
  country VARCHAR (255) NOT NULL,
  city VARCHAR (255) NOT NULL,
  latitude NUMERIC NOT NULL,
  longitude NUMERIC NOT NULL,
  start_year SMALLINT NOT NULL,
  end_year SMALLINT NOT NULL,
  lengths NUMERIC NOT NULL,
  width NUMERIC NOT NULL,
	height NUMERIC NOT NULL,
  wiki_url TEXT NOT NULL,
	img_url TEXT NOT NULL,
	user_id INT NOT NULL,
	created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
	FOREIGN KEY (user_id)
		REFERENCES users (id)
);
