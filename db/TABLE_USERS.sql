-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id serial4 NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	username varchar(50) NOT NULL,
	"password" varchar(100) NOT NULL,
	email varchar(100) NULL,
	fullname varchar(100) NULL,
	is_active bool NOT NULL DEFAULT false
);

