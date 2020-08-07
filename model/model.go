package model

type BlacklistedCompany struct {
	Id   uint64 `json:"-" gorm:"AUTO_INCREMENT;column:id"`
	Name string `json:"name" gorm:"column:name"`
}
type BlockTitle string
type Status string

type Companies struct {
	Companies []Company `json:"companies"`
}

type Company struct {
	Id          uint64       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	JobPostings []JobPosting `json:"jobPostings,omitempty"`
}

type JobPosting struct {
	Id          uint64 `json:"id,omitempty"`
	JobTitle    string `json:"jobTitle,omitempty"`
	Tags        string `json:"tags,omitempty"`
	Href        string `json:"href,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Company     string `json:"company,omitempty"`
	Location    string `json:"location,omitempty"`
	Date        string `json:"date,omitempty"`
	Salary      string `json:"salary,omitempty"`
	JobSite     string `json:"jobSite,omitempty"`
	Description string `json:"description,omitempty"`
	RemoteText  string `json:"remoteText,omitempty"`
	MiscText    string `json:"miscText,omitempty"`
	Status      string `json:"status,omitempty"`
}
