-- filepath: [20250110073909_add_salary_range_column_to_jobs.down.sql](http://_vscodecontentref_/0)
BEGIN;
ALTER TABLE jobs DROP COLUMN salary_max;
ALTER TABLE jobs
    RENAME COLUMN salary_min TO income_range;
COMMIT;
