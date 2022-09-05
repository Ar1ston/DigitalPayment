
CREATE TABLE public.books (
    id SERIAL PRIMARY KEY ,
    "Author" BIGINT DEFAULT 1 NOT NULL CHECK ( "Author">0 ),
    "Publisher" BIGINT DEFAULT 1 NOT NULL CHECK ( "Publisher">0 ),
    "AddedUser" BIGINT DEFAULT 1 NOT NULL CHECK ("AddedUser">0),
    "AddedTime" timestamp without time zone,
    "Description" TEXT DEFAULT 'Нет описания' NOT NULL,
    "Name" TEXT NOT NULL CHECK ( "Name"!='' ),
    "Genre" TEXT NOT NULL CHECK ( "Genre"!='' )
);
INSERT INTO public.books ("Author", "Publisher", "AddedUser", "AddedTime", "Description", "Name", "Genre") VALUES (1,1,1,'2022-09-05 21:10:10.000000',DEFAULT,'Book','Genre');