package service

import (
	"scrapebatch-controller-go/database"
	"scrapebatch-controller-go/model"
)

func GetAllJobsByCompany() model.Companies {
	jobPostings := database.GetAllJobPostings()

	m := map[string][]model.JobPosting{}
	for _, job := range jobPostings {
		m[job.Company] = append(m[job.Company], job)
	}

	var companies []model.Company
	var id uint64 = 1
	for name, postings := range m {
		companies = append(companies, model.Company{
			Id:          id,
			Name:        name,
			JobPostings: postings,
		})
		id++
	}

	return model.Companies{Companies: companies}
}

func UpdateJobStatus(id uint64, status string) model.JobPosting {
	jobPosting := database.UpdateJobStatus(id, status)
	return jobPosting
}

func UpdateMultipleJobStatuses(ids []uint64, status string) []model.JobPosting {
	jobPostings := database.UpdateMultipleJobStatuses(ids, status)
	return jobPostings
}

func GetBlacklistedCompanies() []string {
	blc := database.GetBlacklistedCompanies()
	var ret []string
	for _, b := range blc {
		ret = append(ret, b.Name)
	}
	if ret == nil {
		ret = []string{}
	}
	return ret
}

func AddBlacklistedCompany(company model.BlacklistedCompany) []string {
	database.AddBlacklistedCompany(company)
	return GetBlacklistedCompanies()
}

func RemoveBlacklistedCompany(company model.BlacklistedCompany) []string {
	database.RemoveBlacklistedCompany(company)
	return GetBlacklistedCompanies()
}
