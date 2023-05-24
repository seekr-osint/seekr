package api

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/gin-gonic/gin"
	"github.com/seekr-osint/seekr/api/server"
)

type SaveDBFunc func(ApiConfig) error
type LoadDBFunc func(ApiConfig) (ApiConfig, error)
type ApiConfig struct {
	Server         server.Server `json:"server"`
	LogFile        string        `json:"log_file"`
	DataBaseFile   string        `json:"data_base_file"`
	DataBase       DataBase      `json:"data_base"`
	SetCORSHeader  bool          `json:"set_CORS_header"`
	SaveDBFunc     SaveDBFunc    `json:"save_db_func"`
	LoadDBFunc     LoadDBFunc    `json:"load_db_func"`
	GinRouter      *gin.Engine   `json:"gin_router"`
	ApiKeysComplex ApiKeys       `json:"api_keys_complex"`
	ApiKeysSimple  ApiKeysSimple `json:"api_keys"`
	Testing        bool          `json:"testing"`
	BadgerDBLogger badger.Logger `json:"badger_db_logger"`
	Parsers        []func(ApiConfig) (ApiConfig, error)
}
type ApiKeysSimple map[string][]string // map["serviceName"]["key1","key2"]
type ApiKeys struct {
	Github ApiKeyEnum `json:"github"`
}
type ApiKeyEnum map[string]ApiKey
type ApiKey struct {
}
