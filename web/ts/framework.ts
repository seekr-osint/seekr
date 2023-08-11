declare function saveAs(blob: Blob, filename: string): void;

function saveAsFile(textContent: string, fileName: string) {
  // saveAsFile("text","filename.txt");
  const textEncoding: string = "text/plain;charset=utf-8";

  try {
    var blob = new Blob([textContent], { type: textEncoding });
    saveAs(blob, fileName);
  } catch (exception) {
    window.open("data:" + textEncoding + "," + encodeURIComponent(textContent), '_blank', '');
  }
}

function loadDropdown(dropdownType: string, data: string) {
  const scrollbox = document.querySelector<HTMLDivElement>("body > .edit-container > div > div.scroll-box")!;
  const dropdownElement = scrollbox.querySelector("custom-dropdown[title='" + dropdownType + "']")!.shadowRoot!.querySelector("div > .table > .dropdown-select > .select-selected") as HTMLElement;

  if (data != "") {
    dropdownElement.innerHTML = data;
  }
}

function checkDropdownValue(windowType: "edit", dropdownType: string) {
  const scrollbox = document.querySelector<HTMLDivElement>("body > div." + windowType + "-container > div > div.scroll-box")!;

  const selectedType = scrollbox.querySelector("custom-dropdown[title='" + dropdownType + "']")!.shadowRoot!.querySelector("div > .table > .dropdown-select > .select-selected")!.innerHTML ?? "";

  if (selectedType != "Select " + dropdownType + ":") {
    return "";
  } else {
    return selectedType;
  }
}

function apiCall(endpoint: string): string {
  var hostname = window.location.hostname;
  var port = window.location.port;
  var baseUrl = hostname + ":" + port;
  var apiUrl = "http://" + baseUrl + "/api";

  if (endpoint.charAt(0) === "/") {
    endpoint = endpoint.substring(1);
  }

  return apiUrl + '/' + endpoint
}

export { saveAsFile, loadDropdown, checkDropdownValue, apiCall };
