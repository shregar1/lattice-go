package nosql

import "context"

type InMemoryDocumentDriver struct{}

func (i *InMemoryDocumentDriver) Connect(ctx context.Context) error { return nil }
func (i *InMemoryDocumentDriver) Disconnect(ctx context.Context) error { return nil }
func (i *InMemoryDocumentDriver) InsertDocument(ctx context.Context, collection string, doc any) (string, error) { return "doc_123", nil }
func (i *InMemoryDocumentDriver) FindDocumentByID(ctx context.Context, collection string, id string) (any, error) { return nil, nil }
func (i *InMemoryDocumentDriver) UpdateDocument(ctx context.Context, collection string, id string, update any) (bool, error) { return true, nil }
func (i *InMemoryDocumentDriver) DeleteDocument(ctx context.Context, collection string, id string) (bool, error) { return true, nil }
func (i *InMemoryDocumentDriver) GetDriverName() string { return "in_memory_doc" }
