-- public.api_clients definition

-- Drop table

-- DROP TABLE public.api_clients;

CREATE TABLE public.api_clients (
	id serial4 NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	client_id varchar(50) NOT NULL,
	secret_key varchar(100) NULL,
	grant_type varchar(20) NOT NULL DEFAULT 'credentials'::character varying,
	is_active bool NOT NULL DEFAULT false,
	scopes _text NULL,
	uuid varchar(10) NULL
);

INSERT INTO public.api_clients (created_at,updated_at,client_id,secret_key,grant_type,is_active,scopes,uuid) VALUES
	 ('2021-11-26 21:28:18.601532',NULL,'apiv1','5ae6ea9d886dfb01ca99b8aae3db70d','credentials',true,'{"SCP01","SCP02"}','123456789'),
	 ('2021-11-26 21:28:18.601532',NULL,'guest','5ae6ea9d886dfb01ca99b8aae3db70d','credentials',true,'{"SCP02",""}','123456789');