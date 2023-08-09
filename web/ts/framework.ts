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

function loadDropdown(dropdownType: "gender" | "ethnicity" | "religion" | "civil status", data: string) {
  const scrollbox = document.querySelector<HTMLDivElement>("body > .edit-container > div > div.scroll-box")!;
  const dropdownElement = scrollbox.querySelector("custom-dropdown[title='" + dropdownType + "']")!.shadowRoot!.querySelector("div > .table > .dropdown-select > .select-selected") as HTMLElement;

  if (data != "") {
    dropdownElement.innerHTML = data;
  }
}

function checkDropdownValue(windowType: "edit" | "create", dropdownType: "gender" | "ethnicity" | "religion" | "civil status") {
  const scrollbox = document.querySelector<HTMLDivElement>("body > div." + windowType + "-container > div > div.scroll-box")!;

  if (dropdownType == "gender") {
    const selectedGender = scrollbox.querySelector("custom-dropdown[title='gender']")!.shadowRoot!.querySelector("div > .table > .dropdown-select > .select-selected")!.innerHTML ?? "";
    const gender: { [key: string]: string } = {};

    // English

    gender["Select gender:"] = "";
    gender["Male"] = "Male";
    gender["Female"] = "Female";
    gender["Other"] = "Other";

    return gender[selectedGender];
  } else if (dropdownType == "ethnicity") {
    const selectedEthnicity = scrollbox.querySelector("custom-dropdown[title='ethnicity']")!.shadowRoot!.querySelector("div > .table > .dropdown-select > .select-selected")!.innerHTML ?? "";
    const ethnicity: { [key: string]: string } = {};

    // English

    ethnicity["Select ethnicity:"] = "";
    ethnicity["African"] = "African";
    ethnicity["Asian"] = "Asian";
    ethnicity["Caucasian/White"] = "Caucasian/White";
    ethnicity["Hispanic/Latino"] = "Hispanic/Latino";
    ethnicity["Indigenous/Native American"] = "Indigenous/Native American";
    ethnicity["Multiracial/Mixed"] = "Multiracial/Mixed";

    return ethnicity[selectedEthnicity];
  } else if (dropdownType == "religion") {
    const selectedReligion = scrollbox.querySelector("custom-dropdown[title='religion']")!.shadowRoot!.querySelector("div > .table > .dropdown-select > .select-selected")!.innerHTML ?? "";
    const religion: { [key: string]: string } = {};

    // English

    religion["Select religion:"] = "";
    religion["Christianity"] = "Christianity";
    religion["Atheism"] = "Atheism";
    religion["Islam"] = "Islam";
    religion["Hinduism"] = "Hinduism";
    religion["Buddhism"] = "Buddhism";
    religion["Sikhism"] = "Sikhism";
    religion["Judaism"] = "Judaism";
    religion["Other"] = "Other";

    return religion[selectedReligion];
  } else if (dropdownType == "civil status") {
    const selectedCivilstatus = scrollbox.querySelector("custom-dropdown[title='civil status']")!.shadowRoot!.querySelector("div > .table > .dropdown-select > .select-selected")!.innerHTML ?? "";
    const civilstatus: { [key: string]: string } = {};

    // English

    civilstatus["Select civil status:"] = "";
    civilstatus["Single"] = "Single";
    civilstatus["Married"] = "Married";
    civilstatus["Widowed"] = "Widowed";
    civilstatus["Divorced"] = "Divorced";
    civilstatus["Separated"] = "Separated";

    return civilstatus[selectedCivilstatus];
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
