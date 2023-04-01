const bg_var = getComputedStyle(document.documentElement).getPropertyValue('[data-theme]--bg');


window.addEventListener('message', (event) => {
  if (event.data.type === 'dark-mode') {
    const isDarkMode = event.data.isDarkMode;
    localStorage.setItem('isDarkMode', isDarkMode);
    
    // Update the dark mode switch
    darkModeSwitch.checked = isDarkMode;
    
    // Update all iframes
    const iframes = document.querySelectorAll('iframe');
    iframes.forEach((iframe) => {
      iframe.contentWindow.postMessage({ type: 'dark-mode', isDarkMode }, '*');
    });
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