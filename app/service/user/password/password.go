/**
* @Author: lik
* @Date: 2021/3/7 19:15
* @Version 1.0
 */
package password

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"gopkg.in/hlandau/passlib.v1/abstract"
	"gopkg.in/hlandau/passlib.v1/hash/pbkdf2/raw"
	"hash"
	"strings"
)

//md5验证
func Md5Str(src string) string {

	h := md5.New()
	h.Write([]byte(src))                               // 需要加密的字符串为
	fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	return hex.EncodeToString(h.Sum(nil))
}

//base编码
func Base64EncodeStr(src string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(src)))
}

//base解码
func Base64DecodeStr(src string) string {
	a, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "error"
	}
	return string(a)
}

// PBKDF2-SHA1,PBKDF2-SHA256,PBKDF-SHA512

var SHA1Encrypt abstract.Scheme
var SHA256Encrypt abstract.Scheme
var SHA512Encrypt abstract.Scheme

const (
	RecommendedRoundsSHA1   = 131000
	RecommendedRoundsSHA256 = 29000
	RecommendedRoundsSHA512 = 25000
)

const SaltLength = 16

func init() {
	SHA1Encrypt = New("$pbkdf2$", sha1.New, RecommendedRoundsSHA1)
	SHA256Encrypt = New("$pbkdf2-sha256$", sha256.New, RecommendedRoundsSHA256)
	SHA512Encrypt = New("$pbkdf2-sha512$", sha512.New, RecommendedRoundsSHA512)
}

type scheme struct {
	Ident    string
	HashFunc func() hash.Hash
	Rounds   int
}

func New(ident string, hf func() hash.Hash, rounds int) abstract.Scheme {
	return &scheme{
		Ident:    ident,
		HashFunc: hf,
		Rounds:   rounds,
	}
}

func (s *scheme) Hash(password string) (string, error) {
	salt := make([]byte, SaltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	h := raw.Hash([]byte(password), salt, s.Rounds, s.HashFunc)

	newHash := fmt.Sprintf("%s%d$%s$%s", s.Ident, s.Rounds, raw.Base64Encode(salt), h)
	return newHash, nil
}

func (s *scheme) Verify(password, stub string) (err error) {
	_, rounds, salt, oldHash, err := raw.Parse(stub)
	if err != nil {
		return
	}

	newHash := raw.Hash([]byte(password), salt, rounds, s.HashFunc)

	if len(newHash) == 0 || !abstract.SecureCompare(oldHash, newHash) {
		err = abstract.ErrInvalidPassword
	}

	return
}

func (s *scheme) SupportsStub(stub string) bool {
	return strings.HasPrefix(stub, s.Ident)
}

func (s *scheme) NeedsUpdate(stub string) bool {
	_, rounds, salt, _, err := raw.Parse(stub)
	return err == raw.ErrInvalidRounds || rounds < s.Rounds || len(salt) < SaltLength
}
