INSERT INTO companies (company_name, company_overview)
VALUES ('Test Company 1', 'Overview 1'),
    ('Test Company 2', 'Overview 2');
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
