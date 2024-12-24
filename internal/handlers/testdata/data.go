package testdata

import (
    "time"

    "Aurelia/internal/models"
)

var jobTestData = []models.Job{
	{
		JobID:          1,
		CompanyID:      1,
		HiringType:     "intern",
		TechnologyType: "Software Engineering",
		IncomeRange:    1500000,
		JobTag:         "backend",
		Requirements:   "Experience with distributed system",
		UsedTechnology: "Go, Kubernetes",
	},
	{
		JobID:          2,
		CompanyID:      1,
		HiringType:     "fulltime",
		TechnologyType: "Data Science",
		IncomeRange:    2000000,
		JobTag:         "data",
		Requirements:   "Knowledge of machine learning frameworks",
		UsedTechnology: "Python, TensorFlow",
	},
	{
		JobID:          3,
		CompanyID:      2,
		HiringType:     "parttime",
		TechnologyType: "Frontend Development",
		IncomeRange:    1200000,
		JobTag:         "frontend",
		Requirements:   "Experience with React or Vue",
		UsedTechnology: "JavaScript, React",
	},
	{
		JobID:          4,
		CompanyID:      2,
		HiringType:     "contract",
		TechnologyType: "DevOps",
		IncomeRange:    1800000,
		JobTag:         "ops",
		Requirements:   "CI/CD pipeline knowledge",
		UsedTechnology: "Docker, Jenkins",
	},
	{
		JobID:          5,
		CompanyID:      3,
		HiringType:     "fulltime",
		TechnologyType: "Mobile Development",
		IncomeRange:    2200000,
		JobTag:         "mobile",
		Requirements:   "Experience with native iOS/Android development",
		UsedTechnology: "Swift, Kotlin",
	},
}

