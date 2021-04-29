-- public.contacts definition

-- Drop table

-- DROP TABLE public.contacts;

CREATE TABLE public.contacts (
	contact_name varchar(100) NULL,
	contact_email varchar(100) NULL,
	contact_mobileno varchar(100) NULL,
	is_active int4 NULL,
	contact_address varchar(100) NULL,
	created_at timestamp(0) NULL,
	contactid int4 NOT NULL
);
CREATE INDEX contacts_contact_name_idx ON public.contacts USING btree (contact_name);