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
    ),
    (
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
    ),
    (
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
    ),
    (
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
