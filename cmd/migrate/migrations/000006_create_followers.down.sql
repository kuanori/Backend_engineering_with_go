ALTER TABLE followers DROP CONSTRAINT user_id;
ALTER TABLE followers DROP CONSTRAINT follower_id;

DROP TABLE IF EXISTS followers;
