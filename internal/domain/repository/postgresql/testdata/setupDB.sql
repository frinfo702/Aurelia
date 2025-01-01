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
    application_status VARCHAR(20) DEFAULT 'pending' CHECK (
        application_status IN ('pending', 'approved', 'rejected')
    )
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
        application_status
    )
VALUES (
        'Alpha Inc',
        'A leading tech startup',
        120,
        'Open culture',
        '2020-01-01',
        'https://alpha.example.com',
        'San Francisco',
        '51-100',
        '2M',
        'Startup',
        'Software,AI',
        'approved'
    ),
    (
        'Beta Solutions',
        'Focus on beta testing',
        45,
        'Flexible hours',
        '2019-03-15',
        'http://beta.solutions',
        'London,Remote',
        '21-50',
        '500K',
        'LLC',
        'Testing,Consulting',
        'pending'
    ),
    (
        'Gamma Corp',
        'Global services provider',
        300,
        'Team building events',
        '2018-07-20',
        'https://gamma.corp',
        'New York',
        '101-200',
        '5M',
        'Corporation',
        'Finance,Cloud',
        'rejected'
    ),
    (
        'Delta Innovations',
        'Innovative tech products',
        80,
        'Research incentives',
        '2021-11-10',
        'http://delta.io',
        'Berlin',
        '51-100',
        '1M',
        'Startup',
        'AR,VR',
        'pending'
    ),
    (
        'Epsilon Ventures',
        'Venture capital firm',
        25,
        'Equity sharing',
        '2017-02-28',
        'http://epsilon.vc',
        'Tokyo',
        '1-20',
        '10M',
        'VC',
        'Investment,Finance',
        'pending'
    );
-- jobs
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
        'intern',
        'React',
        300000,
        'Backend',
        'Go experience',
        'Go, Docker'
    ),
    (
        1,
        'fulltime',
        'Golang',
        500000,
        'Frontend',
        'React experience',
        'React, TypeScript'
    ),
    (
        2,
        'intern',
        'Python, ML',
        280000,
        'Mobile',
        'Swift experience',
        'Swift, iOS'
    );
