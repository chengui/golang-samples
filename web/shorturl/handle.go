package shorturl

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	MASK  = 0x3FFFFFFF
	INDEX = 0x0000003D
)

var (
	alnum = []byte("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type Handler struct {
	cache *Cache
}

func NewHandler() *Handler {
	return &Handler{
		cache: NewCache(),
	}
}

func (h *Handler) Shorten(c *gin.Context) {
	long_url := c.Query("url")
	codes := h.transform(long_url)
	if len(codes) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fail Code Transform",
		})
	} else {
		h.cache.Set(codes[0], long_url)
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"long_url":  long_url,
			"short_url": h.urlJoin(c.Request, codes[0]),
		})
	}
}

func (h *Handler) Expand(c *gin.Context) {
	code := c.Param("code")
	if long_url, ok := h.cache.Get(code); !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Page Not Found",
		})
	} else {
		c.Redirect(http.StatusMovedPermanently, long_url)
	}
}

func (h *Handler) transform(long_url string) (codes []string) {
	md5sum := md5.Sum([]byte(long_url))
	hexstr := hex.EncodeToString(md5sum[:])

	for i := 0; i < 32; i += 8 {
		subval, err := strconv.ParseInt(hexstr[i:i+8], 16, 64)
		if err != nil {
			return
		}
		maskval := subval & MASK
		code := []byte{}
		for j := 0; j < 6; j++ {
			index := maskval & INDEX
			code = append(code, alnum[index])
			maskval >>= 5
		}
		codes = append(codes, string(code))
	}
	return
}

func (h *Handler) urlJoin(req *http.Request, code string) (short_url string) {
	var scheme string
	if req.TLS != nil {
		scheme = "https://"
	} else {
		scheme = "http://"
	}
	short := fmt.Sprintf("/s/%s", code)
	return strings.Join([]string{scheme, req.Host, short}, "")
}
