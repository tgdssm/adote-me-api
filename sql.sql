-- drop database adote_me;
create database if not exists adote_me;
use adote_me;

drop table if exists users;
drop table if exists profile_images;

create table users(
    id int auto_increment primary key,
    username varchar(50) not null,
    email varchar(50) not null unique,
    cellphone varchar(15) not null unique,
    passwd varchar(20) not null,
    created_at timestamp default current_timestamp()
    -- default current_timestamp() -> por padrão o valor desse campo sempre será a data atual
);

create table profile_images(
    id int auto_increment primary key,
    file_name varchar(150) not null,
    file_path varchar(200) not null unique,
    user_id int,
    constraint fk_user_id foreign key (user_id)
    references users(id)
    on delete cascade
);