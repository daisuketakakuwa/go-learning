CREATE TABLE IF NOT EXISTS products(
   product_id serial PRIMARY KEY,
   product_name VARCHAR (50),
   product_description VARCHAR(500),
   product_min_stock INTEGER
);