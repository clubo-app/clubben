CREATE TABLE parties (
    id varchar(27) PRIMARY KEY,
    user_id TEXT NOT NULL,
    title TEXT NOT NULL,
    description VARCHAR(150),
    is_public BOOLEAN NOT NULL DEFAULT false,
    max_participants INTEGER NOT NULL DEFAULT 0,
    location geometry(POINT, 4326) NOT NULL,
    street_address TEXT,
    postal_code TEXT,
    state TEXT,
    country TEXT,
    entry_date TIMESTAMP NOT NULL
);

CREATE INDEX parties_by_user_id_idx ON parties (user_id, is_public, id);

CREATE INDEX party_location_idx
ON parties
USING GIST( (location::geography) );
