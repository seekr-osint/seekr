const arcticTheme = document.getElementById("arctic-theme");
const duskTheme = document.getElementById("dusk-theme");

const nordTheme = document.getElementById("nord-theme");

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


if (arcticTheme) {
  arcticTheme.addEventListener("click", () => {
    changeTheme("arctic");
  });
}

if (duskTheme) {
  duskTheme.addEventListener("click", () => {
    changeTheme("dusk");
  });
}

if (nordTheme) {
  nordTheme.addEventListener("click", () => {
    changeTheme("nord");
  });
}

export {};