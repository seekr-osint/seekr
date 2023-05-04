const bg_var = getComputedStyle(document.documentElement).getPropertyValue('[data-theme]--bg');

// Create a new broadcast channel with the same name as in the first code block
const channel = new BroadcastChannel("theme-channel");

// Listen for messages on the broadcast channel
channel.addEventListener('message', (event) => {
  if (event.data.type === "theme") {
    const theme = event.data.theme;
    localStorage.setItem("theme", theme);

    document.documentElement.setAttribute("data-theme", theme);
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
    id: "seekr-window",
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
shortcutMenu.addEventListener('click', function() {
  appBar.classList.toggle('clicked');
});
