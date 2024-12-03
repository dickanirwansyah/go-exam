create database db_go_exam;

create table accounts (
    id bigserial primary key,
    email varchar(150) not null,
    full_name varchar(200) not null,
    phone_number varchar(30) not null,
    roles_id bigint not null,
    roles_name varchar(120) not null,
    image_profile varchar(255),
    address_detail varchar(255),
    created_at date,
    updated_at date, 
    is_deleted int not null
);

create table roles (
    id bigserial primary key,
    name varchar(120) not null,
    created_at date,
    updated_at date,
    is_deleted int not null
);

create table permissions(
    id bigserial primary key,
    endpoint varchar(200) not null,
    name varchar(100) not null,
    parent_id bigint,
    level int not null,
    icon varchar(100),
    is_deleted int not null
);

create table permissions_roles(
    id bigserial primary key,
    permissions_id bigint not null,
    roles_id bigint not null,
    is_deleted int not null
);

create table questions_category(
    id bigserial primary key,
    name varchar(150) not null,
    is_deleted int not null
);

create table questions (
    id bigserial primary key,
    text TEXT not null,
    questions_category_id bigint not null,
    is_deleted int not null,
    created_at date not null,
    updated_at date not null
);

create table answer (
    id bigserial primary key,
    questions_id bigint not null,
    text TEXT not null,
    is_correct boolean default false,
    created_at date not null,
    updated_at date not null
);

