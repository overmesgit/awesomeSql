BEGIN;

ALTER TABLE users DROP COLUMN mood;
DROP TYPE enum_mood;

COMMIT;