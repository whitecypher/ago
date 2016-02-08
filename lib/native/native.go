package native

var (
	pkgMap = map[string]bool{}
	pkgs   = []string{
		"archive",
		"archive/tar",
		"archive/tar/testdata",
		"archive/zip",
		"archive/zip/testdata",
		"bufio",
		"builtin",
		"bytes",
		"cmd",
		"cmd/addr2line",
		"cmd/api",
		"cmd/api/testdata",
		"cmd/api/testdata/src",
		"cmd/api/testdata/src/pkg",
		"cmd/api/testdata/src/pkg/p1",
		"cmd/api/testdata/src/pkg/p2",
		"cmd/api/testdata/src/pkg/p3",
		"cmd/asm",
		"cmd/asm/internal",
		"cmd/asm/internal/arch",
		"cmd/asm/internal/asm",
		"cmd/asm/internal/asm/testdata",
		"cmd/asm/internal/flags",
		"cmd/asm/internal/lex",
		"cmd/cgo",
		"cmd/compile",
		"cmd/compile/internal",
		"cmd/compile/internal/amd64",
		"cmd/compile/internal/arm",
		"cmd/compile/internal/arm64",
		"cmd/compile/internal/big",
		"cmd/compile/internal/gc",
		"cmd/compile/internal/gc/builtin",
		"cmd/compile/internal/ppc64",
		"cmd/compile/internal/x86",
		"cmd/cover",
		"cmd/cover/testdata",
		"cmd/dist",
		"cmd/doc",
		"cmd/doc/testdata",
		"cmd/fix",
		"cmd/go",
		"cmd/go/testdata",
		"cmd/go/testdata/cgocover",
		"cmd/go/testdata/failssh",
		"cmd/go/testdata/generate",
		"cmd/go/testdata/importcom",
		"cmd/go/testdata/importcom/src",
		"cmd/go/testdata/importcom/src/bad",
		"cmd/go/testdata/importcom/src/conflict",
		"cmd/go/testdata/importcom/src/works",
		"cmd/go/testdata/importcom/src/works/x",
		"cmd/go/testdata/importcom/src/wrongplace",
		"cmd/go/testdata/local",
		"cmd/go/testdata/local/easysub",
		"cmd/go/testdata/local/sub",
		"cmd/go/testdata/local/sub/sub",
		"cmd/go/testdata/norunexample",
		"cmd/go/testdata/rundir",
		"cmd/go/testdata/rundir/sub",
		"cmd/go/testdata/shadow",
		"cmd/go/testdata/shadow/root1",
		"cmd/go/testdata/shadow/root1/src",
		"cmd/go/testdata/shadow/root1/src/foo",
		"cmd/go/testdata/shadow/root1/src/math",
		"cmd/go/testdata/shadow/root2",
		"cmd/go/testdata/shadow/root2/src",
		"cmd/go/testdata/shadow/root2/src/foo",
		"cmd/go/testdata/src",
		"cmd/go/testdata/src/badc",
		"cmd/go/testdata/src/badpkg",
		"cmd/go/testdata/src/badtest",
		"cmd/go/testdata/src/badtest/badexec",
		"cmd/go/testdata/src/badtest/badsyntax",
		"cmd/go/testdata/src/badtest/badvar",
		"cmd/go/testdata/src/cgotest",
		"cmd/go/testdata/src/go-cmd-test",
		"cmd/go/testdata/src/main_test",
		"cmd/go/testdata/src/notest",
		"cmd/go/testdata/src/syntaxerror",
		"cmd/go/testdata/src/testcycle",
		"cmd/go/testdata/src/testcycle/p1",
		"cmd/go/testdata/src/testcycle/p2",
		"cmd/go/testdata/src/testcycle/p3",
		"cmd/go/testdata/src/testcycle/q1",
		"cmd/go/testdata/src/testdep",
		"cmd/go/testdata/src/testdep/p1",
		"cmd/go/testdata/src/testdep/p2",
		"cmd/go/testdata/src/testdep/p3",
		"cmd/go/testdata/src/vend",
		"cmd/go/testdata/src/vend/hello",
		"cmd/go/testdata/src/vend/subdir",
		"cmd/go/testdata/src/vend/vendor",
		"cmd/go/testdata/src/vend/vendor/p",
		"cmd/go/testdata/src/vend/vendor/q",
		"cmd/go/testdata/src/vend/vendor/strings",
		"cmd/go/testdata/src/vend/x",
		"cmd/go/testdata/src/vend/x/invalid",
		"cmd/go/testdata/src/vend/x/vendor",
		"cmd/go/testdata/src/vend/x/vendor/p",
		"cmd/go/testdata/src/vend/x/vendor/p/p",
		"cmd/go/testdata/src/vend/x/vendor/r",
		"cmd/go/testdata/src/vetpkg",
		"cmd/go/testdata/src/xtestonly",
		"cmd/go/testdata/testimport",
		"cmd/go/testdata/testimport/p1",
		"cmd/go/testdata/testimport/p2",
		"cmd/go/testdata/testinternal",
		"cmd/go/testdata/testinternal2",
		"cmd/go/testdata/testinternal2/x",
		"cmd/go/testdata/testinternal2/x/y",
		"cmd/go/testdata/testinternal2/x/y/z",
		"cmd/go/testdata/testinternal2/x/y/z/internal",
		"cmd/go/testdata/testinternal2/x/y/z/internal/w",
		"cmd/go/testdata/testinternal3",
		"cmd/go/testdata/testinternal4",
		"cmd/go/testdata/testinternal4/src",
		"cmd/go/testdata/testinternal4/src/p",
		"cmd/go/testdata/testinternal4/src/q",
		"cmd/go/testdata/testinternal4/src/q/internal",
		"cmd/go/testdata/testinternal4/src/q/internal/x",
		"cmd/go/testdata/testinternal4/src/q/j",
		"cmd/go/testdata/testonly",
		"cmd/go/testdata/testvendor",
		"cmd/go/testdata/testvendor/src",
		"cmd/go/testdata/testvendor/src/p",
		"cmd/go/testdata/testvendor/src/q",
		"cmd/go/testdata/testvendor/src/q/vendor",
		"cmd/go/testdata/testvendor/src/q/vendor/x",
		"cmd/go/testdata/testvendor/src/q/y",
		"cmd/go/testdata/testvendor/src/q/z",
		"cmd/go/testdata/testvendor2",
		"cmd/go/testdata/testvendor2/src",
		"cmd/go/testdata/testvendor2/src/p",
		"cmd/go/testdata/testvendor2/vendor",
		"cmd/go/testdata/testvendor2/vendor/x",
		"cmd/gofmt",
		"cmd/gofmt/testdata",
		"cmd/internal",
		"cmd/internal/gcprog",
		"cmd/internal/goobj",
		"cmd/internal/obj",
		"cmd/internal/obj/arm",
		"cmd/internal/obj/arm64",
		"cmd/internal/obj/ppc64",
		"cmd/internal/obj/x86",
		"cmd/internal/objfile",
		"cmd/internal/rsc.io",
		"cmd/internal/rsc.io/arm",
		"cmd/internal/rsc.io/arm/armasm",
		"cmd/internal/rsc.io/arm/armasm/testdata",
		"cmd/internal/rsc.io/x86",
		"cmd/internal/rsc.io/x86/x86asm",
		"cmd/internal/rsc.io/x86/x86asm/testdata",
		"cmd/link",
		"cmd/link/internal",
		"cmd/link/internal/amd64",
		"cmd/link/internal/arm",
		"cmd/link/internal/arm64",
		"cmd/link/internal/ld",
		"cmd/link/internal/ppc64",
		"cmd/link/internal/x86",
		"cmd/nm",
		"cmd/objdump",
		"cmd/objdump/testdata",
		"cmd/pack",
		"cmd/pprof",
		"cmd/pprof/internal",
		"cmd/pprof/internal/commands",
		"cmd/pprof/internal/driver",
		"cmd/pprof/internal/fetch",
		"cmd/pprof/internal/plugin",
		"cmd/pprof/internal/profile",
		"cmd/pprof/internal/report",
		"cmd/pprof/internal/svg",
		"cmd/pprof/internal/symbolizer",
		"cmd/pprof/internal/symbolz",
		"cmd/pprof/internal/tempfile",
		"cmd/trace",
		"cmd/vet",
		"cmd/vet/testdata",
		"cmd/vet/testdata/tagtest",
		"cmd/vet/whitelist",
		"cmd/yacc",
		"cmd/yacc/testdata",
		"cmd/yacc/testdata/expr",
		"compress",
		"compress/bzip2",
		"compress/bzip2/testdata",
		"compress/flate",
		"compress/gzip",
		"compress/gzip/testdata",
		"compress/lzw",
		"compress/testdata",
		"compress/zlib",
		"container",
		"container/heap",
		"container/list",
		"container/ring",
		"crypto",
		"crypto/aes",
		"crypto/cipher",
		"crypto/des",
		"crypto/dsa",
		"crypto/ecdsa",
		"crypto/ecdsa/testdata",
		"crypto/elliptic",
		"crypto/hmac",
		"crypto/md5",
		"crypto/rand",
		"crypto/rc4",
		"crypto/rsa",
		"crypto/rsa/testdata",
		"crypto/sha1",
		"crypto/sha256",
		"crypto/sha512",
		"crypto/subtle",
		"crypto/tls",
		"crypto/tls/testdata",
		"crypto/x509",
		"crypto/x509/pkix",
		"database",
		"database/sql",
		"database/sql/driver",
		"debug",
		"debug/dwarf",
		"debug/dwarf/testdata",
		"debug/elf",
		"debug/elf/testdata",
		"debug/gosym",
		"debug/macho",
		"debug/macho/testdata",
		"debug/pe",
		"debug/pe/testdata",
		"debug/plan9obj",
		"debug/plan9obj/testdata",
		"encoding",
		"encoding/ascii85",
		"encoding/asn1",
		"encoding/base32",
		"encoding/base64",
		"encoding/binary",
		"encoding/csv",
		"encoding/gob",
		"encoding/hex",
		"encoding/json",
		"encoding/json/testdata",
		"encoding/pem",
		"encoding/xml",
		"errors",
		"expvar",
		"flag",
		"fmt",
		"go",
		"go/ast",
		"go/build",
		"go/build/testdata",
		"go/build/testdata/empty",
		"go/build/testdata/multi",
		"go/build/testdata/other",
		"go/build/testdata/other/file",
		"go/constant",
		"go/doc",
		"go/doc/testdata",
		"go/format",
		"go/importer",
		"go/internal",
		"go/internal/gccgoimporter",
		"go/internal/gccgoimporter/testdata",
		"go/internal/gcimporter",
		"go/internal/gcimporter/testdata",
		"go/parser",
		"go/parser/testdata",
		"go/printer",
		"go/printer/testdata",
		"go/scanner",
		"go/token",
		"go/types",
		"go/types/testdata",
		"hash",
		"hash/adler32",
		"hash/crc32",
		"hash/crc64",
		"hash/fnv",
		"html",
		"html/template",
		"image",
		"image/color",
		"image/color/palette",
		"image/draw",
		"image/gif",
		"image/internal",
		"image/internal/imageutil",
		"image/jpeg",
		"image/png",
		"image/png/testdata",
		"image/png/testdata/pngsuite",
		"image/testdata",
		"index",
		"index/suffixarray",
		"internal",
		"internal/format",
		"internal/singleflight",
		"internal/syscall",
		"internal/syscall/unix",
		"internal/syscall/windows",
		"internal/syscall/windows/registry",
		"internal/testenv",
		"internal/trace",
		"io",
		"io/ioutil",
		"log",
		"log/syslog",
		"math",
		"math/big",
		"math/cmplx",
		"math/rand",
		"mime",
		"mime/multipart",
		"mime/multipart/testdata",
		"mime/quotedprintable",
		"mime/testdata",
		"net",
		"net/http",
		"net/http/cgi",
		"net/http/cgi/testdata",
		"net/http/cookiejar",
		"net/http/fcgi",
		"net/http/httptest",
		"net/http/httputil",
		"net/http/internal",
		"net/http/pprof",
		"net/http/testdata",
		"net/internal",
		"net/internal/socktest",
		"net/mail",
		"net/rpc",
		"net/rpc/jsonrpc",
		"net/smtp",
		"net/testdata",
		"net/textproto",
		"net/url",
		"os",
		"os/exec",
		"os/signal",
		"os/user",
		"path",
		"path/filepath",
		"reflect",
		"regexp",
		"regexp/syntax",
		"regexp/testdata",
		"runtime",
		"runtime/cgo",
		"runtime/debug",
		"runtime/pprof",
		"runtime/race",
		"runtime/race/testdata",
		"runtime/trace",
		"sort",
		"strconv",
		"strconv/testdata",
		"strings",
		"sync",
		"sync/atomic",
		"syscall",
		"testing",
		"testing/iotest",
		"testing/quick",
		"text",
		"text/scanner",
		"text/tabwriter",
		"text/template",
		"text/template/parse",
		"text/template/testdata",
		"time",
		"unicode",
		"unicode/utf16",
		"unicode/utf8",
		"unsafe",
	}
)

func init() {
	for _, name := range Packages() {
		pkgMap[name] = true
	}
}

// Packages returns an array of native packages
func Packages() []string {
	return pkgs
}

// IsNative returns whether given package name is a native package or not
func IsNative(name string) bool {
	_, ok := pkgMap[name]
	return ok
}
