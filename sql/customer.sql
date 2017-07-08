
--
-- TABLE: customer
-- 
--  

CREATE TABLE customer (
		username text NOT NULL ,
		firstName text NOT NULL ,
		credit float NOT NULL  DEFAULT 0,
		password text,
		salt text,
		perms bigint NOT NULL  DEFAULT 1,
		lastName text NOT NULL 
);

-- 
ALTER TABLE customer ADD CONSTRAINT constraint_username PRIMARY KEY (username);

CREATE INDEX customer_username_index  ON customer(username);

CREATE INDEX customer_perms_index  ON customer(perms);
ALTER TABLE customer ADD CONSTRAINT  FOREIGN KEY (perms) REFERENCES perms (id);
