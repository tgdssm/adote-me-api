-- drop database adote_me;
create database if not exists adote_me;
use adote_me;

drop table if exists users;
-- drop table if exists profile_images;
drop table if exists pets;
drop table if exists pet_images;

create table IF not exists mysql.user (
    Host           VARCHAR(60),
    User           VARCHAR(32),
    Password       CHAR(41),
    Select_priv    ENUM('N','Y') DEFAULT 'N',
    Insert_priv    ENUM('N','Y') DEFAULT 'N',
    Update_priv    ENUM('N','Y') DEFAULT 'N',
    Delete_priv    ENUM('N','Y') DEFAULT 'N',
    Create_priv    ENUM('N','Y') DEFAULT 'N',
    Drop_priv      ENUM('N','Y') DEFAULT 'N',
    Reload_priv    ENUM('N','Y') DEFAULT 'N',
    Shutdown_priv  ENUM('N','Y') DEFAULT 'N',
    Process_priv   ENUM('N','Y') DEFAULT 'N',
    File_priv      ENUM('N','Y') DEFAULT 'N',
    Grant_priv     ENUM('N','Y') DEFAULT 'N',
    References_priv ENUM('N','Y') DEFAULT 'N',
    Index_priv     ENUM('N','Y') DEFAULT 'N',
    Alter_priv     ENUM('N','Y') DEFAULT 'N',
    Show_db_priv   ENUM('N','Y') DEFAULT 'N',
    Super_priv     ENUM('N','Y') DEFAULT 'N',
    Create_tmp_table_priv ENUM('N','Y') DEFAULT 'N',
    Lock_tables_priv ENUM('N','Y') DEFAULT 'N',
    Execute_priv   ENUM('N','Y') DEFAULT 'N',
    Repl_slave_priv ENUM('N','Y') DEFAULT 'N',
    Repl_client_priv ENUM('N','Y') DEFAULT 'N',
    Create_view_priv ENUM('N','Y') DEFAULT 'N',
    Show_view_priv  ENUM('N','Y') DEFAULT 'N',
    Create_routine_priv ENUM('N','Y') DEFAULT 'N',
    Alter_routine_priv ENUM('N','Y') DEFAULT 'N',
    Create_user_priv ENUM('N','Y') DEFAULT 'N',
    Event_priv     ENUM('N','Y') DEFAULT 'N',
    Trigger_priv   ENUM('N','Y') DEFAULT 'N',
    Create_tablespace_priv ENUM('N','Y') DEFAULT 'N',
    ssl_type       ENUM('','ANY','X509','SPECIFIED') DEFAULT '',
    ssl_cipher     BLOB,
    x509_issuer    BLOB,
    x509_subject   BLOB,
    max_questions  INT UNSIGNED DEFAULT 0,
    max_updates    INT UNSIGNED DEFAULT 0,
    max_connections INT UNSIGNED DEFAULT 0,
    max_user_connections INT UNSIGNED DEFAULT 0,
    plugin         CHAR(64) DEFAULT '',
    authentication_string TEXT,
    password_expired ENUM('N','Y') DEFAULT 'N',
    PRIMARY KEY (Host, User)
    ) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='Users and global privileges';

create table users(
    id int auto_increment primary key,
    username varchar(50) not null,
    email varchar(50) not null unique,
    cellphone varchar(15) not null unique,
    passwd varchar(100) not null,
    created_at timestamp default current_timestamp()
    -- default current_timestamp() -> por padrão o valor desse campo sempre será a data atual
);

-- create table profile_images(
   -- id int auto_increment primary key,
   -- file_name varchar(150) not null,
   -- file_path varchar(200) not null unique,
   -- user_id int unique,
   -- foreign key (user_id) 
   -- references users(id)
   -- on delete cascade
-- );

create table pets(
	id int auto_increment primary key,
    pet_name varchar(50) not null,
    age int,
    weight decimal(4,2),
    requirements varchar(300),
    created_at timestamp default current_timestamp(),
    user_id int,
    foreign key (user_id)
    references users(id)
    on delete cascade
);

create table pet_images(
	id int auto_increment primary key,
    file_name varchar(150) not null,
    file_path varchar(200) not null unique,
    pet_id int,
    foreign key (pet_id)
    references pets(id)
    on delete cascade
);