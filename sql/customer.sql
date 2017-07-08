
--
-- TABLE: customer
-- 
--  

CREATE TABLE customer (
		username string NOT NULL ,
		fullName string,
		credit float NOT NULL  DEFAULT 0,
		password string,
		salt string NOT NULL ,
		perms unsigned int NOT NULL  DEFAULT 0,
		tag string
);

-- 
ALTER TABLE customer ADD CONSTRAINT constraint_username PRIMARY KEY (username);

CREATE INDEX customer_username_index  ON customer(username);

CREATE INDEX customer_perms_index  ON customer(perms);
ALTER TABLE customer ADD CONSTRAINT  FOREIGN KEY (perms) REFERENCES perms (id);
