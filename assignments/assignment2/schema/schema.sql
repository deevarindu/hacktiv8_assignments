CREATE TABLE IF NOT EXISTS orders (
  order_id VARCHAR(10) PRIMARY KEY,
  customer_name VARCHAR(255) NOT NULL,
  ordered_at DATE NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS items (
  item_id VARCHAR(10) PRIMARY KEY,
	item_code VARCHAR(255) NOT NULL,
	description VARCHAR(255),
	quantity INT NOT NULL,
  order_id VARCHAR(10) NOT NULL,
  FOREIGN KEY (order_id) REFERENCES orders(order_id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
