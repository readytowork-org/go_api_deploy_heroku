CREATE TABLE IF NOT EXISTS users (
  id SERIAL,  
  email VARCHAR(100) NOT NULL,
  name VARCHAR(20) NOT NULL,
  age INT,
  birthday TIMESTAMPTZ,
  member_number VARCHAR(100),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  CONSTRAINT email_unique UNIQUE(email)
);