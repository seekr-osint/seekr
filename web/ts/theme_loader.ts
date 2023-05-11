import { createThemeCards } from "./settings.js";

const head = document.getElementsByTagName("head")[0];
const cssFolder = "./themes/";

// This loads all css files in the themes directory

fetch(cssFolder)
  .then((response) => response.text())
  .then((html) => {
    const parser = new DOMParser();
    const doc = parser.parseFromString(html, "text/html");
    const links = doc.querySelectorAll('a[href$=".css"]');

    links.forEach((link) => {
      const href = cssFolder + link.getAttribute("href");
      const cssLink = document.createElement("link");

      cssLink.rel = "stylesheet";
      cssLink.type = "text/css";
      cssLink.href = href;

      head.appendChild(cssLink);

      const theme = href.trim().replace(/^\.\/themes\/|\.css$/g, '');

      if (theme != "default") {
        createThemeCards(theme);
      }
    });
  }
);