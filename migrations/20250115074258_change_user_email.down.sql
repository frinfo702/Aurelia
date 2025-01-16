BEGIN;
-- Remove trigger and function
DROP TRIGGER IF EXISTS update_users_updated_at ON users;
DROP FUNCTION IF EXISTS update_updated_at_column();
-- Remove timestamp columns
ALTER TABLE users DROP COLUMN created_at,
    DROP COLUMN updated_at;
-- Revert password column
ALTER TABLE users
ALTER COLUMN password_hash TYPE VARCHAR(100),
    RENAME COLUMN password_hash TO user_password;
-- Revert email column
ALTER TABLE users
    RENAME COLUMN email TO user_mail;
COMMIT;
