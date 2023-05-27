package seekrd

type ApiConfig interface{
	ParsePointer() (error)
 	SaveDB() error
 	LoadDBPointer() error
} // FIXME circular dependencies when using real type

type SeekrdFunc func(ApiConfig) (ApiConfig,error)

type SeekrdService struct{
	Name string `json:"name"`
	Func SeekrdFunc `json:"-"`
	Repeat bool `json:"repeat"`
}

type SeekrdServices []SeekrdService

type SeekrdInstance struct{
	Services SeekrdServices `json:"services"`
	Interval int `json:"interval"` // in minutes
	initialRun bool `default:"true"`
	ApiConfig ApiConfig
}


