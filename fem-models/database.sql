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
    address text NOT NULL,
    phone text NOT NULL,
    win integer NOT NULL,
    draw integer NOT NULL,
    loss integer NOT NULL
);
CREATE TABLE IF NOT EXISTS "players" (
    rut integer PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL UNIQUE,
    phone text NOT NULL UNIQUE,
    team_name NOT NULL,
    dorsal_number text NOT NULL,
    goal number NOT NULL
);
CREATE TABLE IF NOT EXISTS "matchs" (
    local_name text NOT NULL,
    visit_name text NOT NULL,
    local_goal integer NOT NULL,
    visit_goal integer NOT NULL,
    referee_rut text NOT NULL,
    active integer NOT NULL,
    day integer NOT NULL,
    PRIMARY KEY (local_name, visit_name)
);
COMMIT;
