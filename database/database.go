package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"scrapebatch-controller-go/model"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

var db *gorm.DB

func InitDatabase() {
	mysqlPass := os.Getenv("MYSQL_PASS")
	userName := "batchuser"
	dbName := "batchjobs"
	database, err := gorm.Open("mysql", userName+":"+mysqlPass+"@/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	database.DB().SetConnMaxLifetime(time.Hour * 1)
	database.DB().SetMaxIdleConns(10)
	database.DB().SetMaxOpenConns(100)
	db = database
}

func Close() {
	err := db.Close()
	if err != nil {
		_ = fmt.Errorf("%s", "Error closing database connection pool.")
	}
}

// Update the status on a single jobPosting. Returns the updated jobPosting.
func UpdateJobStatus(id uint64, status string) model.JobPosting {
	db.Table("job_posting").
		Where("id = ?", id).
		Update("status", status)

	jobPosting := model.JobPosting{}
	db.Table("job_posting").
		Where("id = ?", id).
		First(&jobPosting)
	return jobPosting
}

// Update all the jobs with a new status. Returns all the jobPostings.
func UpdateMultipleJobStatuses(ids []uint64, status string) []model.JobPosting {
	db.Table("job_posting").
		Where("id IN (?)", ids).
		Update("status", status)

	var jobPostings []model.JobPosting
	db.Table("job_posting").
		Where("id IN (?)", ids).
		Find(&jobPostings)

	return jobPostings
}

// Get all the blacklisted companies. Returns a list of strings.
func GetBlacklistedCompanies() []model.BlacklistedCompany {
	var companies []model.BlacklistedCompany
	db.Table("blacklisted_company").
		Find(&companies)
	return companies
}

// Adds a blacklisted company. Returns the blacklisted companies.
func AddBlacklistedCompany(company model.BlacklistedCompany) []model.BlacklistedCompany {
	exists, _ := FindBlacklistedCompany(company)
	if exists == false {
		db.Table("blacklisted_company").
			Create(&company)
	}
	return GetBlacklistedCompanies()
}

func FindBlacklistedCompany(company model.BlacklistedCompany) (bool, model.BlacklistedCompany) {
	var existing model.BlacklistedCompany
	db.Table("blacklisted_company").
		Where("name = ?", company.Name).
		First(&existing)
	var exists = existing.Name != ""
	return exists, existing
}

// Removes a blacklisted company. Returns the blacklisted companies.
func RemoveBlacklistedCompany(company model.BlacklistedCompany) []model.BlacklistedCompany {
	exists, existing := FindBlacklistedCompany(company)
	if !exists {
		return GetBlacklistedCompanies()
	}
	db.Table("blacklisted_company").
		Delete(&existing)
	return GetBlacklistedCompanies()
}

func GetAllJobPostings() []model.JobPosting {
	var jobPostings []model.JobPosting
	db.Table("job_posting").Find(&jobPostings)
	return jobPostings
}

// Updates the job posting with whatever data it currently has. Returns the updated jobPosting.
func UpdateJobPosting(jobPosting model.JobPosting) model.JobPosting {
	db.Table("job_posting").
		Update(&jobPosting)
	return jobPosting
}

// Get all the blocked titles. Returns a list of strings.
func GetBlockedTitles() []string {
	var titles []string
	db.Table("block_title").
		Find(&titles)
	return titles
}

// Get a jobPosting by id.
func GetJobPostingById(id uint64) model.JobPosting {
	var jobPosting model.JobPosting
	db.Table("job_posting").
		Where("id = ?", id).
		First(&jobPosting)
	return jobPosting
}
