import { createThemeCards, changeTheme } from "./settings.js";

const head = document.getElementsByTagName("head")[0];
const cssFolder = "./themes";

const localStorageTheme = localStorage.getItem("theme");

const defaultTheme = "arctic";

// This loads all css files in the themes directory

function setDefaultIfNotStored(): void {
  if (!localStorage.getItem("theme")) {
    localStorage.setItem("theme", defaultTheme);

    changeTheme(defaultTheme);
  }
}

setDefaultIfNotStored();

fetch(cssFolder)
  .then((response) => response.text())
  .then((html) => {
    const parser = new DOMParser();
    const doc = parser.parseFromString(html, "text/html");
    const links = doc.querySelectorAll('a[href$=".css"]');

    let hasThemeBeenApplied = false;

    links.forEach((link) => {
      const href = link.getAttribute("href");
      const cssLink = document.createElement("link");

      cssLink.rel = "stylesheet";
      cssLink.type = "text/css";
      cssLink.href = href!;

      head.appendChild(cssLink);

      const theme = href!.trim().replace("/web/themes/", "").replace(".css", "");
      
      if (localStorageTheme == theme) {
        hasThemeBeenApplied = true;
      }

      createThemeCards(theme);
    });

    if (!hasThemeBeenApplied) {
      changeTheme(defaultTheme);
    }
  }
);
