package watermark

import (
	"context"

	"../"
)

type Service interface {
	// Get the list of all documents
	Get(ctx context.Context, filters ...internal.filters) ([]internal.Document, error)
	Status(ctx context.Context, ticketID string) (internal.Status, error)
	Watermark(ctx context.Context, ticketID, mark string) (int, error)
	AddDocument(ctx context.Context, doc *internal.Document) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}