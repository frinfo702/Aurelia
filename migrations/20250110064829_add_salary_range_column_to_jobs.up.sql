BEGIN;
ALTER TABLE jobs
    RENAME COLUMN income_range TO salary_min;
ALTER TABLE jobs
ADD COLUMN salary_max INT;
COMMIT;
