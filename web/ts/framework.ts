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
type DropdownType = "gender" | "ethnicity" | "religion" | "civil status" | "language";

function getDropdownElementIndex(dropdownType: DropdownType, dropdownValue: string, customLangParameter?: string): string {
  if (dropdownType == "gender") {
    const genderIndex: { [key: string]: string } = {};

    genderIndex[""] = "";

    // English

    genderIndex["Male"] = "0";
    genderIndex["Female"] = "1";
    genderIndex["Other"] = "2";

    return genderIndex[dropdownValue];
  } else if (dropdownType == "ethnicity") {
    const ethnicityIndex: { [key: string]: string } = {};

    ethnicityIndex[""] = "";

    // English

    ethnicityIndex["African"] = "0";
    ethnicityIndex["Asian"] = "1";
    ethnicityIndex["Caucasian/White"] = "2";
    ethnicityIndex["Hispanic/Latino"] = "3";
    ethnicityIndex["Indigenous/Native American"] = "4";
    ethnicityIndex["Multiracial/Mixed"] = "5";

    return ethnicityIndex[dropdownValue];
  } else if (dropdownType == "religion") {
    const religionIndex: { [key: string]: string } = {};

    religionIndex[""] = "";

    // English

    religionIndex["Christianity"] = "0";
    religionIndex["Atheism"] = "1";
    religionIndex["Islam"] = "2";
    religionIndex["Hinduism"] = "3";
    religionIndex["Buddhism"] = "4";
    religionIndex["Sikhism"] = "5";
    religionIndex["Judaism"] = "6";
    religionIndex["Other"] = "7";

    return religionIndex[dropdownValue];
  } else if (dropdownType == "civil status") {
    const civilstatusIndex: { [key: string]: string } = {};

    civilstatusIndex[""] = "";

    // English

    civilstatusIndex["Single"] = "0";
    civilstatusIndex["Married"] = "1";
    civilstatusIndex["Widowed"] = "2";
    civilstatusIndex["Divorced"] = "3";
    civilstatusIndex["Separated"] = "4";

    return civilstatusIndex[dropdownValue];
  } else if (dropdownType == "language") {
    const languageIndex: { [key: string]: string } = {};

    languageIndex[""] = "";

    // English

    languageIndex["English"] = "0";
    languageIndex["Spanish"] = "1";
    languageIndex["German"] = "2";
    languageIndex["Italian"] = "3";
    languageIndex["Gaelic"] = "4";
    languageIndex["Latin"] = "5";

    return languageIndex[dropdownValue];
  }

  return "";
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

export { saveAsFile, loadDropdown, getDropdownElementIndex, checkDropdownValue, apiCall, DropdownType };
