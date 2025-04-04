/*
Repository package for data access
*/
package repository

import (
	"database/sql"
	"encoding/json"
	"log"
	"strings"

	"github.com/AkyurekDogan/exinity-task/internal/app/infrastructure/model"
	"github.com/lib/pq"
)

// Match represents the usage database interface
type Match interface {
	Get(filter model.MatchFilter) (model.MatchList, error)
}

type match struct {
	driverRead *sql.DB
}

// NewMatch creates new database interface to access database records
func NewMatch(driverR *sql.DB) Match {
	return &match{
		driverRead: driverR,
	}
}

// Get gets the records regarding to the filter list
func (u *match) Get(filter model.MatchFilter) (model.MatchList, error) {

	// Execute a SELECT query
	rows, err := u.driverRead.Query(`
		select
			p.id,
			p.name,
			ST_X(p."location"::geometry) as lat,
			ST_Y(p."location"::geometry) as long,
			p.radius,
			ST_Distance(
				p."location",
				ST_SetSRID(ST_MakePoint($1, $2), 4326)::GEOGRAPHY
			) AS distance,
			pr.avg,
			ps.craftsmanship_tags
		from public.partner p
		inner join public.skill ps on ps.partner_id = p.id
		left join public.rating pr on pr.partner_id = p.id 
		where ST_DWithin(
			p."location",            
			ST_SetSRID(ST_MakePoint($3, $4), 4326)::GEOGRAPHY,
			p.radius*1000
		) and ps.craftsmanship_tags @> $5
		order by pr.avg desc, distance asc
	`, filter.Loc.Lat, filter.Loc.Long, filter.Loc.Lat, filter.Loc.Long, convertToQueryArrayParameter(filter.MaterialType))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result model.MatchList
	// Iterate through the result set
	i := 1
	for rows.Next() {
		item := model.Match{}
		err := rows.Scan(&item.PartnerId, &item.Name, &item.Loc.Lat, &item.Loc.Long, &item.Radius, &item.Distance, &item.Rating, pq.Array(&item.Skills))
		item.Rank = i
		if err != nil {
			return nil, err
		}
		result = append(result, item)
		i++
	}
	// Check for errors from iterating over rows
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func convertToQueryArrayParameter(param string) string {
	// Convert the parameter to JSON
	jsonParam, err := json.Marshal([]string{param})
	if err != nil {
		log.Fatal(err)
	}
	return strings.ReplaceAll(strings.ReplaceAll(string(jsonParam), "[", "{"), "]", "}")
}
