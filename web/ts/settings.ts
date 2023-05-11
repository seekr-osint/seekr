const channel = new BroadcastChannel("theme-channel");

const themeCardContainer = document.querySelector(".theme-option") as HTMLDivElement;

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

function createThemeCards(theme: string) {
  if (themeCardContainer ) {
    const themeCardOuter = document.createElement("div") as HTMLDivElement;
    themeCardOuter.id = `${theme}-theme`;
    themeCardOuter.classList.add("theme-card", "big-card", "chip");
    themeCardOuter.dataset.theme = theme;

    const themeCardInner = document.createElement("div") as HTMLDivElement;
    themeCardInner.classList.add("theme-card", "small-card", "chip");
    themeCardInner.dataset.theme = theme;

    const themeCardColors = document.createElement("div") as HTMLDivElement;
    themeCardColors.classList.add("colors");

    const themeCardText = document.createElement("p") as HTMLParagraphElement;

    themeCardText.classList.add("theme-text");
    themeCardText.dataset.theme = theme;
    themeCardText.textContent = theme.charAt(0).toUpperCase() + theme.slice(1);


    themeCardContainer.appendChild(themeCardOuter);
    themeCardOuter.appendChild(themeCardInner);
    themeCardInner.appendChild(themeCardColors);
    themeCardColors.appendChild(themeCardText);

    themeCardOuter.addEventListener("click", () => {
      changeTheme(theme);
    });
  }
}

export { createThemeCards };
