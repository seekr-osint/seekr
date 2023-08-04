const bg_var = getComputedStyle(document.documentElement).getPropertyValue('[data-theme]--bg');

const channelDesktop = new BroadcastChannel("seekr-channel");
// Listen for messages on the broadcast channel
channelDesktop.addEventListener('message', (event) => {
  if (event.data.type === "theme") {
    const theme = event.data.theme;

    document.documentElement.setAttribute("data-theme", theme);

    replaceIconColors();
  }
});


function createSeekrWindow() {
  new WinBox("SEEKR", {
    id: "seekr-window",
    html: '<iframe class="frame" src="./lite.html"></iframe>',
    background: bg_var,
    header: 45,
    // viewport boundaries:
    top: 10,
    right: 10,
    bottom: 0,
    left: 10,
  });
}

function createGuideWindow() {
  new WinBox("GUIDE", {
    id: "seekr-window",
    html: '<iframe class="frame" src="./guide.html"></iframe>',
    background: bg_var,
    header: 45,
    // viewport boundaries:
    top: 10,
    right: 10,
    bottom: 0,
    left: 10,
  });
}

function createCrtshWindow() {
  new WinBox("CRTSH", {
    id: "seekr-window-crtsh",
    html: '<iframe class="frame" src="https://crt.sh"></iframe>',
    background: bg_var,
    header: 45,
    // viewport boundaries:
    top: 10,
    right: 10,
    bottom: 0,
    left: 10,
  });
}

function createBlockchainExplorerWindow() {
  new WinBox("Blockchain Explorer", {
    id: "seekr-window",
    html: '<iframe class="frame" src="https://blockexplorer.one"></iframe>',
    background: bg_var,
    header: 45,
    // viewport boundaries:
    top: 10,
    right: 10,
    bottom: 0,
    left: 10,
  });
}

function createTempMailWindow() {
  new WinBox("TEMPMAIL", {
    id: "tempmail-window",
    html: '<iframe class="frame" src="./tempmail.html"></iframe>',
    background: bg_var,
    header: 45,
    // viewport boundaries:
    top: 10,
    right: 10,
    bottom: 0,
    left: 10,
  });
}

function createWhoisWindow() {
  new WinBox("WHOIS", {
    id: "seekr-window",
    html: '<iframe class="frame" src="https://who.is"></iframe>',
    background: bg_var,
    header: 45,
    // viewport boundaries:
    top: 10,
    right: 10,
    bottom: 0,
    left: 10,
  });
}

function createSeekrSettingsWindow() {
  new WinBox("SETTINGS", {
    id: "seekr-window",
    html: '<iframe class="frame" src="./settings.html"></iframe>',
    background: bg_var,
    header: 45,
    // viewport boundaries:
    top: 10,
    right: 10,
    bottom: 0,
    left: 10,
  });
}


// const appIcon = document.querySelector('.app-icon');
// const seekrIcon = document.querySelector('.seekr-icon');
// appIcon.addEventListener("hover", function() {
//   appIcon.src = "./img/seekr-icon.png";
// });

const shortcutMenu = document.querySelector('.shortcut-menu');
const appBar = document.querySelector('.app-bar');
shortcutMenu!.addEventListener('click', function() {
  appBar!.classList.toggle('clicked');
});


// Replace the hex color code with the desired one
const replaceHexColor = (element: HTMLImageElement, originalHex: string, newHex: string, tolerance: number) => {
  const canvas = document.createElement("canvas");
  const context = canvas.getContext("2d") as CanvasRenderingContext2D;
  const width = element.width;
  const height = element.height;
  canvas.width = width;
  canvas.height = height;

  context.drawImage(element, 0, 0, width, height);

  const imageData = context.getImageData(0, 0, width, height);
  const data = imageData.data;

  const originalRGB = hexToRgb(originalHex);
  const newRGB = hexToRgb(newHex);

  if (originalRGB && newRGB) {
    for (let i = 0; i < data.length; i += 4) {
      const r = data[i];
      const g = data[i + 1];
      const b = data[i + 2];

      if (
        Math.abs(r - originalRGB.r) <= tolerance &&
        Math.abs(g - originalRGB.g) <= tolerance &&
        Math.abs(b - originalRGB.b) <= tolerance
      ) {
        data[i] = newRGB.r;
        data[i + 1] = newRGB.g;
        data[i + 2] = newRGB.b;
      }
    }
  }

  context.putImageData(imageData, 0, 0);
  element.src = canvas.toDataURL();
};

// Helper function to convert hex color to RGB
const hexToRgb = (hexColor: string) => {
  const shorthandRegex = /^#?([a-f\d])([a-f\d])([a-f\d])$/i;
  const hex = hexColor.replace(shorthandRegex, (m, r, g, b) => {
    return r + r + g + g + b + b;
  });

  const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex);
  return result
    ? { r: parseInt(result[1], 16), g: parseInt(result[2], 16), b: parseInt(result[3], 16) }
    : null;
};

const appIcons = document.querySelectorAll(".app-icon") as NodeListOf<HTMLImageElement>;

let hasBeenLoaded = false;

// While hasBeenLoaded is false, repeat every 100ms
const interval = setInterval(() => {
  if (!hasBeenLoaded) {
    if (localStorage.getItem("theme")) {
      replaceIconColors();
    } else {
      localStorage.setItem("theme", "arctic");

      replaceIconColors();
    }
  }
}, 100);

let lastReplaced = "#9BAACF";

function replaceIconColors() {
  // If the .css link has been loaded
  if (document.head.innerHTML.includes(localStorage.getItem("theme")!)) {
    const computedStyle = getComputedStyle(document.documentElement);

    const currentTextColor = computedStyle.getPropertyValue("--textColor-1").trim();

    // Replace the hex color code with the desired one
    appIcons.forEach((appIcon) => {
      replaceHexColor(appIcon, lastReplaced, currentTextColor, 10);
    });

    lastReplaced = currentTextColor;
    hasBeenLoaded = true;
    clearInterval(interval);
  }
}