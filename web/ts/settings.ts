const channel = new BroadcastChannel("seekr-channel");

channel.addEventListener('message', (event) => {
  if (event.data.type === "theme") {
    const theme = event.data.theme;

    document.documentElement.setAttribute("data-theme", theme);
  } else if (event.data.type === "language") {
    translate()
  }
});

const themeCardContainer = document.querySelector(".theme-option") as HTMLDivElement;

function changeTheme(theme: string): void {
  // Get the file path of the target iframe
  const targetFilePaths = ["./lite.html", "./guide.html", "./desktop.html", "./index.html", "./settings.html", "./tempmail.html"];

  // Send a message to the broadcast channel with the target file path and the new state of the switch for each html file in targetFilePaths
  targetFilePaths.forEach(targetFilePath => {
    channel.postMessage({ type: "theme", targetFilePath, theme });
  });
  localStorage.setItem("theme", theme);
  document.documentElement.setAttribute("data-theme", theme);
}

function createThemeCards(theme: string) {
  if (themeCardContainer ) {
    const themeCardOuter = document.createElement("div") as HTMLDivElement;
    themeCardOuter.id = `${theme}-theme`;
    themeCardOuter.classList.add("theme-card", "big-card", "chip");
    themeCardOuter.dataset.theme = theme;

    const themeCardInner = document.createElement("div") as HTMLDivElement;
    themeCardInner.classList.add("theme-card", "small-card", "chip");
    themeCardInner.dataset.theme = theme;

    const themeCardColors = document.createElement("div") as HTMLDivElement;
    themeCardColors.classList.add("colors");

    const themeCardText = document.createElement("p") as HTMLParagraphElement;

    themeCardText.classList.add("theme-text");
    themeCardText.dataset.theme = theme;
    themeCardText.textContent = theme.charAt(0).toUpperCase() + theme.slice(1);


    themeCardContainer.appendChild(themeCardOuter);
    themeCardOuter.appendChild(themeCardInner);
    themeCardInner.appendChild(themeCardColors);
    themeCardColors.appendChild(themeCardText);

    themeCardOuter.addEventListener("click", () => {
      changeTheme(theme);
    });
  }
}

const selectedLanguage = document.querySelector(".language-select > .select-selected") as HTMLDivElement;

function checkLanguage(): "en" | "de" | undefined {
  if (document) {
    if (selectedLanguage) {
      const languages: { [key: string]: "en" | "de" } = {};

      // English

      languages["English"] = "en";
      languages["German"] = "de";

      // Translations

      if (languages[selectedLanguage.innerHTML] == undefined) {
        languages[translateText("english")!] = "en";
        languages[translateText("german")!] = "de";
      }

      return languages[selectedLanguage.innerHTML];
    }
  }
}

function handleLanguageChange() {
  const language = checkLanguage();

  if (language) {
    const targetFilePaths = ["./lite.html", "./guide.html", "./desktop.html", "./index.html", "./settings.html"];

    setLanguage(language);

    translate();

    targetFilePaths.forEach(targetFilePath => {
      channel.postMessage({ type: "language", targetFilePath, language });
    });
  }
}

function preLanguageChangeHandler () {
  if (selectedLanguage && selectedLanguage.innerHTML != "") {
    handleLanguageChange();
  }
}

if (selectedLanguage) {
  setTimeout(() => {
    selectedLanguage.addEventListener("DOMSubtreeModified", preLanguageChangeHandler);
  }, 100); // Triggered when loaded, this is a workaround (might cause problems on slow devices)
}

export { createThemeCards, changeTheme, checkLanguage };