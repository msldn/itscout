CREATE DATABASE ITSCOUT;
\c itscout

CREATE TABLE cis (
	id serial PRIMARY KEY,
	type VARCHAR ( 50 ) NOT NULL,
	name VARCHAR ( 50 ) NOT NULL,
	created_on TIMESTAMP NOT NULL
);


INSERT INTO cis(id, type, name, created_on ) VALUES (1, 'SERVER', 'FIRSTDSERVER', NOW());
INSERT INTO cis(id, type, name, created_on ) VALUES (2, 'SERVER', 'SECONDSERVER', NOW());
INSERT INTO cis(id, type, name, created_on ) VALUES (3, 'SERVER', 'THIRDSERVER', NOW());
INSERT INTO cis(id, type, name, created_on ) VALUES (4, 'SERVER', 'FOURTHSERVER', NOW());
INSERT INTO cis(id, type, name, created_on ) VALUES (5, 'APPLICATOIN', 'FIRSTAPP', NOW());
INSERT INTO cis(id, type, name, created_on ) VALUES (6, 'APPLICATOIN', 'SECONDAPP', NOW());
INSERT INTO cis(id, type, name, created_on ) VALUES (7, 'APPLICATOIN', 'THIRDAPP', NOW());
INSERT INTO cis(id, type, name, created_on ) VALUES (8, 'APPLICATOIN', 'FOURTHAPP', NOW());



create or replace procedure getcis(
)
language plpgsql    
as $$
begin
    select * From cis
    commit;
end;$$
