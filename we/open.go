package we

import (
	"context"
	"encoding/json"
	"io"
	"net/url"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
)

// OpenURLs splits filenames by "," and reads each of the files from the left
// If the definition of the environment variable duplicates in the files, the last one is used.
func OpenURLs(filenames string) (Config, error) {
	envvarsMap := map[string]*EnvVar{}

	for _, filename := range strings.Split(filenames, ",") {
		config, err := OpenURL(filename)
		if err != nil {
			return nil, errors.Errorf("Failed to read %s: %w", filename, err)
		}
		envvars, err := config.EnvVars()
		if err != nil {
			return nil, errors.Errorf("Failed to get envvar from %s: %w", filename, err)
		}

		for _, e := range envvars {
			envvarsMap[e.Name] = e
		}
	}

	ret := &configStruct{}
	for _, e := range envvarsMap {
		ret.envvars = append(ret.envvars, e)
	}

	return ret, nil
}

// OpenURL opens and parses the given file path and returns Config object
//
// OpenURL supports the following path
// - Local file relative|absolute path which doesn't have URL Scheme
// - Google Cloud Storage starting with gs://
func OpenURL(filename string) (Config, error) {
	ctx := context.Background()
	u, err := url.Parse(filename)
	if err != nil {
		return nil, errors.Errorf("Failed to parse %s: %w", filename, err)
	}
	var b []byte
	switch u.Scheme {
	case "gs":
		b, err = openGs(ctx, u)
		if err != nil {
			return nil, errors.Errorf("Failed to open GCS file %s: %w", filename, err)
		}

	case "":
		b, err = openLocalFile(u)
		if err != nil {
			return nil, errors.Errorf("Failed to open local file %s: %w", filename, err)
		}

	case "s3":
		return nil, errors.Errorf("s3 was not supported")
	}

	config := ConfigFile{}
	if err := json.Unmarshal(b, &config); err != nil {
		return nil, errors.Errorf("Failed to unmarshal config file: %s: %w", filename, err)
	}

	return &config, nil
}

func openGs(ctx context.Context, u *url.URL) ([]byte, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, errors.Errorf("Failed to create storage client: %w", err)
	}

	bucket := u.Host
	object := strings.TrimPrefix(u.Path, "/")

	rdr, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, errors.Errorf("Failed to create reader bucket(%s) object(%s): %w", bucket, object, err)
	}
	defer rdr.Close()

	b, err := io.ReadAll(rdr)
	if err != nil {
		return nil, errors.Errorf("Failed to read bucket(%s) object(%s): %w", bucket, object, err)
	}
	return b, nil
}

func openLocalFile(u *url.URL) ([]byte, error) {
	b, err := os.ReadFile(u.Path)
	if err != nil {
		return nil, errors.Errorf("Failed to open config file: %w", err)
	}
	return b, nil
}
