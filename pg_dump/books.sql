
CREATE TABLE public.books (
    id SERIAL PRIMARY KEY ,
    author BIGINT DEFAULT 1 NOT NULL CHECK ( author>0 ),
    publisher BIGINT DEFAULT 1 NOT NULL CHECK ( publisher>0 ),
    "addedUser" BIGINT DEFAULT 1 NOT NULL CHECK ("addedUser">0),
    "addedTime" timestamp without time zone,
    "description" TEXT DEFAULT 'Нет описания' NOT NULL,
    "name" TEXT NOT NULL CHECK ( "name"!='' ),
    "genre" TEXT NOT NULL CHECK ( genre!='' )
);