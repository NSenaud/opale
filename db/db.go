package db

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/shibukawa/configdir"
)

const dbname = "opale.db"

var log = logrus.New()

func GetDbPath() (path string) {
	cache := GetCachePath()
	path = cache.Path + "/opale.db"
	return
}

func GetCachePath() (cache *configdir.Config) {
	configDirs := configdir.New("TechnicalDebtFactory", "opale")
	cache = configDirs.QueryCacheFolder()
	log.WithFields(logrus.Fields{
		"path": cache.Path,
	}).Debug("Cache directory.")

	return
}
