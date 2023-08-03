//go:build windows && 386

package sherpa_ncnn

// #cgo LDFLAGS: -L ${SRCDIR}/lib/i686-pc-windows-gnu -lsherpa-ncnn-c-api -lsherpa-ncnn-core -lkaldi-native-fbank-core -lncnn
import "C"
