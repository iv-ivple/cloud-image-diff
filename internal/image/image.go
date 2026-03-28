package image

import "time"

// ImageMeta is the normalised cloud image representation.
// Every provider translates its API response into this struct.
type ImageMeta struct {
    Provider   string
    Region     string
    ImageID    string
    Name       string
    Release    string // e.g. "24.04", "22.04"
    Arch       string // "amd64", "arm64"
    KernelVer  string
    BuildDate  time.Time
    VirtType   string // "hvm", "pv", "gen2"
    RootDevice string // "ebs", "instance-store", "disk"
}

// ImageFilter is passed to each provider to narrow the query.
type ImageFilter struct {
    Release string
    Arch    string
}
