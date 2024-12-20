INSERT INTO companies (company_name, company_overview)
VALUES (
        'Google',
        'A multinational technology company specializing in Internet-related services and products.'
    ),
    (
        'Apple',
        'A technology company that designs, develops, and sells consumer electronics, software, and services.'
    ),
    (
        'Amazon',
        'An e-commerce and cloud computing company.'
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
