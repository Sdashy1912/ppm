package stats

import (
	"ppm/libnvth/internal/database"
	"gopkg.in/mgo.v2/bson"
)

// DAO dao for stats
type DAO interface {
	Stats() (map[string]interface{}, error)
}

// DAOImpl implements...
type DAOImpl struct {
	session *database.DBSession
}

// NewDAO return new statdao
func NewDAO(session *database.DBSession) DAO {
	return DAOImpl{session: session}
}

var _ DAO = (*DAOImpl)(nil)

// Stats lets stats =))
func (dao DAOImpl) Stats() (map[string]interface{}, error) {
	result := make(map[string]interface{})
	collection := dao.session.Collection("Customers")
	defer collection.Close()
	// Customers by industries
	data := []bson.M{}
	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id":       "$industry",
				"customers": bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"_id":       false,
				"industry":  "$_id",
				"customers": true,
			},
		},
	}
	err := collection.Pipe(pipeline).All(&data)
	if err != nil {
		return nil, err
	}
	result["stats_by_industries"] = data
	// Vulnerabilities by risk rating
	data = []bson.M{}
	tgCollection := dao.session.Collection("Targets")
	defer tgCollection.Close()
	pipeline = []bson.M{
		{"$unwind": "$vulnerabilities"},
		{
			"$group": bson.M{
				"_id":             "$vulnerabilities.rating",
				"vulnerabilities": bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"_id":             false,
				"rating":          "$_id",
				"vulnerabilities": true,
			},
		},
	}
	err = tgCollection.Pipe(pipeline).All(&data)
	if err != nil {
		return nil, err
	}
	result["stats_by_vul_ratings"] = data
	// requests count
	data = []bson.M{}
	pipeline = []bson.M{
		{"$unwind": "$details"},
	}
	err = tgCollection.Pipe(pipeline).All(&data)
	if err != nil {
		return nil, err
	}
	result["requests"] = len(data)
	// Top 10 vulnerabilities
	data = []bson.M{}
	pipeline = []bson.M{
		{"$unwind": "$vulnerabilities"},
		{
			"$group": bson.M{
				"_id": bson.M{
					"template_id":     "$vulnerabilities.template_id",
					"identifier_name": "$vulnerabilities.identifier_name",
					"rating":          "$vulnerabilities.rating",
					"name":            "$vulnerabilities.name",
				},
				"count": bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"_id":             false,
				"rating":          "$_id.rating",
				"template_id":     "$_id.template_id",
				"identifier_name": "$_id.identifier_name",
				"name":            "$_id.name",
				"count":           true,
			},
		},
		{
			"$sort": bson.M{"count": -1},
		},
	}
	err = tgCollection.Pipe(pipeline).All(&data)
	if err != nil {
		return nil, err
	}
	if len(data) > 10 {
		data = data[:10]
	}
	result["stats_top_10_vuls"] = data
	// Vul stats by industries
	data = []bson.M{}
	pipeline = []bson.M{
		{
			"$lookup": bson.M{
				"from":         "Scopes",
				"localField":   "scope_id",
				"foreignField": "_id",
				"as":           "scope"},
		},
		{"$unwind": "$scope"},
		{
			"$lookup": bson.M{
				"from":         "Projects",
				"localField":   "scope.project_id",
				"foreignField": "_id",
				"as":           "project"},
		},
		{"$unwind": "$project"},
		{
			"$lookup": bson.M{
				"from":         "Customers",
				"localField":   "project.customer_id",
				"foreignField": "_id",
				"as":           "customer"},
		},
		{"$unwind": "$customer"},
		{"$unwind": "$vulnerabilities"},
		{"$group": bson.M{
			"_id":    "$customer.industry",
			"high":   bson.M{"$sum": bson.M{"$cond": []interface{}{bson.M{"$eq": []string{"$vulnerabilities.rating", "High"}}, 1, 0}}},
			"medium": bson.M{"$sum": bson.M{"$cond": []interface{}{bson.M{"$eq": []string{"$vulnerabilities.rating", "Medium"}}, 1, 0}}},
			"low":    bson.M{"$sum": bson.M{"$cond": []interface{}{bson.M{"$eq": []string{"$vulnerabilities.rating", "Low"}}, 1, 0}}},
			"info":   bson.M{"$sum": bson.M{"$cond": []interface{}{bson.M{"$eq": []string{"$vulnerabilities.rating", "Info"}}, 1, 0}}},
		}},
	}
	err = tgCollection.Pipe(pipeline).All(&data)
	if err != nil {
		return nil, err
	}
	result["vuls_by_rating_by_industry"] = data
	return result, nil
}
