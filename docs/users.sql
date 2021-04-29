-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	userid int4 NOT NULL,
	full_name varchar(100) NULL,
	email varchar(100) NULL,
	mobile_no varchar(100) NULL,
	username varchar(100) NULL,
	passwd text NULL,
	created_at timestamp(0) NULL,
	CONSTRAINT user_pk PRIMARY KEY (userid)
);