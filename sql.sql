create database if not exists adote_me;
use adote_me;

drop table if exists users;

create table users(
    id int auto_increment primary key,
    username varchar(50) not null,
    email varchar(50) not null unique,
    cellphone varchar(15) not null unique,
    passwd varchar(20) not null,
    picturePath varchar(200) null,
    createdAt timestamp default current_timestamp()
    -- default current_timestamp() -> por padrão o valor desse campo sempre será a data atual
);