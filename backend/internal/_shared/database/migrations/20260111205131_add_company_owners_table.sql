-- Create "company_owners" table
CREATE TABLE "public"."company_owners" (
  "id" uuid NOT NULL,
  "email" character varying(255) NOT NULL,
  "phone" character varying(20) NOT NULL,
  "password" character varying(255) NOT NULL,
  "company_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id"),
  CONSTRAINT "company_owners_email_key" UNIQUE ("email"),
  CONSTRAINT "company_owners_company_id_fkey" FOREIGN KEY ("company_id") REFERENCES "public"."companies" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
