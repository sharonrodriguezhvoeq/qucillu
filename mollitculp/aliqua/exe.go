import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

// buildNumber gets the build number of the GCS bucket.
func buildNumber(w io.Writer, bucketName string) error {
	// bucketName := "bucket-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	attrs, err := client.Bucket(bucketName).Attrs(ctx)
	if err != nil {
		return fmt.Errorf("Bucket(%q).Attrs: %v", bucketName, err)
	}
	fmt.Fprintf(w, "Build number: %v\n", attrs.MetaGeneration)
	return nil
}
  
