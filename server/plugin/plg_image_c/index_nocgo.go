//go:build !linux && !freebsd

package plg_image_c

func init() {
	// no-op - this plugin only works on Linux/FreeBSD
}
