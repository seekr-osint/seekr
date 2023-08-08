import { apiCall } from './framework.js';
import * as tsGenConfig from './../ts-gen/config.js'

const channel = new BroadcastChannel("seekr-channel");

channel.addEventListener('message', (event) => {
  if (event.data.type === "theme") {
    const theme = event.data.theme;

    document.documentElement.setAttribute("data-theme", theme);
  }
});

// Theme stuff

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

// SEEKR config stuff

type SeekrConfig = tsGenConfig.Config
//interface SeekrConfig {
//  general: {
//    browser: boolean;
//    discord: boolean;
//    force_port: boolean;
//  };
//  server: {
//    ip: string;
//    port: number;
//  };
//}

async function restartSeekr(): Promise<void> {
  await fetch(apiCall('/restart'));
  return;
}

async function getCurrentConfig(): Promise<SeekrConfig> {
  const response = await fetch(apiCall('/config'));
  const data = await response.json();
  return data as SeekrConfig;
}

async function postSeekrConfig(config: SeekrConfig): Promise<string> {
  const response = await fetch(apiCall('/config'), {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(config),
  });
  const data = await response.json();
  return data.message;
}

function setValues(config: SeekrConfig): void {
  // Port
  const portInput = document.getElementById("port-tag") as HTMLInputElement;
  if (portInput) {
    portInput.value = config.server.port.toString();
  }

  // Port
  const ipInput = document.getElementById("ip-tag") as HTMLInputElement;

  if (ipInput) {
    ipInput.value = config.server.ip.toString();
  }
}

function getValues(): { port: number, ip: string } {
  const portInput = document.getElementById("port-tag") as HTMLInputElement;
  const port = parseInt(portInput.value, 10);

  const ipInput = document.getElementById("ip-tag") as HTMLInputElement;
  const ip = ipInput.value;

  return { port, ip };
}

async function getUpdatedSeekrConfig(): Promise<SeekrConfig> {
  let currentConfig = await getCurrentConfig();
  const values = getValues();

  currentConfig.server.ip = values.ip;
  currentConfig.server.port = values.port;

  return currentConfig;
}

document.addEventListener('DOMContentLoaded', () => {
  const saveBtn = document.getElementById('settings-savebtn-p');
  if (saveBtn) {
    (async () => {
      const currentConfig = await getCurrentConfig();
      // console.log('Current Config:', currentConfig);
      setValues(currentConfig);
    })();

    saveBtn.addEventListener('click', async () => {
      const newConfig = await getUpdatedSeekrConfig();
      const message = await postSeekrConfig(newConfig);
      restartSeekr();

      setTimeout(() => {
        window.location.reload();
      }, 1000);
    });
  }
});


export { createThemeCards, changeTheme };
