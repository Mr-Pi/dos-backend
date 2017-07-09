
--
-- TABLE: perms
-- 
--  

CREATE TABLE perms (
		type text NOT NULL ,
		patchDrinkEveryone bool NOT NULL  DEFAULT false,
		modSupplier bool NOT NULL  DEFAULT false,
		modDrink bool NOT NULL  DEFAULT false,
		modUser bool NOT NULL  DEFAULT false,
		setOwnPassword bool NOT NULL  DEFAULT true
);

-- 
ALTER TABLE perms ADD CONSTRAINT constraint_id PRIMARY KEY (type);

CREATE INDEX perms_type_index  ON perms(type);
ALTER TABLE perms ADD CONSTRAINT  FOREIGN KEY () REFERENCES customer ();
