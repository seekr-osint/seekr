const frame = document.querySelector('.frame');

function createSeekrWindow() {
new WinBox("SEEKR", {
  id: "seekr-window",
  html: '<iframe class="frame" src="./index.html"></iframe>',
  background: "#E4EBF5",
  header: 45,
  // viewport boundaries:
  top: 10,
  right: 10,
  bottom: 0,
  left: 10,
});
}

function createSeekrSettingsWindow() {
new WinBox("SEEKR SETTINGS", {
  id: "seekr-window",
  html: '<iframe class="frame" src="./settings.html"></iframe>',
  background: "#E4EBF5",
  header: 45,
  // viewport boundaries:
  top: 10,
  right: 10,
  bottom: 0,
  left: 10,
});
}

function toggleAppBar() {
  var appBar = document.querySelector(".app-bar");
  if (appBar.style.display === "none") {
    appBar.style.display = "block";
  } else {
    appBar.style.display = "none";
  }
}
