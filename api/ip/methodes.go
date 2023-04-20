package ip

import "github.com/seekr-osint/seekr/api/functions"

func (ips Ips) Parse() (Ips, error) {
	newIps, err := functions.FullParseMapRet(ips, "Ip")
	return newIps, err
}

func (ip Ip) Parse() (Ip, error) {
	return ip, nil
}

func (ip Ip) Markdown() (string, error) {
	return functions.Markdown(ip)
}

func (ips Ips) Markdown() (string, error) {
	return functions.MarkdownMap(ips)
}
