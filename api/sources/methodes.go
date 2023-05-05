package sources

import "github.com/seekr-osint/seekr/api/functions"

func (sources Sources) Parse() (Sources, error) {
	return functions.FullParseMapRet(sources, "url")
}

func (source Source) Parse() (Source, error) {
	return source, nil
}

func (source Source) Markdown() (string, error) {
	return functions.Markdown(source)
}
