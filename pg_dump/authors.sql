
CREATE TABLE public.authors(
    id serial PRIMARY KEY,
    "FirstName" CHARACTER VARYING(30) DEFAULT '' NOT NULL ,
    "LastName"  CHARACTER VARYING(30) DEFAULT '' NOT NULL ,
    "Description" TEXT DEFAULT 'Нет описания' NOT NULL
);

INSERT INTO public.authors ("FirstName", "LastName", "Description") VALUES ('Александр','Блок','Известный русский писатель эпохи просвещения');
INSERT INTO public.authors ("FirstName", "LastName", "Description") VALUES ('Александр','Пушкин','Русский поэт, драматург и прозаик, заложивший основы русского реалистического направления ');
INSERT INTO public.authors ("FirstName", "LastName", "Description") VALUES ('Лев','Толстой','один из наиболее известных русских писателей и мыслителей, один из величайших писателей-романистов мира');
