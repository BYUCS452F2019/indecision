CREATE TABLE restaurant (
    restaurant_id INTEGER NOT NULL AUTO_INCREMENT,
    choice_id INTEGER NOT NULL,
    name varchar(255),
    category varchar(255),
    location varchar(255),
    price varchar(255),
    primary key (restaurant_id),
    foreign key (choice_id) REFERENCES choice(id)
);
