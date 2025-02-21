package bean

import "gopkg.in/mgo.v2/bson"

// TargetBean bean
type TargetBean struct {
	ID              bson.ObjectId       `json:"id" bson:"_id"`
	Name            string              `json:"name" bson:"name"`
	Description     string              `json:"description" bson:"description"`
	Assignee        AssigneeBean        `json:"assignee" bson:"assignee"`
	Platform        string              `json:"platform" bson:"platform"`
	Project         ProjectBean         `json:"project" bson:"project"`
	Scope           ScopeBean           `json:"scope" bson:"scope"`
	Requests        int                 `json:"requests" bson:"requests"`
	Vulnerabilities []VulnerabilityBean `json:"vulnerabilities" bson:"vulnerabilities"`
}
