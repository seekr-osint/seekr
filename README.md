<p align="center">
  <a href="https://github.com/seekr-osint/seekr" target="blank"><img src="https://user-images.githubusercontent.com/67828948/230488036-00377252-180b-4202-a0f7-392f79b0aafc.png" width="450" alt="Seekr Logo" /></a>
 </p>

 <p align="center">
 A multi-purpose toolkit for gathering and managing OSINT-Data with a neat web-interface.
</p>
<p align="center">
<img alt="GitHub" src="https://img.shields.io/github/license/seekr-osint/seekr">
<img alt="GitHub Workflow Status" src="https://img.shields.io/github/actions/workflow/status/seekr-osint/seekr/go.yml">
<img alt="GitHub repo size" src="https://img.shields.io/github/repo-size/seekr-osint/seekr">
<img alt="GitHub issues" src="https://img.shields.io/github/issues/seekr-osint/seekr">
<img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/seekr-osint/seekr">
</p>

## Introduction
Seekr is a multi-purpose toolkit for gathering and managing OSINT-data with a sleek web interface. The backend is written in Go and offers a wide range of features for data collection, organization, and analysis. Whether you're a researcher, investigator, or just someone looking to gather information, seekr makes it easy to find and manage the data you need. Give it a try and see how it can streamline your OSINT workflow!

Check the wiki for setup guide, etc.


<img width="800" src="https://user-images.githubusercontent.com/67828948/216688806-6cfd4344-e1b6-4a69-870c-ec8d2763c5b7.png">


## Why use seekr over my current tool ?
Seekr combines note taking and OSINT in one application. Seekr can be used alongside your current tools.
Seekr is desingned with OSINT in mind and optimized for real world usecases.
### Key features
- Desktop interface
- Database for OSINT targets
- Integration / adaptation of many popular OSINT-tools (e.g. phoneinfoga)
- GitHub to email
- Account cards for each person in the database
- Account discovery intigrating with the account cards
- Pre defined commonly used fields in the database
## Getting Started - Installation
### Windows
Download and run the latest exe [here](https://github.com/seekr-osint/seekr/releases/latest)

Now open [the web interface](http://127.0.0.1:5050/web/) in your browser of choice.
### Linux (stable)
Download the latest stable binary [here](https://github.com/seekr-osint/seekr/releases/latest)
### Linux (unstable)
To install seekr on linux simply run:
```sh
git clone https://github.com/seekr-osint/seekr
cd seekr
go run main.go
```
Now open [the web interface](http://127.0.0.1:5050/web/) in your browser of choice.
### Run on NixOS
Seekr is build with NixOS in mind and therefore supports nix flakes.
To run seekr on NixOS run following commands.
```sh
nix shell github:seekr-osint/seekr
seekr
```

## Intigrating seekr into your current workflow
```mermaid
journey
	title How to Intigrate seekr into your current workflow.
	section Initial Research
		Create a person in seekr: 100: seekr
    Simple web research: 100: Known tools
		Account scan: 100: seekr
	section Deeper account investigation
		Investigate the accounts: 100: seekr, Known tools
		Keep notes: 100: seekr
  section Deeper Web research
    Deep web research: 100: Known tools
    Keep notes: 100: seekr
	section Finishing the report
		Export the person with seekr: 100: seekr
		Done.: 100
```
## Feedback
We would love to hear from you. Tell us about your opinions on seekr. Where do we need to improve?...
You can do this by just opeing up an issue or maybe even telling others in your blog or somewhere else about your experience.
## Legal Disclaimer
This tool is intended for legitimate and lawful use only. It is provided for educational and research purposes, and should not be used for any illegal or malicious activities, including doxxing. Doxxing is the practice of researching and broadcasting private or identifying information about an individual, without their consent and can be illegal. The creators and contributors of this tool will not be held responsible for any misuse or damage caused by this tool. By using this tool, you agree to use it only for lawful purposes and to comply with all applicable laws and regulations. It is the responsibility of the user to ensure compliance with all relevant laws and regulations in the jurisdiction in which they operate. Misuse of this tool may result in criminal and/or civil prosecution.
## Thanks to
[![Stargazers repo roster for @seekr-osint/seekr](https://reporoster.com/stars/seekr-osint/seekr)](https://github.com/seekr-osint/seekr/stargazers)

- [WinBox.js](https://github.com/nextapps-de/winbox)
- [WAU/discord.go](https://github.com/alpkeskin/wau/blob/main/cmd/apps/discord.go)
