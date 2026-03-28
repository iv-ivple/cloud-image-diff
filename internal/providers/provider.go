package providers
 
import (
    "context"
    "github.com/iv-ivple/cloud-image-diff/internal/image"
)
 
// Provider is the interface every cloud implementation must satisfy.
// Adding a new cloud = implementing these two methods only.
type Provider interface {
    Name() string
    ListImages(ctx context.Context,
        filter image.ImageFilter) ([]image.ImageMeta, error)
}
