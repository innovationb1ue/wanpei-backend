create table if not exists wanpei.users
(
    Id         bigint auto_increment comment 'user id'
        primary key,
    games      varchar(512) null comment 'tag of games',
    nickname   varchar(256) null,
    username   varchar(256) not null,
    password   varchar(256) not null,
    email      varchar(256) not null,
    created_at datetime     null,
    deleted_at datetime     null,
    updated_at datetime     null
);

