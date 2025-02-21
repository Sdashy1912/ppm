package model

import "gopkg.in/mgo.v2/bson"

// Target target
type Target struct {
	ID              bson.ObjectId   `json:"id" bson:"_id"`
	Name            string          `json:"name" bson:"name"`
	Description     string          `json:"description" bson:"description"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities" bson:"vulnerabilities"`
}

func (target Target) GetVulByID(id string) (Vulnerability, bool) {
	if target.Vulnerabilities == nil {
		return Vulnerability{}, false
	}
	for _, vul := range target.Vulnerabilities {
		if vul.ID.Hex() == id {
			return vul, true
		}
	}
	return Vulnerability{}, false
}