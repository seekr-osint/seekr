package info

type Info struct {
	Version string `json:"version"`
	IsLatest bool `json:"is_latest"`
	Latest string `json:"latest"`
	DownloadUrl string `json:"download_url"`
}
