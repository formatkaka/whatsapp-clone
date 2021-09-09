package v1

import (
	"database/sql"
	"encoding/json"
	"whatsapp-clone/db"
	"whatsapp-clone/es"

	"github.com/olivere/elastic/v7"
	"go.elastic.co/apm/module/apmgoredis"
	"golang.org/x/net/context"
)

type SampleService struct {
	version   string
	dbFactory db.DBFactory
	redis     apmgoredis.Client
	es        *elastic.Client
}

func NewSampleService(dbFactory db.DBFactory, redis apmgoredis.Client, es *elastic.Client) *SampleService {
	return &SampleService{
		version:   "v1",
		dbFactory: dbFactory,
		redis:     redis,
		es:        es,
	}
}

func (s *SampleService) Hello() string {
	return "Hello"
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (s *SampleService) FromDb(ctx context.Context) ([]User, error) {
	query := db.WrapQuery("SELECT userId, userName, userEmail FROM users limit 2;")
	rows, err := s.dbFactory("reader").QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var users []User

	for rows.Next() {
		var user User
		// null needs to be handled separately than string
		var nullEmail sql.NullString
		err := rows.Scan(&user.Id, &user.Name, &nullEmail)

		if nullEmail.Valid {
			user.Email = nullEmail.String
		}

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *SampleService) FromRedis(ctx context.Context) (string, error) {
	info, err := s.redis.WithContext(ctx).Info().Result()
	if err != nil {
		return "", err
	}
	return info, nil
}

func (s *SampleService) FromES() ([]map[string]interface{}, error) {
	multiMatchQuery := elastic.NewMultiMatchQuery("pulkit", "Name", "Handle").
		FieldWithBoost("Name", 2).
		Type("phrase_prefix")
	boolQuery := elastic.NewBoolQuery().Should(multiMatchQuery)
	q := elastic.
		NewFunctionScoreQuery().
		Query(boolQuery).
		AddScoreFunc(elastic.NewFieldValueFactorFunction().Modifier("sqrt").Factor(1).Field("Following")).
		AddScoreFunc(elastic.NewFieldValueFactorFunction().Modifier("sqrt").Factor(2.5).Field("Trails")).
		AddScoreFunc(elastic.NewFieldValueFactorFunction().Modifier("sqrt").Factor(5).Field("Followers"))

	searchSource := elastic.NewSearchSource()
	searchSource.Query(q)
	searchResult, err := es.Client().Search().Index("users").Type("users").SearchSource(searchSource).Size(20).Do(context.Background())

	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, hit := range searchResult.Hits.Hits {
		var u map[string]interface{}
		err = json.Unmarshal(hit.Source, &u)
		if err != nil {
			continue
		}
		result = append(result, u)
	}

	return result, nil
}
