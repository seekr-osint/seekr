package accounts

var DefaultServices Services = Services{
	{
		Name:   "GitHub",
		Domain: "github.com",
		URLTemplates: map[string]URLTemplate{
			"HtmlURL": "{{ .Domain }}/{{ .Username }}",
		},
		UserExistsCheck: `{{ status .URLs.HtmlURL 200 -1 }}`,
	},
}
