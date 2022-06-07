CREATE TABLE users (
	user_id bigserial NOT NULL,
    name varchar(255) NOT NULL,
    hair_color varchar(255) NOT NULL,
    age int4 NOT NULL,
	CONSTRAINT user_id_pkey PRIMARY KEY (user_id)
);