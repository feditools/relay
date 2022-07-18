package http

type (
	Mime   string
	Suffix string
)

const (
	// MimeAll matches any mime type.
	MimeAll Mime = `*/*`

	// MimeAppJRDJSON represents a JSON Resource Descriptor type.
	MimeAppJRDJSON Mime = `application/jrd+json`
	// MimeAppJSON represents a JavaScript Object Notation type.
	MimeAppJSON Mime = `application/json`
	// MimeAppActivityJSON represents a JSON activity pub action type.
	MimeAppActivityJSON Mime = `application/activity+json`
	// MimeAppActivityLDJSON represents JSON-based Linked Data for activity streams type.
	MimeAppActivityLDJSON Mime = `application/ld+json; profile="https://www.w3.org/ns/activitystreams"`
	// MimeImageGIF represents a gif image type.
	MimeImageGIF Mime = `image/gif`
	// MimeImageJPG represents a jpg image type.
	MimeImageJPG Mime = `image/jpeg`
	// MimeImagePNG represents a png image type.
	MimeImagePNG Mime = `image/png`
	// MimeImageSVG represents a svg image type.
	MimeImageSVG Mime = `image/svg+xml`
	// MimeImageWebP represents a webp image type.
	MimeImageWebP Mime = `image/webp`
	// MimeTextHTML represents a html type.
	MimeTextHTML Mime = `text/html`

	// SuffixImageGIF represents a gif image suffix.
	SuffixImageGIF Suffix = `gif`
	// SuffixImageJPG represents a jpg image suffix.
	SuffixImageJPG Suffix = `jpg`
	// SuffixImagePNG represents a png image suffix.
	SuffixImagePNG Suffix = `png`
	// SuffixImageSVG represents a svg image suffix.
	SuffixImageSVG Suffix = `svg`
	// SuffixImageWebP represents a webp image suffix.
	SuffixImageWebP Suffix = `webp`
	// SuffixTextHTML represents a html suffix.
	SuffixTextHTML Suffix = `html`
)

var (
	suffixToMime = map[Suffix]Mime{
		SuffixImageGIF:  MimeImageGIF,
		SuffixImageJPG:  MimeImageJPG,
		SuffixImagePNG:  MimeImagePNG,
		SuffixImageSVG:  MimeImageSVG,
		SuffixImageWebP: MimeImageWebP,
		SuffixTextHTML:  MimeTextHTML,
	}

	mimeToSuffix = map[Mime]Suffix{
		MimeImageGIF:  SuffixImageGIF,
		MimeImageJPG:  SuffixImageJPG,
		MimeImagePNG:  SuffixImagePNG,
		MimeImageSVG:  SuffixImageSVG,
		MimeImageWebP: SuffixImageWebP,
		MimeTextHTML:  SuffixTextHTML,
	}
)

func ToMime(s Suffix) Mime {
	m, ok := suffixToMime[s]
	if ok {
		return m
	}

	return ""
}

func ToSuffix(m Mime) Suffix {
	s, ok := mimeToSuffix[m]
	if ok {
		return s
	}

	return ""
}
