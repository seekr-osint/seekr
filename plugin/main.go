package plugin

import (
  "plugin"
  api "github.com/seekr-osint/seekr/api"
)


func Load(path string,config api.ApiConfig) api.ApiConfig {
  plug, err := plugin.Open(path)
  api.CheckAndLog(err,"error loading plugin",config)
  run, err := plug.Lookup("Run")
  api.CheckAndLog(err,"error loading Run function",config)
  config = run.(func(api.ApiConfig) api.ApiConfig)(config)
  return config
}
