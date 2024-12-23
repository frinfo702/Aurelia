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
INSERT INTO companies (
        company_name,
        company_overview,
        working_people_id,
        culture_and_benefit,
        establish_date,
        company_website,
        company_locations,
        company_size,
        total_raised,
        company_type,
        company_markets,
        is_authorized
    )
VALUES (
        'TechNova Solutions',
        'A leading provider of innovative tech solutions specializing in AI and machine learning applications.',
        250,
        'Flexible working hours, comprehensive health benefits, and continuous learning opportunities.',
        '2015-06-15',
        'https://www.technova.com',
        'San Francisco, New York, London',
        '201-500 employees',
        '$50M',
        'Private',
        'Artificial Intelligence, Machine Learning, Software Development',
        TRUE
    );
INSERT INTO companies (
        company_name,
        company_overview,
        working_people_id,
        culture_and_benefit,
        establish_date,
        company_website,
        company_locations,
        company_size,
        total_raised,
        company_type,
        company_markets,
        is_authorized
    )
VALUES (
        'GreenEarth Industries',
        'Committed to sustainable manufacturing and eco-friendly products for a greener planet.',
        120,
        'Remote work options, wellness programs, and eco-friendly office spaces.',
        '2010-09-01',
        'https://www.greenearth.com',
        'Berlin, Amsterdam, Toronto',
        '51-200 employees',
        '$20M',
        'Public',
        'Sustainable Manufacturing, Eco-Friendly Products',
        FALSE
    );
INSERT INTO companies (
        company_name,
        company_overview,
        working_people_id,
        culture_and_benefit,
        establish_date,
        company_website,
        company_locations,
        company_size,
        total_raised,
        company_type,
        company_markets,
        is_authorized
    )
VALUES (
        'HealthWave Inc.',
        'Providing cutting-edge healthcare solutions and telemedicine services to improve patient outcomes.',
        500,
        'Competitive salaries, health insurance, and opportunities for career advancement.',
        '2008-03-22',
        'https://www.healthwave.com',
        'Chicago, Miami, Seattle',
        '501-1000 employees',
        '$150M',
        'Private',
        'Healthcare, Telemedicine',
        TRUE
    );
INSERT INTO companies (
        company_name,
        company_overview,
        working_people_id,
        culture_and_benefit,
        establish_date,
        company_website,
        company_locations,
        company_size,
        total_raised,
        company_type,
        company_markets,
        is_authorized
    )
VALUES (
        'EduNext Learning',
        'An online education platform offering a wide range of courses and professional certifications.',
        80,
        'Inclusive work environment, tuition reimbursement, and flexible schedules.',
        '2018-11-05',
        'https://www.edunext.com',
        'Austin, Denver, Boston',
        '11-50 employees',
        '$5M',
        'Startup',
        'Online Education, Professional Development',
        FALSE
    );
INSERT INTO companies (
        company_name,
        company_overview,
        working_people_id,
        culture_and_benefit,
        establish_date,
        company_website,
        company_locations,
        company_size,
        total_raised,
        company_type,
        company_markets,
        is_authorized
    )
VALUES (
        'FinSecure LLC',
        'Specializing in cybersecurity solutions for financial institutions and safeguarding digital assets.',
        300,
        'Health and wellness programs, stock options, and a collaborative work culture.',
        '2012-07-30',
        'https://www.finsecure.com',
        'New York, London, Singapore',
        '201-500 employees',
        '$75M',
        'Private',
        'Cybersecurity, Financial Services',
        TRUE
    );
INSERT INTO jobs (
        company_id,
        hiring_type,
        technology_type,
        income_range,
        job_tag,
        requirements,
        used_technology
    )
VALUES (
        1,
        'fulltime',
        'Software Engineering',
        1500000,
        'Backend',
        'Experience with distributed systems',
        'Go, Kubernetes'
    ),
    (
        2,
        'fulltime',
        'Hardware Engineering',
        1700000,
        'Hardware',
        'Experience with circuit design',
        'VHDL, Verilog'
    ),
    (
        3,
        'fulltime',
        'E-commerce Platform',
        1400000,
        'Full Stack',
        'Experience with large-scale web applications',
        'Java, AWS'
    );
