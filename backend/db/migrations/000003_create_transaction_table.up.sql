CREATE TABLE IF NOT EXISTS "transaction"(
  tranaction_id serial PRIMARY KEY,
  user_id VARCHAR (32),
  datetime timestamp,
  category_id integer,
  subcategory_id integer,
  debit integer,
  credit integer,

  CONSTRAINT fk_transactions_category
    FOREIGN KEY (category_id) 
    REFERENCES category (category_id)
    ON DELETE CASCADE,

  CONSTRAINT fk_transactions_subcategory
    FOREIGN KEY (subcategory_id) 
    REFERENCES subcategory (subcategory_id)
    ON DELETE CASCADE
);
