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
	{
		Name:   "YouTube",
		Domain: "youtube.com",
		URLTemplates: map[string]URLTemplate{
			"HtmlURL": "{{ .Domain }}/@{{ .Username }}",
			"bio":     "{{.Domain}}/@{{.Username}}/about",
		},
		UserExistsCheck: `{{ status .URLs.HtmlURL 200 -1 }}`,
		TestData: TestData[DefaultAccount]{
			MockData: []MockData[DefaultAccount]{
				{
					User: User{
						Username: "mrbeast",
					},
					Result: ScanResult[DefaultAccount]{
						Account: &DefaultAccount{
							Name: "YouTube",
							URL:  "https://youtube.com/@mrbeast",
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
							Name: "YouTube",
						},
						Exists: false,
					},
				},
			},
		},
	},
	{
		Name:   "TikTok",
		Domain: "tiktok.com",
		URLTemplates: map[string]URLTemplate{
			"HtmlURL": "{{ .Domain }}/@{{ .Username }}",
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
							Name: "TikTok",
							URL:  "https://tiktok.com/@greg",
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
							Name: "TikTok",
						},
						Exists: false,
					},
				},
			},
		},
	},
	{
		Name:   "Npm",
		Domain: "npmjs.com",
		URLTemplates: map[string]URLTemplate{
			"HtmlURL": "{{ .Domain }}/~{{ .Username }}",
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
							Name: "Npm",
							URL:  "https://npmjs.com/~greg",
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
							Name: "Npm",
						},
						Exists: false,
					},
				},
			},
		},
	},
	{
		Name:   "chess.com",
		Domain: "chess.com",
		URLTemplates: map[string]URLTemplate{
			"HtmlURL": "{{ .Domain }}/member/{{ .Username }}",
			"ApiURL":  "api.{{ .Domain }}/pub/player/{{ .Username }}",
		},
		UserExistsCheck: `{{ status .URLs.ApiURL 200 -1 }}`,
		TestData: TestData[DefaultAccount]{
			MockData: []MockData[DefaultAccount]{
				{
					User: User{
						Username: "danielnaroditsky",
					},
					Result: ScanResult[DefaultAccount]{
						Account: &DefaultAccount{
							Name: "chess.com",
							URL:  "https://chess.com/member/danielnaroditsky",
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
							Name: "chess.com",
						},
						Exists: false,
					},
				},
			},
		},
	},

	{
		Name:   "Replit",
		Domain: "replit.com",
		URLTemplates: map[string]URLTemplate{
			"HtmlURL": "{{ .Domain }}/@{{ .Username }}",
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
							Name: "Replit",
							URL:  "https://replit.com/@greg",
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
							Name: "Replit",
						},
						Exists: false,
					},
				},
			},
		},
	},
	{
		Name:   "Asciinema",
		Domain: "asciinema.org",
		URLTemplates: map[string]URLTemplate{
			"HtmlURL": "{{ .Domain }}/~{{ .Username }}",
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
							Name: "Asciinema",
							URL:  "https://asciinema.org/~greg",
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
							Name: "Asciinema",
						},
						Exists: false,
					},
				},
			},
		},
	},
	{
		Name:   "Vimeo",
		Domain: "vimeo.com",
		URLTemplates: map[string]URLTemplate{
			"HtmlURL": "{{ .Domain }}/{{ .Username }}/",
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
							Name: "Vimeo",
							URL:  "https://vimeo.com/greg/",
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
							Name: "Vimeo",
						},
						Exists: false,
					},
				},
			},
		},
	},

	// {
	// 	Name:   "Twitch",
	// 	Domain: "twitch.tv",
	// 	URLTemplates: map[string]URLTemplate{
	// 		"HtmlURL": "{{ .Domain }}/{{ .Username }}/",
	// 	},
	// 	UserExistsCheck: `{{ status .URLs.HtmlURL 200 -1 }}`,
	// 	TestData: TestData[DefaultAccount]{
	// 		MockData: []MockData[DefaultAccount]{
	// 			{
	// 				User: User{
	// 					Username: "greg",
	// 				},
	// 				Result: ScanResult[DefaultAccount]{
	// 					Account: &DefaultAccount{
	// 						Name: "Twitch",
	// 						URL:  "https://m.twitch.tv/greg/",
	// 					},
	// 					Exists: true,
	// 				},
	// 			},
	// 			{
	// 				User: User{
	// 					Username: "thatoneguywholikesseekr",
	// 				},
	// 				Result: ScanResult[DefaultAccount]{
	// 					Account: &DefaultAccount{
	// 						Name: "Twitch",
	// 					},
	// 					Exists: false,
	// 				},
	// 			},
	// 		},
	// 	},
	// },
	{
		Name:   "PyPi",
		Domain: "pypi.org",
		URLTemplates: map[string]URLTemplate{
			"HtmlURL": "{{ .Domain }}/user/{{ .Username }}/",
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
							Name: "PyPi",
							URL:  "https://pypi.org/user/greg/",
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
							Name: "PyPi",
						},
						Exists: false,
					},
				},
			},
		},
	},
	{
		Name:   "Instagram",
		Domain: "instagram.com",
		URLTemplates: map[string]URLTemplate{
			"HtmlURL": "{{ .Domain }}/{{ .Username }}/",
		},
		UserExistsCheck: `{{ bodyPatternMatch .URLs.HtmlURL "user?username=" 1 }}`,
		TestData: TestData[DefaultAccount]{
			MockData: []MockData[DefaultAccount]{
				{
					User: User{
						Username: "greg",
					},
					Result: ScanResult[DefaultAccount]{
						Account: &DefaultAccount{
							Name: "Instagram",
							URL:  "https://instagram.com/greg/",
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
							Name: "Instagram",
						},
						Exists: false,
					},
				},
			},
		},
	},
}

func AddServices(a Services) func(*Services) {
	return func(s *Services) {
		*s = append(*s, a...)
	}
}
func AddService(a ...AccountScanner) func(*Services) {
	return func(s *Services) {
		*s = append(*s, a...)
	}
}

func DefaultServices(options ...func(*Services)) *Services {
	services := &defaultServices
	for _, o := range options {
		o(services)
	}
	return services
}
