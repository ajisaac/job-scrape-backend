package main

import (
	"github.com/gin-gonic/gin"
	"scrapebatch-controller-go/database"
	"scrapebatch-controller-go/model"
	"scrapebatch-controller-go/service"
	"strconv"
)

func main() {
	database.InitDatabase()
	defer database.Close()

	r := gin.Default()
	r.Use(corsResponse())

	// get all jobs by company
	// returns an array of jobs grouped by company
	r.GET("/jobs/all/bycompany", func(c *gin.Context) {
		companies := service.GetAllJobsByCompany()
		c.JSON(200, companies)
	})

	// update the status of multiple jobs
	// takes a list of ids
	// takes a status as string
	// returns the job postings
	r.PUT("/jobs/status/updatemultiple/:status", func(c *gin.Context) {
		status := c.Param("status")
		if status == "" {
			c.JSON(400, gin.H{
				"error": "status must not be blank",
			})
		}

		var ids []uint64
		err := c.Bind(&ids)

		if err != nil {
			c.JSON(400, gin.H{
				"error": "unable to parse body",
			})
		}

		jobPostings := service.UpdateMultipleJobStatuses(ids, status)
		c.JSON(200, jobPostings)
	})

	// update the status of a particular job
	// takes an id and status
	// returns the updated job posting
	r.PUT("/jobs/status/update/:id/:status", func(c *gin.Context) {
		idStr := c.Param("id")
		if idStr == "" {
			c.JSON(400, gin.H{
				"error": "id must not be blank",
			})
		}
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "id must be numeric",
			})
		}
		status := c.Param("status")
		if status == "" {
			c.JSON(400, gin.H{
				"error": "status must not be blank",
			})
		}
		jobPosting := service.UpdateJobStatus(id, status)

		c.JSON(200, jobPosting)
	})

	// get the blacklisted companies
	// returns the list of companies
	r.GET("/jobs/blacklistedcompanies", func(c *gin.Context) {
		blacklistedCompanies := service.GetBlacklistedCompanies()
		c.JSON(200, blacklistedCompanies)
	})

	// blacklist a company
	// takes a company as string
	// returns the list of companies
	r.POST("/jobs/blacklistcompany", func(c *gin.Context) {
		var company model.BlacklistedCompany
		err := c.Bind(&company)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err,
			})
		}
		blacklistedCompanies := service.AddBlacklistedCompany(company)
		c.JSON(200, blacklistedCompanies)
	})

	// unblacklist a company
	// takes a company as string
	// returns the list of companies
	r.POST("/jobs/blacklistcompanyremove", func(c *gin.Context) {
		var company model.BlacklistedCompany
		err := c.Bind(&company)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err,
			})
		}
		blacklistedCompanies := service.RemoveBlacklistedCompany(company)
		c.JSON(200, blacklistedCompanies)
	})

	_ = r.Run()
}

func corsResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
