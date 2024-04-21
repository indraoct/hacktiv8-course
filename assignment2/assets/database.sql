
CREATE TABLE Orders (
                              order_id serial PRIMARY KEY,
                              customer_name VARCHAR (70) UNIQUE  NOT NULL,
                              ordered_at TIMESTAMPTZ,
                              created_at TIMESTAMPTZ,
                              updated_at TIMESTAMPTZ
);

CREATE TABLE Items (
                       item_id serial PRIMARY KEY,
                       item_code VARCHAR (35) UNIQUE  NOT NULL,
                       description VARCHAR (70)   NULL,
                       quantity INT NULL,
                       order_id INT REFERENCES Orders(order_id),
                       created_at TIMESTAMPTZ,
                       updated_at TIMESTAMPTZ
);