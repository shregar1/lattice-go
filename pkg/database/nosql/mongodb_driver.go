package nosql

import "context"

type MongoDbDriver struct{}

func (m *MongoDbDriver) Connect(ctx context.Context) error { return nil }
func (m *MongoDbDriver) Disconnect(ctx context.Context) error { return nil }
func (m *MongoDbDriver) InsertDocument(ctx context.Context, collection string, doc any) (string, error) { return "mongo_123", nil }
func (m *MongoDbDriver) FindDocumentByID(ctx context.Context, collection string, id string) (any, error) { return nil, nil }
func (m *MongoDbDriver) UpdateDocument(ctx context.Context, collection string, id string, update any) (bool, error) { return true, nil }
func (m *MongoDbDriver) DeleteDocument(ctx context.Context, collection string, id string) (bool, error) { return true, nil }
func (m *MongoDbDriver) GetDriverName() string { return "mongodb" }
