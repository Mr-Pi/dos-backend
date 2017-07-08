
--
-- TABLE: drink
-- 
--  

CREATE TABLE drink (
		EAN string NOT NULL ,
		name string,
		amount unsigned int NOT NULL  DEFAULT 0,
		supplier unsigned int NOT NULL  DEFAULT 0,
		redeliverAmount unsigned int NOT NULL  DEFAULT 0,
		priceOrder float NOT NULL ,
		priceResell float NOT NULL 
);

-- 
ALTER TABLE drink ADD CONSTRAINT constraint_EAN PRIMARY KEY (EAN);

CREATE INDEX drink_EAN_index  ON drink(EAN);

CREATE INDEX drink_supplier_index  ON drink(supplier);
