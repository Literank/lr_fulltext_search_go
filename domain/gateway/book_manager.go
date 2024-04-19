/*
Package gateway contains all domain gateways.
*/
package gateway

import (
	"context"

	"literank.com/fulltext-books/domain/model"
)

// BookManager manages all books
type BookManager interface {
	IndexBook(ctx context.Context, b *model.Book) (string, error)
}
