
CREATE TABLE public.authors(
    id serial PRIMARY KEY,
    firstName CHARACTER VARYING(30) DEFAULT '' NOT NULL ,
    lastName  CHARACTER VARYING(30) DEFAULT '' NOT NULL ,
    description TEXT
);

INSERT INTO public.authors (firstName, lastName, description) VALUES ('Александр','Блок','Известный русский писатель эпохи просвещения');
INSERT INTO public.authors (firstName, lastName, description) VALUES ('Александр','Пушкин','Русский поэт, драматург и прозаик, заложивший основы русского реалистического направления ');
INSERT INTO public.authors (firstName, lastName, description) VALUES ('Лев','Толстой','один из наиболее известных русских писателей и мыслителей, один из величайших писателей-романистов мира');
