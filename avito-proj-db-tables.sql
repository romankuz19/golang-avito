CREATE TABLE users
(
    id            serial       primary key,
    name          varchar(255) not null,
);

CREATE TABLE sections
(
    id          serial       primary key,
    name       varchar(255) not null,
);

CREATE TABLE users_sections
(
    id      serial                                           not null unique,
    user_id int not null ,
    section_id int not null ,
    FOREIGN KEY (user_id) REFERENCES users ON DELETE CASCADE,
    FOREIGN KEY (section_id) REFERENCES sections ON DELETE CASCADE
);
