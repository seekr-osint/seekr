const whiteTheme = document.getElementById("white-theme");
const darkTheme = document.getElementById("dark-theme");

const channel = new BroadcastChannel("theme-channel");

function changeTheme(theme) {
  // Get the file path of the target iframe
  const targetFilePaths = ["./lite.html", "./guide.html", "./desktop.html", "./index.html", "./settings.html", "./tempmail.html"];

  // Send a message to the broadcast channel with the target file path and the new state of the switch for each html file in targetFilePaths
  targetFilePaths.forEach(targetFilePath => {
    channel.postMessage({ type: "theme", targetFilePath, theme });
    localStorage.setItem("theme", theme);
  });

  document.documentElement.setAttribute("data-theme", localStorage.getItem("theme"));
}


whiteTheme.addEventListener("click", () => {
  changeTheme("white");
});

darkTheme.addEventListener("click", () => {
  changeTheme("dark");
});