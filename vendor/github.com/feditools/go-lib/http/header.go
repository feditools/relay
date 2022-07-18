package http

const (
	// CacheControlMustRevalidate that the response can be stored in caches and can be reused while fresh.
	CacheControlMustRevalidate = "must-revalidate"
	// CacheControlNoCache indicates the response must be validated with the origin server before each reuse.
	CacheControlNoCache = "no-cache"
	// CacheControlNoStore indicates that all caches should not store this response.
	CacheControlNoStore = "no-store"
	// CacheControlNoTransform indicates that any intermediary shouldn't transform the response contents.
	CacheControlNoTransform = "no-transform"

	// HeaderAccept is the key for the accept header.
	HeaderAccept = "Accept"
	// HeaderCacheControl is the key for the cache control header.
	HeaderCacheControl = "Cache-Control"
	// HeaderContentType is the key for the content type header.
	HeaderContentType = "Content-Type"
)
