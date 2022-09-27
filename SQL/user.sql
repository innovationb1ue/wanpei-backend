-- auto-generated definition
create table users
(
    ID         bigint auto_increment comment 'user id'
        primary key,
    games      varchar(512)  null comment 'tag of games',
    nickname   varchar(256)  null,
    username   varchar(256)  not null,
    password   varchar(256)  not null,
    email      varchar(256)  not null,
    created_at datetime      null,
    deleted_at datetime      null,
    updated_at datetime      null,
    user_role  int default 0 not null comment '0 - normal, 1- admin',
    avatar_url varchar(256)  null comment '头像url',
    steam_code varchar(64)   null comment 'steam好友代码'
);

