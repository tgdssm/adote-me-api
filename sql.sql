-- drop database adote_me;
create database if not exists adote_me;
use adote_me;

drop table if exists users;
drop table if exists profile_images;
drop table if exists pets;
drop table if exists pet_images;

create table users(
    id int auto_increment primary key,
    username varchar(50) not null,
    email varchar(50) not null unique,
    cellphone varchar(15) not null unique,
    passwd varchar(100) not null,
    created_at timestamp default current_timestamp()
    -- default current_timestamp() -> por padrão o valor desse campo sempre será a data atual
);

create table profile_images(
    id int auto_increment primary key,
    file_name varchar(150) not null,
    file_path varchar(200) not null unique,
    user_id int unique,
    foreign key (user_id) 
    references users(id)
    on delete cascade
);

create table pets(
	id int auto_increment primary key,
    pet_name varchar(50) not null,
    age int,
    weight decimal(4,2),
    requirements varchar(300),
    user_id int,
    foreign key (user_id)
    references users(id)
    on delete cascade
);

create table pet_images(
	id int auto_increment primary key,
    file_name varchar(150) not null,
    file_path varchar(200) not null unique,
    user_id int,
    pet_id int,
    foreign key (user_id)
    references users(id)
    on delete cascade,
    foreign key (pet_id)
    references pets(id)
    on delete cascade
);