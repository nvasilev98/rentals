package rentals

import (
	"fmt"
	"strings"
)

const withinMiles = 100

func buildSQLQuery(query string, clauses map[string][]string) string {
	return addPagination(addSorting(addWhereClause(query, clauses), clauses), clauses)
}

func addWhereClause(query string, clauses map[string][]string) string {
	addedConditions := 0
	resultQuery := fmt.Sprintf("%s WHERE", query)
	priceMin, ok := clauses["price_min"]
	if ok {
		resultQuery = fmt.Sprintf("%s r.price_per_day >= %s", resultQuery, priceMin[0])
		addedConditions++
	}

	priceMax, ok := clauses["price_max"]
	if ok {
		if addedConditions > 0 {
			resultQuery = fmt.Sprintf("%s AND", resultQuery)
		}

		resultQuery = fmt.Sprintf("%s r.price_per_day <= %s", resultQuery, priceMax[0])
		addedConditions++
	}

	ids, ok := clauses["ids"]
	if ok {
		if addedConditions > 0 {
			resultQuery = fmt.Sprintf("%s AND", resultQuery)
		}

		resultQuery = fmt.Sprintf("%s r.id IN (%s)", resultQuery, ids[0])
		addedConditions++
	}

	near, ok := clauses["near"]
	if ok {
		if addedConditions > 0 {
			resultQuery = fmt.Sprintf("%s AND", resultQuery)
		}
		coordinates := strings.Split(near[0], ",")
		distanceFormula := "%s (3959 * acos(cos(radians(%s)) * cos(radians(lat)) * cos(radians(lng) - radians(%s)) + sin(radians(%s)) * sin(radians(lat)))) < %d"
		resultQuery = fmt.Sprintf(distanceFormula, resultQuery, coordinates[0], coordinates[1], coordinates[0], withinMiles)
		addedConditions++
	}

	if addedConditions == 0 {
		return query
	}

	return resultQuery
}

func addPagination(query string, clauses map[string][]string) string {
	resultQuery := query
	offset, ok := clauses["offset"]
	if ok {
		resultQuery = fmt.Sprintf("%s OFFSET %s", resultQuery, offset[0])
	}

	limit, ok := clauses["limit"]
	if ok {
		resultQuery = fmt.Sprintf("%s LIMIT %s", resultQuery, limit[0])
	}

	return resultQuery
}

func addSorting(query string, clauses map[string][]string) string {
	resultQuery := query
	sort, ok := clauses["sort"]
	if ok {
		// if the column does not exist, a order by clause will not be added
		if column, ok := toDBColumnName(sort[0]); ok {
			resultQuery = fmt.Sprintf("%s ORDER BY %s", resultQuery, column)
		}
	}

	return resultQuery
}

func toDBColumnName(key string) (string, bool) {
	//supports price and year sorting, adding more columns will extend sorting options
	columns := map[string]string{
		"price": "price_per_day",
		"year":  "vehicle_year",
	}
	dbColumnName, ok := columns[key]

	return dbColumnName, ok
}
