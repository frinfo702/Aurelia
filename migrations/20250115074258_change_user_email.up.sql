BEGIN;
-- 1) user_mail → email
ALTER TABLE users
  RENAME COLUMN user_mail TO email;

-- 2) user_password → password_hash かつ型変更
ALTER TABLE users
  ALTER COLUMN user_password TYPE CHAR(60);
ALTER TABLE users
  RENAME COLUMN user_password TO password_hash;

-- 3) UNIQUE制約を確実にする
ALTER TABLE users
  ADD CONSTRAINT users_email_uk UNIQUE (email);
COMMIT;
