
--
-- TABLE: drink
-- 
--  

CREATE TABLE drink (
		EAN text NOT NULL ,
		name text,
		amount bigint NOT NULL  DEFAULT 0,
		supplier bigint NOT NULL  DEFAULT 0,
		redeliverAmount bigint NOT NULL  DEFAULT 0,
		priceOrder float NOT NULL ,
		priceResell float NOT NULL 
);

-- 
ALTER TABLE drink ADD CONSTRAINT constraint_EAN PRIMARY KEY (EAN);

CREATE INDEX drink_EAN_index  ON drink(EAN);

CREATE INDEX drink_supplier_index  ON drink(supplier);
