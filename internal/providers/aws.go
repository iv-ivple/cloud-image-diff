package providers
 
import (
    "context"
    "fmt"
    "github.com/iv-ivple/cloud-image-diff/internal/image"
)
 
type AWSProvider struct{ Regions []string }
 
func NewAWSProvider(regions []string) *AWSProvider {
    return &AWSProvider{Regions: regions}
}
 
func (p *AWSProvider) Name() string { return "aws" }
 
func (p *AWSProvider) ListImages(ctx context.Context,
    filter image.ImageFilter) ([]image.ImageMeta, error) {
    var results []image.ImageMeta
    for _, region := range p.Regions {
        results = append(results, image.ImageMeta{
            Provider:  "aws",
            Region:    region,
            ImageID:   fmt.Sprintf("ami-stub-%s", region),
            Release:   filter.Release,
            Arch:      filter.Arch,
            KernelVer: "6.8.0-1008-aws",
        })
    }
    return results, nil
}
