CREATE TABLE movie (
    movie_id INTEGER NOT NULL AUTO_INCREMENT,
    choice_id INTEGER NOT NULL,
    title varchar(255),
    genre varchar(255),
    primary key (movie_id),
    foreign key (choice_id) REFERENCES choice(id)
);
