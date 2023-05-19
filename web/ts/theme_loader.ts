import { createThemeCards, changeTheme } from "./settings.js";

const head = document.getElementsByTagName("head")[0];
const cssFolder = "./themes/";

const defaultTheme = "arctic";

// This loads all css files in the themes directory

function setDefaultIfNotStored(): void {
  if (!localStorage.getItem("theme")) {
    localStorage.setItem("theme", defaultTheme);

    changeTheme(defaultTheme);

    console.log(localStorage.getItem("theme"));
  }
}

setDefaultIfNotStored();

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

function recolorImage(img: any, oldRed: Number, oldGreen: Number, oldBlue: Number, newRed: Number, newGreen: Number, newBlue: Number) {
  var c = document.createElement('canvas');
  var ctx = c.getContext("2d");
  var w = img.width;
  var h = img.height;

  c.width = w;
  c.height = h;

  // draw the image on the temporary canvas
  ctx.drawImage(img, 0, 0, w, h);

  // pull the entire image into an array of pixel data
  var imageData = ctx.getImageData(0, 0, w, h);

  // examine every pixel, 
  // change any old rgb to the new-rgb
  for (var i = 0; i < imageData.data.length; i += 4) {
      // is this pixel the old rgb?
      if (imageData.data[i] == oldRed && imageData.data[i + 1] == oldGreen && imageData.data[i + 2] == oldBlue) {
          // change to your new rgb
          imageData.data[i] = newRed;
          imageData.data[i + 1] = newGreen;
          imageData.data[i + 2] = newBlue;
      }
  }
  // put the altered data back on the canvas  
  ctx.putImageData(imageData, 0, 0);
  // put the re-colored image back on the image
  var img1 = document.getElementById("image1");
  img1.src = c.toDataURL('image/png');

}