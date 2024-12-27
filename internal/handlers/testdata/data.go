package testdata

import "Aurelia/internal/domain/models"

var JobTestData = []models.Job{
	models.Job{
		JobID:          1,
		CompanyID:      1,
		HiringType:     "fulltime",
		TechnologyType: "golang",
		IncomeRange:    100000,
		JobTag:         "backend",
		Requirements:   "3 years experience in golang",
		UsedTechnology: "golang",
	},
	models.Job{
		JobID:          2,
		CompanyID:      2,
		HiringType:     "intern",
		TechnologyType: "python",
		IncomeRange:    30000,
		JobTag:         "data science",
		Requirements:   "Basic knowledge of Python",
		UsedTechnology: "pandas, numpy",
	},
	models.Job{
		JobID:          3,
		CompanyID:      3,
		HiringType:     "contract",
		TechnologyType: "javascript",
		IncomeRange:    60000,
		JobTag:         "frontend",
		Requirements:   "2 years experience in React",
		UsedTechnology: "React, Node.js",
	},
	models.Job{
		JobID:          4,
		CompanyID:      4,
		HiringType:     "parttime",
		TechnologyType: "ruby",
		IncomeRange:    50000,
		JobTag:         "backend",
		Requirements:   "1 year experience in Ruby on Rails",
		UsedTechnology: "Rails, PostgreSQL",
	},
	models.Job{
		JobID:          5,
		CompanyID:      1,
		HiringType:     "fulltime",
		TechnologyType: "java",
		IncomeRange:    90000,
		JobTag:         "middleware",
		Requirements:   "4 years experience in Java",
		UsedTechnology: "Spring, Hibernate",
	},
}
