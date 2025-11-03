package storages

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type FirebaseStorage struct {
	client     *storage.Client
	bucket     *storage.BucketHandle
	bucketName string
}

func NewFirebaseStorage(ctx context.Context) (*FirebaseStorage, error) {
	credsPath := os.Getenv("FIREBASE_CREDENTIALS")
	bucketName := os.Getenv("FIREBASE_STORAGE_BUCKET")

	if credsPath == "" || bucketName == "" {
		return nil, fmt.Errorf("FIREBASE_CREDENTIALS atau FIREBASE_STORAGE_BUCKET belum di-set")
	}

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credsPath))
	if err != nil {
		return nil, fmt.Errorf("gagal membuat storage client: %w", err)
	}

	bucket := client.Bucket(bucketName)
	return &FirebaseStorage{
		client:     client,
		bucket:     bucket,
		bucketName: bucketName,
	}, nil
}

func (fs *FirebaseStorage) UploadFile(ctx context.Context, objectPath string, r io.Reader, contentType string) (string, error) {
	wc := fs.bucket.Object(objectPath).NewWriter(ctx)
	wc.ContentType = contentType

	if _, err := io.Copy(wc, r); err != nil {
		return "", fmt.Errorf("gagal menyalin data ke bucket: %w", err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("gagal menutup writer bucket: %w", err)
	}

	escapedObjectPath := url.PathEscape(objectPath)
	publicURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", fs.bucketName, escapedObjectPath)

	return publicURL, nil
}

func (fs *FirebaseStorage) DeleteFile(ctx context.Context, objectPath string) error {
	o := fs.bucket.Object(objectPath)

	if _, err := o.Attrs(ctx); err != nil {
		return fmt.Errorf("file tidak ditemukan: %w", err)
	}

	if err := o.Delete(ctx); err != nil {
		return fmt.Errorf("gagal menghapus objek: %w", err)
	}
	return nil
}

// CheckBucket adalah fungsi debug untuk memverifikasi koneksi dan izin
func (fs *FirebaseStorage) CheckBucket(ctx context.Context) error {
	// Coba dapatkan atribut bucket.
	// Ini akan gagal jika bucket tidak ada ATAU service account tidak punya izin
	_, err := fs.bucket.Attrs(ctx)
	if err != nil {
		return fmt.Errorf(
			"gagal memverifikasi bucket '%s'. Error: %w. Pastikan bucket ada dan service account memiliki peran 'Storage Object Admin'",
			fs.bucketName,
			err,
		)
	}
	return nil
}
