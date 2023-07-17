CREATE TABLE "users" (
     "id" bigserial PRIMARY KEY,
     "firstname" varchar,
     "lastname" varchar,
     "username" varchar UNIQUE NOT NULL,
     "email" varchar UNIQUE NOT NULL,
     "password" varchar NOT NULL,
     "verification_code" varchar,
     "verified" boolean DEFAULT false,
     "created_at" timestamptz NOT NULL DEFAULT (now())
);
