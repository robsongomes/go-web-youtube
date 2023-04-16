CREATE TABLE users (
    id int not null auto_increment,
    email varchar(255) unique,
    password varchar(255),
    primary key (id)
);

INSERT INTO users (email, password) values ('robson@gmail.com', '123456');

CREATE TABLE posts (
    id int not null auto_increment,
    title varchar(255) not null,
    slug varchar(255) not null unique,
    content text,
    user_id int not null,
    created_at timestamp default current_timestamp(),
    updated_at timestamp default current_timestamp(),
    primary key (id),
    foreign key (user_id) references users(id)
);