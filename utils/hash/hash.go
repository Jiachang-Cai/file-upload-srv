package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func HashByteMd5(file []byte) string {
	hash := md5.New()
	hash.Write(file)
	return hex.EncodeToString(hash.Sum(nil))
}
