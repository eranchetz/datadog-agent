// +build !zlib,!zstd

package compression

// ContentEncoding describes the HTTP header value associated with the compression method
// empty here since there's no compression
// var instead of const to ease testing
var ContentEncoding = ""

// Compress will not compress anything
func Compress(dst []byte, src []byte) ([]byte, error) {
	dst = src
	return dst, nil
}

// Decompress will not decompress anything
func Decompress(dst []byte, src []byte) ([]byte, error) {
	dst = src
	return dst, nil
}
