BEGIN;

-- Modify user table columns
ALTER TABLE users 
    RENAME COLUMN user_mail TO email;

ALTER TABLE users
    ALTER COLUMN user_password TYPE CHAR(60),
    RENAME COLUMN user_password TO password_hash;

-- Add timestamps
ALTER TABLE users
    ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- Add trigger for updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

COMMIT;
