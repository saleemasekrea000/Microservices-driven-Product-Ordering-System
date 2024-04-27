
CREATE TABLE IF NOT EXISTS order_items (
  id SERIAL PRIMARY KEY,
  product_id int,
  order_id int,
  amount int,
  price float
);

CREATE INDEX ON order_items (order_id);

ALTER TABLE order_items ADD CONSTRAINT order_id_forg FOREIGN KEY (order_id) REFERENCES orders (id);