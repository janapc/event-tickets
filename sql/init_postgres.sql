CREATE TYPE cur AS ENUM ('BRL', 'USD');

CREATE TABLE IF NOT EXISTS events(
  id uuid NOT NULL,
  name VARCHAR(70) NOT NULL,
  description VARCHAR(150) NOT NULL,
  image_url VARCHAR(250) NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  currency cur NOT NULL,
  event_date TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL, 
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(id)
)