
--
-- TABLE: supplier
-- 
--  

CREATE TABLE supplier (
		id unsigned int NOT NULL  DEFAULT 0,
		address string,
		name string,
		deliverTime unsigned int NOT NULL  DEFAULT 3
);
CREATE SEQUENCE supplier_id_seq START 1 INCREMENT 1 ;
ALTER TABLE supplier ALTER COLUMN id SET NOT 0;
ALTER TABLE supplier ALTER COLUMN id SET DEFAULT nextval('supplier_id_seq');

-- 
ALTER TABLE supplier ADD CONSTRAINT constraint_id PRIMARY KEY (id);

CREATE INDEX supplier_id_index  ON supplier(id);
ALTER TABLE supplier ADD CONSTRAINT  FOREIGN KEY (id) REFERENCES drink (supplier);
