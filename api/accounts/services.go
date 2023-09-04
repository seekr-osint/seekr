package accounts

var defaultServices Services = Services{
	{
		Name:   "GitHub",
		Domain: "github.com",
		URLTemplates: map[string]URLTemplate{
			"HtmlURL": "{{ .Domain }}/{{ .Username }}",
		},
		UserExistsCheck: `{{ status .URLs.HtmlURL 200 -1 }}`,
		TestData: TestData[DefaultAccount]{
			MockData: []MockData[DefaultAccount]{
				{
					User: User{
						Username: "greg",
					},
					Result: ScanResult[DefaultAccount]{
						Account: &DefaultAccount{
							Name: "GitHub",
							URL:  "https://github.com/greg",
						},
						Exists: true,
					},
				},
				{
					User: User{
						Username: "gregdoesnotexsist",
					},
					Result: ScanResult[DefaultAccount]{
						Account: &DefaultAccount{
							Name: "GitHub",
							URL:  "https://github.com/gregdoesnotexsist",
						},
						Exists: false,
					},
				},
			},
		},
	},
}

func DefaultServices() Services {
	return defaultServices
}
