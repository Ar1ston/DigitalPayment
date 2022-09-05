CREATE TABLE public.users (
    id SERIAL PRIMARY KEY ,
    "login" CHARACTER VARYING(30) NOT NULL CHECK ( "login"!='' ),
    "password" TEXT NOT NULL CHECK ( "password"!='' ),
    "name" CHARACTER VARYING(30) DEFAULT '' NOT NULL,
    "level" SMALLINT DEFAULT 1 NOT NULL CHECK ( "level">0 AND "level"<4 )
);
