package s3

// todo 应该废弃这个 用配置文件
type Bucket string

func (b Bucket) Domain() string {
	return BucketDomain(b)
}

func (b Bucket) Uri(name string) string {
	return b.Domain() + "/" + name
}

const (
	Image Bucket = "guapitu-images"
	Music Bucket = "guapitu-master"
)

const (
	ImageDomain = "https://img.ohdat.io"
	MusicDomain = "https://master.ohdat.io"
)

func BucketDomain(b Bucket) string {
	switch b {
	case Image:
		return ImageDomain
	case Music:
		return MusicDomain
	}
	return ""
}

//var (
//	BucketMapDomain = map[string]string{
//		"ohdat-images": "https://images.megauniverse.io",
//		"ohdat-master": "https://master.ohdat.io",
//	}
//)
