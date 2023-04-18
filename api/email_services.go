package api

var DefaultMailServices = MailServices{
	MailService{
		Name:           "Discord",
		UserExistsFunc: DiscordMail,
		Icon:           "./images/mail/discord.png",
	},
	MailService{
		Name:           "Spotify",
		UserExistsFunc: SpotifyMail,
		Icon:           "./images/mail/spotify.png",
	},
	MailService{
		Name:           "Twitter",
		UserExistsFunc: TwitterMail,
		Icon:           "./images/mail/twitter.png",
	},
	MailService{
		Name:           "GitHub",
		UserExistsFunc: GitHubEmail,
		Icon:           "./images/mail/github.svg",
	},
	MailService{
		Name:           "Ubuntu GPG",
		UserExistsFunc: UbuntuGPGUserExists,
		Icon:           "./images/mail/ubuntu.png",
		Url:            "https://keyserver.ubuntu.com/pks/lookup?search={{ email }}&op=index",
	},
	MailService{
		Name:           "keys.gnupg.net",
		UserExistsFunc: KeysGnuPGUserExists,
		Icon:           "./images/mail/gnupg.ico",
		Url:            "https://keys.gnupg.net/pks/lookup?search={{ email }}&op=index",
	},
}

// Makes trouble FIXME #234
//MailService{
//	Name:           "keyserver.pgp.com",
//	UserExistsFunc: KeyserverPGPUserExists,
//	Icon:           "https://pgp.com/favicon.ico",
//	Url:            "https://keyserver.pgp.com/pks/lookup?search={{ email }}&op=index",
//},
//MailService{ // FIXME
//    Name: "pgp.mit.edu",
//    UserExistsFunc: PgpMitUserExists,
//    Icon: "https://pgp.mit.edu/favicon.ico",
//},

// MailService{ // FIXME
//
//	   Name: "pool.sks-keyservers.net",
//	   UserExistsFunc: PoolSKSUserExists,
//	   Icon: "https://sks-keyservers.net/favicon.ico",
//	},
