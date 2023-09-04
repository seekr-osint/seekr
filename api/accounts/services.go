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
						},
						Exists: false,
					},
				},
			},
		},
	},
	{
		Name:   "Lichess",
		Domain: "lichess.org",
		URLTemplates: map[string]URLTemplate{
			"HtmlURL": "{{ .Domain }}/@/{{ .Username }}",
			"ApiURL":  "{{ .Domain }}/api/user/{{ .Username }}",
		},
		UserExistsCheck: `{{ status .URLs.ApiURL 200 -1 }}`,
		TestData: TestData[DefaultAccount]{
			MockData: []MockData[DefaultAccount]{
				{
					User: User{
						Username: "hansniemman", // FIXME recognize disabled accoubts
					},
					Result: ScanResult[DefaultAccount]{
						Account: &DefaultAccount{
							Name: "Lichess",
							URL:  "https://lichess.org/@/hansniemman",
						},
						Exists: true,
					},
				},
				{
					User: User{
						Username: "starwars",
					},
					Result: ScanResult[DefaultAccount]{
						Account: &DefaultAccount{
							Name: "Lichess",
							URL:  "https://lichess.org/@/starwars",
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
							Name: "Lichess",
						},
						Exists: false,
					},
				},
			},
		},
	},
	{
		Name:   "Snapchat",
		Domain: "snapchat.com",
		URLTemplates: map[string]URLTemplate{
			"HtmlURL": "{{ .Domain }}/add/{{ .Username }}",
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
							Name: "Snapchat",
							URL:  "https://snapchat.com/add/greg",
						},
						Exists: true,
					},
				},
				{
					User: User{
						Username: "thatoneguywholikesseekr",
					},
					Result: ScanResult[DefaultAccount]{
						Account: &DefaultAccount{
							Name: "Snapchat",
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
