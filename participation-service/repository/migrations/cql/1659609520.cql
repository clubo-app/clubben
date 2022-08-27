CREATE TABLE IF NOT EXISTS party_participants (
    user_id text,
    party_id text,
    requested boolean,
    joined_at timestamp,
    requested_at timestamp,
    PRIMARY KEY ((party_id, requested), user_id)
) WITH cdc = {'enabled':true};

CREATE INDEX ON party_participants(user_id);

CREATE TABLE IF NOT EXISTS party_invites (
  user_id text,
  inviter_id text,
  party_id text,
  valid_until timestamp,
  PRIMARY KEY (user_id, inviter_id, party_id)
) WITH cdc = {'enabled':true};
