INSERT INTO supplier ( name ) VALUES ( 'DrinkBitch' );
INSERT INTO supplier ( name ) VALUES ( 'DrinkBitch2' );
INSERT INTO drink ( ean, name, priceorder, priceresell, imgurl) VALUES ( '4260310553382', 'MATE MATE', 0.7, 1, '/static/img/matemate.jpg');
INSERT INTO perms (type ) VALUES ('user');
INSERT INTO perms ( type, patchdrinkEveryone, modsupplier, moddrink, moduser, setownpassword ) VALUES ( 'admin', true, true, true, true, true);
INSERT INTO customer ( username, firstname, lastname ) VALUES ( 'muster', 'Max', 'Mustermann' );
INSERT INTO customer ( username, firstname, lastname, perms ) VALUES ( 'mrpi', 'Markus', 'Mroch', 'admin' );
INSERT INTO customer ( username, firstname, lastname, perms, credit ) VALUES ( 'arsch', 'Kinder', 'Scheisse', 'user', -100);
