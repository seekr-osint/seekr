<p align="center">
  <a href="https://github.com/seekr-osint/seekr" target="blank"><img src="https://user-images.githubusercontent.com/69043370/212559304-17f44716-e748-4ac2-b86b-097ad1e49bb3.png" width="250" alt="Seekr Logo" /></a>
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
eekrS is a multi-purpose toolkit for gathering and managing OSINT-data with a sleek web interface. The backend is written in Go and offers a wide range of features for data collection, organization, and analysis. Whether you're a researcher, investigator, or just someone looking to gather information, SEEKR makes it easy to find and manage the data you need. Give it a try and see how it can streamline your OSINT workflow!

Check the wiki for setup guide, etc.
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

## Getting Started
To install seekr on linux simply run:
```sh
git clone https://github.com/seekr-osint/seekr
cd seekr
go run main.go
```
Now open [the web interface](http://localhost:5050) in your browser of choice.

## Legal Disclaimer
This tool is intended for legitimate and lawful use only. It is provided for educational and research purposes, and should not be used for any illegal or malicious activities, including doxxing. Doxxing is the practice of researching and broadcasting private or identifying information about an individual, without their consent and can be illegal. The creators and contributors of this tool will not be held responsible for any misuse or damage caused by this tool. By using this tool, you agree to use it only for lawful purposes and to comply with all applicable laws and regulations. It is the responsibility of the user to ensure compliance with all relevant laws and regulations in the jurisdiction in which they operate. Misuse of this tool may result in criminal and/or civil prosecution.
## Thanks to
[![Stargazers repo roster for @seekr-osint/seekr](https://reporoster.com/stars/seekr-osint/seekr)](https://github.com/seekr-osint/seekr/stargazers)
