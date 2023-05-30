class Translate {
  attribute: string;
  lng: string;

  constructor(attr: string, language: string) {
    this.attribute = attr;
    this.lng = language;
  }

  translateElement(element: HTMLElement): void {
    const xhrFile = new XMLHttpRequest();

    xhrFile.open("GET", `translations/${this.lng}.json`, false);
    xhrFile.onreadystatechange = () => {
      if (xhrFile.readyState === 4) {
        if (xhrFile.status === 200 || xhrFile.status === 0) {
          const LngObject = JSON.parse(xhrFile.responseText);
          const key = element.getAttribute(this.attribute);
          if (key !== null) {
            if (element.hasAttribute("placeholder")) {
              element.setAttribute("placeholder", LngObject[key]);
            } else if (element instanceof HTMLInputElement) {
              element.value = LngObject[key];
            } else if (element instanceof HTMLTextAreaElement) {
              element.value = LngObject[key];
            } else if (element instanceof HTMLSelectElement) {
              element.value = LngObject[key];
            } else {
              element.innerHTML = LngObject[key];
            }
          }
        }
      }
    };
    xhrFile.send();
  }

  translateText(word: string): string | undefined {
    const xhrFile = new XMLHttpRequest();
    let translatedWord: string | undefined;
  
    xhrFile.open("GET", `translations/${this.lng}.json`, false);
    xhrFile.onreadystatechange = () => {
      if (xhrFile.readyState === 4) {
        if (xhrFile.status === 200 || xhrFile.status === 0) {
          const LngObject = JSON.parse(xhrFile.responseText);
          translatedWord = LngObject[word];
        }
      }
    };
    xhrFile.send();
  
    return translatedWord;
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

function translateText(word: string): string | undefined {
  const translator = new Translate("lng-tag", localStorage.getItem("language") || "en");
  return translator.translateText(word);
}

function translateRawWord(word: string): string | undefined {
  word = word.toLowerCase().replace(/ /g, "_").replace(/:/g, "_colon");

  console.log(word);

  const translator = new Translate("lng-tag", localStorage.getItem("language") || "en");
  return translator.translateText(word);
}

console.log(translateRawWord("Pets:"));