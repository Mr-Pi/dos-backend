
--
-- TABLE: perms
-- 
--  

CREATE TABLE perms (
		id bigint NOT NULL ,
		type text,
		patchDrinkEveryone bool NOT NULL  DEFAULT false,
		modSupplier bool NOT NULL  DEFAULT false,
		modDrink bool NOT NULL  DEFAULT false,
		modUser bool NOT NULL  DEFAULT false,
		setOwnPassword bool NOT NULL  DEFAULT true
);
CREATE SEQUENCE perms_id_seq START 1 INCREMENT 1 ;
ALTER TABLE perms ALTER COLUMN id SET NOT 0;
ALTER TABLE perms ALTER COLUMN id SET DEFAULT nextval('perms_id_seq');

-- 
ALTER TABLE perms ADD CONSTRAINT constraint_id PRIMARY KEY (id);

CREATE INDEX perms_id_index  ON perms(id);

CREATE INDEX perms_type_index  ON perms(type);
ALTER TABLE perms ADD CONSTRAINT  FOREIGN KEY () REFERENCES customer ();
