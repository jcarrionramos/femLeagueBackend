PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "referees" (
    rut integer PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL UNIQUE,
    phone text NOT NULL UNIQUE,
    yellow_card integer NOT NULL,
    red_card integer NOT NULL
);
CREATE TABLE IF NOT EXISTS "teams" (
    name text PRIMARY KEY,
    win integer NOT NULL,
    draw integer NOT NULL,
    loss integer NOT NULL,
    total integer NOT NULL,
    active integer NOT NULL
);
INSERT INTO teams VALUES('Colo-Colo',0,0,0,0,0);
INSERT INTO teams VALUES('Universidad de Chile',0,0,0,0,0);
INSERT INTO teams VALUES('Cobreloa',0,0,0,0,0);
CREATE TABLE IF NOT EXISTS "players" (
    rut integer PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL UNIQUE,
    phone text NOT NULL UNIQUE,
    team_name text NOT NULL,
    dorsal_number integer NOT NULL,
    score integer NOT NULL
);
INSERT INTO players VALUES(1,'Javier','Carrion','javier@udp.cl','987373439','Colo-Colo',7,5);
INSERT INTO players VALUES(2,'Felipe','Delgado','felipe@udp.cl','2943078','Colo-Colo',11,3);
INSERT INTO players VALUES(3,'Jaime','Carbone','jaime@udp.cl','999999','Universidad de Chile',10,2);
CREATE TABLE IF NOT EXISTS "matches" (
    local_name text NOT NULL,
    visit_name text NOT NULL,
    season text NOT NULL,
    day integer NOT NULL,
    played integer NOT NULL,
    local_score integer NOT NULL,
    visit_score integer NOT NULL,
    referee_rut text NOT NULL,
    PRIMARY KEY (local_name, visit_name, season)
);
INSERT INTO matches VALUES('Colo-Colo', 'Universidad de Chile', '2-2018', 1, 0, 0, 0, '19137116-k');
INSERT INTO matches VALUES('Colo-Colo', 'Cobreloa', '2-2018', 2, 0, 0, 0, '19137116-k');
INSERT INTO matches VALUES('Cobreloa', 'Universidad de Chile', '2-2018', 3, 0, 0, 0, '19137116-k');
INSERT INTO matches VALUES('Universidad de Chile', 'Colo-Colo', '2-2018', 4, 0, 0, 0, '19137116-k');
COMMIT;
