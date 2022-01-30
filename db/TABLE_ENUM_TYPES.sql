-- public.enum_types definition

-- Drop table

-- DROP TABLE public.enum_types;

CREATE TABLE public.enum_types (
	id varchar(10) NOT NULL,
	type_name varchar(100) NULL
);

INSERT INTO public.enum_types (id,type_name) VALUES
	 ('SCP','Scope');