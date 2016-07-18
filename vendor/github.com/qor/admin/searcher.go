package admin

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/qor"
)

type scopeFunc func(db *gorm.DB, context *qor.Context) *gorm.DB

// Pagination is used to hold pagination related information when rendering tables
type Pagination struct {
	Total       int
	Pages       int
	CurrentPage int
	PerPage     int
}

// Searcher is used to search results
type Searcher struct {
	*Context
	scopes     []*Scope
	filters    map[string]string
	Pagination Pagination
}

func (s *Searcher) clone() *Searcher {
	return &Searcher{Context: s.Context, scopes: s.scopes, filters: s.filters}
}

// Page set current page, if current page equal -1, then show all records
func (s *Searcher) Page(num int) *Searcher {
	s.Pagination.CurrentPage = num
	return s
}

// PerPage set pre page count
func (s *Searcher) PerPage(num int) *Searcher {
	s.Pagination.PerPage = num
	return s
}

// Scope filter with defined scopes
func (s *Searcher) Scope(names ...string) *Searcher {
	newSearcher := s.clone()
	for _, name := range names {
		for _, scope := range s.Resource.scopes {
			if scope.Name == name && !scope.Default {
				newSearcher.scopes = append(newSearcher.scopes, scope)
				break
			}
		}
	}
	return newSearcher
}

// Filter filter with defined filters, filter with columns value
func (s *Searcher) Filter(name, query string) *Searcher {
	newSearcher := s.clone()
	if newSearcher.filters == nil {
		newSearcher.filters = map[string]string{}
	}
	newSearcher.filters[name] = query
	return newSearcher
}

// FindMany find many records based on current conditions
func (s *Searcher) FindMany() (interface{}, error) {
	context := s.parseContext()
	result := s.Resource.NewSlice()
	err := s.Resource.CallFindMany(result, context)
	return result, err
}

// FindOne find one record based on current conditions
func (s *Searcher) FindOne() (interface{}, error) {
	context := s.parseContext()
	result := s.Resource.NewStruct()
	err := s.Resource.CallFindOne(result, nil, context)
	return result, err
}

var filterRegexp = regexp.MustCompile(`^filters\[(.*?)\]$`)

func (s *Searcher) callScopes(context *qor.Context) *qor.Context {
	db := context.GetDB()

	// call default scopes
	for _, scope := range s.Resource.scopes {
		if scope.Default {
			db = scope.Handle(db, context)
		}
	}

	// call scopes
	for _, scope := range s.scopes {
		db = scope.Handle(db, context)
	}

	// call filters
	if s.filters != nil {
		for key, value := range s.filters {
			filter := s.Resource.filters[key]
			if filter != nil && filter.Handler != nil {
				db = filter.Handler(key, value, db, context)
			} else {
				db = defaultFilterHandler(key, value, db, context)
			}
		}
	}

	// add order by
	if orderBy := context.Request.Form.Get("order_by"); orderBy != "" {
		if regexp.MustCompile("^[a-zA-Z_]+$").MatchString(orderBy) {
			if field, ok := db.NewScope(s.Context.Resource.Value).FieldByName(strings.TrimSuffix(orderBy, "_desc")); ok {
				if strings.HasSuffix(orderBy, "_desc") {
					db = db.Order(field.DBName+" DESC", true)
				} else {
					db = db.Order(field.DBName, true)
				}
			}
		}
	}
	
	context.SetDB(db)

	// call search
	var keyword string
	if keyword = context.Request.Form.Get("keyword"); keyword == "" {
		keyword = context.Request.URL.Query().Get("keyword")
	}

	if keyword != "" && s.Resource.SearchHandler != nil {
		context.SetDB(s.Resource.SearchHandler(keyword, context))
		return context
	}

	return context
}

func (s *Searcher) parseContext() *qor.Context {
	var (
		searcher = s.clone()
		context  = searcher.Context.Context.Clone()
	)

	if context != nil && context.Request != nil {
		// parse scopes
		scopes := context.Request.Form["scopes"]
		searcher = searcher.Scope(scopes...)

		// parse filters
		for key, value := range context.Request.Form {
			if matches := filterRegexp.FindStringSubmatch(key); len(matches) > 0 {
				searcher = searcher.Filter(matches[1], value[0])
			}
		}
	}

	searcher.callScopes(context)

	db := context.GetDB()

	// pagination
	context.SetDB(db.Model(s.Resource.Value).Set("qor:getting_total_count", true))
	s.Resource.CallFindMany(&s.Pagination.Total, context)

	if s.Pagination.CurrentPage == 0 {
		if s.Context.Request != nil {
			if page, err := strconv.Atoi(s.Context.Request.Form.Get("page")); err == nil {
				s.Pagination.CurrentPage = page
			}
		}

		if s.Pagination.CurrentPage == 0 {
			s.Pagination.CurrentPage = 1
		}
	}

	if s.Pagination.PerPage == 0 {
		if perPage, err := strconv.Atoi(s.Context.Request.Form.Get("per_page")); err == nil {
			s.Pagination.PerPage = perPage
		} else {
			s.Pagination.PerPage = s.Resource.Config.PageCount
		}
	}

	if s.Pagination.CurrentPage > 0 {
		s.Pagination.Pages = (s.Pagination.Total-1)/s.Pagination.PerPage + 1

		db = db.Limit(s.Pagination.PerPage).Offset((s.Pagination.CurrentPage - 1) * s.Pagination.PerPage)
	}

	context.SetDB(db)

	return context
}
