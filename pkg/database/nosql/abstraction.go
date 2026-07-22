package nosql

import "context"

type BaseDocumentDatabaseDriver interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	InsertDocument(ctx context.Context, collection string, doc any) (string, error)
	FindDocumentByID(ctx context.Context, collection string, id string) (any, error)
	UpdateDocument(ctx context.Context, collection string, id string, update any) (bool, error)
	DeleteDocument(ctx context.Context, collection string, id string) (bool, error)
	GetDriverName() string
}
