const frame = document.querySelector('.frame');

new WinBox("CUSTOM", {
  id: "seekr-window",
  html: '<iframe class="frame" src="https://www.google.com/webhp?igu=1"></iframe>',
  background: "#E4EBF5",
  header: 45,
  // viewport boundaries:
  top: 10,
  right: 10,
  bottom: 0,
  left: 10,
});

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