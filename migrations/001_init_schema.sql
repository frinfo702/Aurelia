CREATE TABLE IF NOT EXISTS companies (
    company_id SERIAL PRIMARY KEY,
    company_name VARCHAR(50) NOT NULL,
    company_overview TEXT,
    working_people_id INT,
    culture_and_benefit TEXT,
    establish_date DATE,
    company_website VARCHAR(100),
    company_locations TEXT,
    company_size VARCHAR(50),
    total_raised VARCHAR(20),
    company_type VARCHAR(50),
    company_markets TEXT,
    is_authorized BOOLEAN DEFAULT FALSE
);
CREATE TABLE IF NOT EXISTS jobs (
    job_id SERIAL PRIMARY KEY,
    company_id INT NOT NULL REFERENCES companies(company_id),
    hiring_type VARCHAR(20) CHECK (
        hiring_type IN ('intern', 'fulltime', 'parttime', 'contract')
    ),
    technology_type VARCHAR(20),
    income_range INT,
    job_tag TEXT,
    requirements TEXT,
    used_technology TEXT
);
CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR(30),
    user_address VARCHAR(100),
    user_mail VARCHAR(100) UNIQUE,
    -- メールはユニークの方がいい
    user_password VARCHAR(100) -- パスワードはハッシュして保存想定
    -- OAuth対応するなら provider, provider_user_id などのカラム追加も検討
);
CREATE TABLE IF NOT EXISTS inquiries (
    inquiry_id SERIAL PRIMARY KEY,
    company_name VARCHAR(50),
    company_overview TEXT,
    working_people_id INT,
    culture_and_benefit TEXT,
    establish_date DATE,
    company_website VARCHAR(100),
    company_locations TEXT,
    company_size VARCHAR(50),
    total_raised VARCHAR(20),
    company_type VARCHAR(50),
    company_markets TEXT,
    is_authorized BOOLEAN DEFAULT FALSE -- inquiriesは新規掲載申請用
    -- 承認されたらcompaniesテーブルへINSERTするフローを想定
);
CREATE TABLE IF NOT EXISTS auth_tokens (
    token_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(user_id),
    auth_token VARCHAR(200),
    issued_at TIMESTAMP DEFAULT NOW(),
    expires_at TIMESTAMP -- 期限付きのトークンを想定
);
