CREATE TABLE IF NOT EXISTS subcategory(
  subcategory_id serial PRIMARY KEY,
  "name" VARCHAR (64),
  user_id VARCHAR (32)
);