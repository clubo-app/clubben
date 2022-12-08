CREATE TABLE profiles (
    id varchar(30) PRIMARY KEY, -- custom id from our lib. Prefix 2 char, ksuid 27, seperator 1
    username TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT,
    avatar TEXT
);

CREATE UNIQUE INDEX username_idx ON profiles (username);
