package blob

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/romankravchuk/pastebin/internal/entity"
	"github.com/romankravchuk/pastebin/internal/usecase"
	"github.com/romankravchuk/pastebin/pkg/minio"
)

const public = "public"

var _ usecase.PastesBlobStorage = &PastesBlobStorage{}

type PastesBlobStorage struct {
	m *minio.Minio
}

func NewPastesBlobStorage(m *minio.Minio) *PastesBlobStorage {
	return &PastesBlobStorage{m: m}
}

// Create uploads paste to minio storage.
//
// Hash is used as object name, user id is used as bucket name.
// If user id is empty, it will be set to public.
func (bs *PastesBlobStorage) Create(ctx context.Context, p *entity.Paste) error {
	bucket := public
	if p.UserID.Valid {
		bucket = p.UserID.String
	}

	reader := bytes.NewReader(p.File)

	err := bs.m.UploadObject(ctx, bucket, p.Hash, p.File.Size(), reader)
	if err != nil {
		return fmt.Errorf("PasteBlobStorage.Create: %w", err)
	}

	return nil
}

// Delete deletes a file from obj storage.
func (bs *PastesBlobStorage) Delete(ctx context.Context, userID, id string) error {
	if userID == "" {
		userID = public
	}

	if err := bs.m.DeleteObject(ctx, userID, id); err != nil {
		return fmt.Errorf("PastesBlobStorage.Delete: %w", err)
	}

	return nil
}

// Get returns a file from obj storage.
func (bs *PastesBlobStorage) Get(ctx context.Context, userID, id string) (entity.File, error) {
	if userID == "" {
		userID = public
	}

	obj, err := bs.m.GetObject(ctx, userID, id)
	if err != nil {
		return nil, fmt.Errorf("PastesBlobStorage.Get: %w", err)
	}

	objInfo, err := obj.Stat()
	if err != nil {
		return nil, fmt.Errorf("PastesBlobStorage.Get: %w", err)
	}

	data := make(entity.File, objInfo.Size)

	_, err = obj.Read(data)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("PastesBlobStorage.Get: %w", err)
	}

	return data, nil
}

// Update updates a file in obj storage.
func (bs *PastesBlobStorage) Update(ctx context.Context, p *entity.Paste) error {
	err := bs.m.UploadObject(ctx, p.UserID.String, p.Hash, p.File.Size(), bytes.NewReader(p.File))
	if err != nil {
		return fmt.Errorf("PastesBlobStorage.Update: %w", err)
	}

	return nil
}
