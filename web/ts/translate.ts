let languageData: { [key: string]: { [key: string]: string } } | undefined = undefined;

class Translate {
  attribute: string; lng: string;

  constructor(attr: string, language: string) {
    this.attribute = attr;
    this.lng = language;
  }

  private loadLanguageData(): void {
    if (languageData && languageData[this.lng]) {
      return; // Language data is already loaded
    }

    const xhrFile = new XMLHttpRequest();
    xhrFile.open("GET", `translations/${this.lng}.json`, false);
    xhrFile.onreadystatechange = () => {
      if (xhrFile.readyState === 4 && (xhrFile.status === 200 || xhrFile.status === 0)) {
        languageData = languageData || {};
        languageData[this.lng] = JSON.parse(xhrFile.responseText);
      }
    };
    xhrFile.send();
  }

  private getTranslation(key: string): string | undefined {
    this.loadLanguageData();
    return languageData && languageData[this.lng] ? languageData[this.lng][key] : undefined;
  }

  translateElement(element: HTMLElement): void {
    const key = element.getAttribute(this.attribute);
    if (key !== null) {
      const translation = this.getTranslation(key);
      if (translation !== undefined) {
        if (element.hasAttribute("placeholder")) {
          element.setAttribute("placeholder", translation);
        } else if (element instanceof HTMLInputElement || element instanceof HTMLTextAreaElement || element instanceof HTMLSelectElement) {
          element.value = translation;
        } else {
          element.innerHTML = translation;
        }
      }
    }
  }

  translateText(word: string): string | undefined {
    return this.getTranslation(word);
  }

  translateAllElements(): void {
    const allDom = document.querySelectorAll(`[${this.attribute}]`);
    allDom.forEach((element) => {
      if (element instanceof HTMLElement) {
        this.translateElement(element);
      }
    });
  }
}

// This function will be called when the user clicks to change the language
function translate(): void {
  let lang = localStorage.getItem("language");

  if (!lang) {
    lang = "en";
    
    setLanguage(lang);
  }

  const translator = new Translate("lng-tag", lang);
  translator.translateAllElements();
}

// This function will be called when the user clicks to change the language
function onLoadTranslate(): void {
  let lang = localStorage.getItem("language");

  if (!lang) {
    lang = "en";
    
    setLanguage(lang);
  } else if (lang == "en") {
    return;
  }

  const translator = new Translate("lng-tag", lang);
  translator.translateAllElements();
}

// This function is used to refresh translation
function refreshTranslation(): void {
  let lang = localStorage.getItem("language");

  if (!lang) {
    lang = "en";
    
    setLanguage(lang);
  }

  const translator = new Translate("lng-tag", lang);
  translator.translateAllElements();
}

function setLanguage(language: string) {
  localStorage.setItem("language", language);
}

function translateElement(element: HTMLElement): void {
  const translator = new Translate("lng-tag", localStorage.getItem("language") || "en");
  translator.translateElement(element);
}

function translateText(word: string, customLang?: string): string | undefined {
  if (customLang) {
    const translator = new Translate("lng-tag", customLang);

    return translator.translateText(word);
  } else {
    const translator = new Translate("lng-tag", localStorage.getItem("language") || "en");

    return translator.translateText(word);
  }
}

function translateRawWord(word: string): string | undefined {
  if (word != "" && word != undefined) {
    word = word.toLowerCase().replace(/\//g, "_slash_").replace(/:/g, "_colon").replace(/ /g, "_");

    const translator = new Translate("lng-tag", localStorage.getItem("language") || "en");

    return translator.translateText(word);
  } else {
    return "";
  }
}
