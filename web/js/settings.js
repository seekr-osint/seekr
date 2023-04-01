const frame = document.querySelector('.frame');

function checkSettingsFromLocalStorage() {
  const isDarkMode = localStorage.getItem("isDarkMode");
  const isDarkModeChecked = isDarkMode === "true" ? true : false;

  document.querySelector(".darkMode-switch-input").checked = isDarkModeChecked;
}


checkSettingsFromLocalStorage();



document.querySelector(".darkMode-switch-input").addEventListener("change", e => {
    const isDarkMode = e.target.checked;


    if(isDarkMode === true) {
        localStorage.setItem("isDarkMode", true);
        window.parent.postMessage({ type: 'dark-mode', isDarkMode }, '*');
    } else {
        localStorage.setItem("isDarkMode", false);
        window.parent.postMessage({ type: 'dark-mode', isDarkMode }, '*');
    }
});