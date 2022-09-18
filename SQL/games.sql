create table if not exists wanpei.games
(
    game_name        varchar(128) null,
    game_description varchar(512) null,
    ID               int          not null
        primary key
);

