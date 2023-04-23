const defaultWhiteTheme = document.getElementById("default-white-theme");
const defaultDarkTheme = document.getElementById("default-dark-theme");

const nordDarkTheme = document.getElementById("nord-dark-theme");

const channel = new BroadcastChannel("theme-channel");


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


if (defaultWhiteTheme) {
  defaultWhiteTheme.addEventListener("click", () => {
    changeTheme("default-white");
  });
}

if (defaultDarkTheme) {
  defaultDarkTheme.addEventListener("click", () => {
    changeTheme("default-dark");
  });
}

if (nordDarkTheme) {
  nordDarkTheme.addEventListener("click", () => {
    changeTheme("nord-dark");
  });
}
