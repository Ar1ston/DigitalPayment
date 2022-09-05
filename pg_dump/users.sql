CREATE TABLE public.users (
    id SERIAL PRIMARY KEY ,
    "Login" CHARACTER VARYING(30) NOT NULL CHECK ( "Login"!='' ),
    "Password" TEXT NOT NULL CHECK ( "Password"!='' ),
    "Name" CHARACTER VARYING(30) DEFAULT '' NOT NULL,
    "Level" SMALLINT DEFAULT 1 NOT NULL CHECK ( "Level">0 AND "Level"<4 )
);
INSERT INTO public.users ("Login", "Password", "Name", "Level") VALUES ('Login1','PASSWORD','Name1',1);