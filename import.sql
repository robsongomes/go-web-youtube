CREATE TABLE users (
    id int not null auto_increment,
    email varchar(255) unique,
    password varchar(255),
    primary key (id)
);

INSERT INTO users (email, password) values ('robson@gmail.com', '123456');