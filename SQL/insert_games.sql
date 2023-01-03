-- temporal solution for games.
truncate table wanpei.games;
-- replace index count
alter table wanpei.games auto_increment=1;
INSERT INTO wanpei.games (game_name, game_description) VALUES ('Dota2', 'Valve 5v5');
INSERT INTO wanpei.games (game_name, game_description) VALUES ('LOL', 'Riot 5V5');
INSERT INTO wanpei.games (game_name, game_description) VALUES ('CSGO', 'FPS');
INSERT INTO wanpei.games (game_name, game_description) VALUES ('Valorant', 'RIOT FPS');
INSERT INTO wanpei.games (game_name, game_description) VALUES ('Heart Of Iron 4', 'Be a war criminal');
INSERT INTO wanpei.games (game_name, game_description) VALUES ('BattleField 2042', 'nuts shooting each other');
INSERT INTO wanpei.games (game_name, game_description) VALUES ('DeadByDaylight', 'the farmer and 4 thieves');
INSERT INTO wanpei.games (game_name, game_description) VALUES ('Apex Legends', 'aim!');
INSERT INTO wanpei.games (game_name, game_description) VALUES ('Destiny 2', 'My liver is dead');
INSERT INTO wanpei.games (game_name, game_description) VALUES ('OverWatch', 'Gonna be next boom');
INSERT INTO wanpei.games (game_name, game_description) VALUES ('Age of Empires IV', 'AOE4');
INSERT INTO wanpei.games (game_name, game_description) VALUES ('Age of Empires II', 'AOE2');
