/*
Package application provides all common structures and functions of the application layer.
*/
package application

import (
	"literank.com/fulltext-books/domain/gateway"
	"literank.com/fulltext-books/infrastructure/config"
	"literank.com/fulltext-books/infrastructure/search"
)

// WireHelper is the helper for dependency injection
type WireHelper struct {
	engine *search.ElasticSearchEngine
}

// NewWireHelper constructs a new WireHelper
func NewWireHelper(c *config.Config) (*WireHelper, error) {
	engine, err := search.NewEngine(c.Search.Address, c.App.PageSize)
	if err != nil {
		return nil, err
	}

	return &WireHelper{engine}, nil
}

// BookManager returns an instance of BookManager
func (w *WireHelper) BookManager() gateway.BookManager {
	return w.engine
}