var companyTestData = []models.Company{
    {
        CompanyID:         1,
        CompanyName:       "TechNova Solutions",
        CompanyOverview:   "Leading provider of AI solutions",
        WorkingPeopleID:   250,
        CultureAndBenefit: "Flexible hours, health benefits",
        EstablishDate:     time.Date(2015, 6, 15, 0, 0, 0, 0, time.UTC),
        CompanyWebsite:    "https://www.technova.com",
        CompanyLocations:  "San Francisco, New York, London",
        CompanySize:       "201-500 employees",
        TotalRaised:       "$50M",
        CompanyType:       "Private",
        CompanyMarkets:    "Software, AI",
        IsAuthorized:      true,
    },
    {
        CompanyID:         2,
        CompanyName:       "GreenEarth Industries",
        CompanyOverview:   "Eco-friendly manufacturing solutions",
        WorkingPeopleID:   120,
        CultureAndBenefit: "Community support, remote-friendly",
        EstablishDate:     time.Date(2010, 4, 1, 0, 0, 0, 0, time.UTC),
        CompanyWebsite:    "https://www.greenearth.co",
        CompanyLocations:  "Portland, Seattle",
        CompanySize:       "51-200 employees",
        TotalRaised:       "$20M",
        CompanyType:       "Private",
        CompanyMarkets:    "Manufacturing, Sustainability",
        IsAuthorized:      false,
    },
    {
        CompanyID:         3,
        CompanyName:       "HealthWave Inc.",
        CompanyOverview:   "Advanced telemedicine platform",
        WorkingPeopleID:   340,
        CultureAndBenefit: "On-site gym, mental health support",
        EstablishDate:     time.Date(2018, 9, 10, 0, 0, 0, 0, time.UTC),
        CompanyWebsite:    "https://www.healthwave.com",
        CompanyLocations:  "Chicago, Boston",
        CompanySize:       "201-500 employees",
        TotalRaised:       "$75M",
        CompanyType:       "Public",
        CompanyMarkets:    "Healthcare, Telemedicine",
        IsAuthorized:      true,
    },
    {
        CompanyID:         4,
        CompanyName:       "EduNext Learning",
        CompanyOverview:   "Online education and remote learning solutions",
        WorkingPeopleID:   150,
        CultureAndBenefit: "Tuition reimbursement, flexible schedules",
        EstablishDate:     time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC),
        CompanyWebsite:    "https://www.edunext.com",
        CompanyLocations:  "Remote",
        CompanySize:       "51-200 employees",
        TotalRaised:       "$10M",
        CompanyType:       "Startup",
        CompanyMarkets:    "Education, E-Learning",
        IsAuthorized:      false,
    },
    {
        CompanyID:         5,
        CompanyName:       "FinSecure LLC",
        CompanyOverview:   "Innovative cybersecurity for financial institutions",
        WorkingPeopleID:   80,
        CultureAndBenefit: "Competitive salaries, flexible PTO",
        EstablishDate:     time.Date(2013, 3, 20, 0, 0, 0, 0, time.UTC),
        CompanyWebsite:    "https://www.finsecure.com",
        CompanyLocations:  "New York, Los Angeles",
        CompanySize:       "51-200 employees",
        TotalRaised:       "$30M",
        CompanyType:       "Private",
        CompanyMarkets:    "Cybersecurity, Finance",
        IsAuthorized:      true,
    },
    {
        CompanyID:         6,
        CompanyName:       "DataMinds Analytics",
        CompanyOverview:   "Data analytics and BI consulting",
        WorkingPeopleID:   40,
        CultureAndBenefit: "Flat organization, remote opportunities",
        EstablishDate:     time.Date(2016, 11, 15, 0, 0, 0, 0, time.UTC),
        CompanyWebsite:    "https://www.dataminds.io",
        CompanyLocations:  "Austin, Denver",
        CompanySize:       "11-50 employees",
        TotalRaised:       "$5M",
        CompanyType:       "Consulting",
        CompanyMarkets:    "Data, Analytics",
        IsAuthorized:      false,
    },
    {
        CompanyID:         7,
        CompanyName:       "CloudSphere Systems",
        CompanyOverview:   "Cloud computing and infrastructure automation",
        WorkingPeopleID:   500,
        CultureAndBenefit: "Global mobility, bonus incentives",
        EstablishDate:     time.Date(2012, 12, 25, 0, 0, 0, 0, time.UTC),
        CompanyWebsite:    "https://www.cloudsphere.com",
        CompanyLocations:  "San Jose, Tokyo, Berlin",
        CompanySize:       "501-1000 employees",
        TotalRaised:       "$100M",
        CompanyType:       "Public",
        CompanyMarkets:    "Cloud, Infrastructure",
        IsAuthorized:      true,
    },
    {
        CompanyID:         8,
        CompanyName:       "InnovateX Labs",
        CompanyOverview:   "R&D for cutting-edge tech prototypes",
        WorkingPeopleID:   60,
        CultureAndBenefit: "A playful culture promoting creativity",
        EstablishDate:     time.Date(2019, 5, 1, 0, 0, 0, 0, time.UTC),
        CompanyWebsite:    "https://www.innovatex.com",
        CompanyLocations:  "Remote",
        CompanySize:       "11-50 employees",
        TotalRaised:       "$2M",
        CompanyType:       "Startup",
        CompanyMarkets:    "Research, Prototyping",
        IsAuthorized:      false,
    },
    {
        CompanyID:         9,
        CompanyName:       "FutureTech Ventures",
        CompanyOverview:   "Futuristic investments and accelerator programs",
        WorkingPeopleID:   25,
        CultureAndBenefit: "Equity benefits, coworking spaces",
        EstablishDate:     time.Date(2017, 2, 14, 0, 0, 0, 0, time.UTC),
        CompanyWebsite:    "https://www.futuretechv.com",
        CompanyLocations:  "Silicon Valley",
        CompanySize:       "11-50 employees",
        TotalRaised:       "$15M",
        CompanyType:       "VC",
        CompanyMarkets:    "Investments, Startups",
        IsAuthorized:      true,
    },
    {
        CompanyID:         10,
        CompanyName:       "BlueSky Robotics",
        CompanyOverview:   "Robotics solutions for industrial automation",
        WorkingPeopleID:   300,
        CultureAndBenefit: "Free lunches, creative freedom",
        EstablishDate:     time.Date(2015, 9, 9, 0, 0, 0, 0, time.UTC),
        CompanyWebsite:    "https://www.blueskyrobotics.com",
        CompanyLocations:  "Chicago, Toronto",
        CompanySize:       "201-500 employees",
        TotalRaised:       "$40M",
        CompanyType:       "Private",
        CompanyMarkets:    "Robotics, Automation",
        IsAuthorized:      false,
    },
}
