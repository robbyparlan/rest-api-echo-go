-- public.enums definition

-- Drop table

-- DROP TABLE public.enums;

CREATE TABLE public.enums (
	id varchar(20) NOT NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	enum_name varchar(50) NOT NULL,
	enum_type_id varchar(10) NOT NULL,
	order_no int4 NULL
);

INSERT INTO public.enums (id,created_at,updated_at,enum_name,enum_type_id,order_no) VALUES
	 ('SCP01','2021-12-05 21:15:33.554835',NULL,'ADMIN','SCP',1),
	 ('SCP02','2021-12-05 21:17:28.474678',NULL,'GUEST','SCP',2);