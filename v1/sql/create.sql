create database cmdshock
default character set utf8
default collate utf8_general_ci;

use cmdshock;

create table if not exists os(
id int auto_increment primary key,
platform varchar(100) not null,
architecture varchar(6) default 'x64'
)
default charset = utf8;

create table if not exists terminal(
id int auto_increment primary key, 
name varchar(100) not null,
os int not null,
constraint terminal_os_fk foreign key(os) references os(id)
)
default charset = utf8;

create table if not exists commands(
id int auto_increment primary key, 
name varchar(100) not null,
description varchar(255) default 'no description',
terminal int not null,
createdAt timestamp default current_timestamp(),
constraint commands_terminal_fk foreign key(terminal) references terminal(id)
)
default charset = utf8;
