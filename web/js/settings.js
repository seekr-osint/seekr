const frame = document.querySelector('.frame');

function checkSettingsFromLocalStorage() {
  const isDarkMode = localStorage.getItem("isDarkMode");
  const isDarkModeChecked = isDarkMode === "true" ? true : false;

  document.querySelector(".darkMode-switch-input").checked = isDarkModeChecked;
}


checkSettingsFromLocalStorage();



// Create a new broadcast channel with a specific name
const channel = new BroadcastChannel('dark-mode-channel');

document.querySelector(".darkMode-switch-input").addEventListener("change", e => {
  const isDarkMode = e.target.checked;

  // Get the file path of the target iframe
  const targetFilePaths = ["./lite.html", "./guide.html", "./desktop.html", "./index.html"];

  if (isDarkMode) {
    document.documentElement.setAttribute('data-theme', 'dark');
  } else {
    document.documentElement.setAttribute('data-theme', 'light');
  }

  // Send a message to the broadcast channel with the target file path and the new state of the switch for each html file in targetFilePaths
  targetFilePaths.forEach(targetFilePath => {
    channel.postMessage({ type: 'dark-mode', targetFilePath, isDarkMode });
    localStorage.setItem("isDarkMode", true);
  });
});
