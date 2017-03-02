package ctemp

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
)

func hashMyself(salt string) (hash string, err error) {
	const BUFSIZE = 1024

	shaHash := sha256.New()
	file, err := os.Open(os.Args[0])
	if err != nil {
		return "", err
	}
	defer file.Close()
	buf := make([]byte, BUFSIZE)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			// Readエラー処理
			break
		}
		shaHash.Write(buf[:n])
	}
	shaHash.Write([]byte(salt))
	return fmt.Sprintf("%x", shaHash.Sum(nil)), nil

}

// ConsistentTempDir creates a new temporary directory
// 自分自身のハッシュを元にした一貫性のあるTempディレクトリを作ります。
// したがって毎回同じディレクトリを作ることになります。
func ConsistentTempDir(dir, prefix, salt string) (name string, err error) {
	if dir == "" {
		dir = os.TempDir()
	}
	hash, err := hashMyself(salt)
	if err != nil {
		return "", err
	}
	try := filepath.Join(dir, prefix, hash)
	os.Mkdir(try, 0700)
	return try, nil
}
