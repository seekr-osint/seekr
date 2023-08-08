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

const elements: { [key: string]: HTMLCollectionOf<Element> } = { // used to select a element in the dropdown
  "gender-select": document.getElementsByClassName("gender-select"),
  "ethnicity-select": document.getElementsByClassName("ethnicity-select"),
  "religion-select": document.getElementsByClassName("religion-select"),
  "civilstatus-select": document.getElementsByClassName("civilstatus-select"),
  "country-select": document.getElementsByClassName("country-select"),
  "language-select": document.getElementsByClassName("language-select")
};

for (const [className, nodeList] of Object.entries(elements)) {
  for (let i = 0; i < nodeList.length; i++) {
    const node = nodeList[i] as HTMLElement;
    const selElmnt = node.getElementsByTagName("select")[0] as HTMLSelectElement;
    const selElmntLength = selElmnt.length;
    let selectSelectedDiv = document.createElement("DIV");
    selectSelectedDiv.setAttribute("class", "select-selected");

    const labelText = selElmnt.options[0].innerHTML;

    selectSelectedDiv.innerHTML = labelText;

    node.appendChild(selectSelectedDiv);
    let b = document.createElement("DIV");
    b.setAttribute("class", "select-items select-hide");
    for (let j = 1; j < selElmntLength; j++) {
      const c = document.createElement("DIV");
      const optionValue = selElmnt.options[j].innerHTML;

      c.innerHTML = optionValue;

      c.addEventListener("click", function (e) {
        if (this.parentNode && this.parentNode.parentNode && this.parentNode.parentNode.querySelectorAll("select")[0]) {
          const y = this.parentNode.parentNode.querySelectorAll("select")[0] as HTMLSelectElement;
          const h = this.parentNode.previousSibling as HTMLElement;

          for (let k = 0; k < y.length; k++) {
            // FIXME **** (bad swear word) this **** (bad swear word) this should not be used never do anything like this its totally bad and buggy.

            y.selectedIndex = k;
            h.innerHTML = this.innerHTML;
            let yl = this.parentNode.querySelector(".same-as-selected") as HTMLSelectElement;
            if (yl) {
              for (let l = 0; l < yl.length; l++) {
                yl[l].removeAttribute("class");
              }
              this.setAttribute("class", "same-as-selected");
              break;
            }
          }
          h.click();
        }
      });
      b.appendChild(c);
    }
    node.appendChild(b);
    selectSelectedDiv.addEventListener("click", function (e) {
      e.stopPropagation();
      closeAllSelect(this);
      if (this.nextSibling) {
        const s = this.nextSibling as HTMLElement;
        s.classList.toggle("select-hide");
        this.classList.toggle("select-arrow-active");
      }
    });
  }
}

function closeAllSelect(elmnt: HTMLElement) {
  const arrNo = [];
  const selectItemsElements = document.getElementsByClassName("select-items") as HTMLCollectionOf<HTMLElement>;
  const selectSelectedElements = document.getElementsByClassName("select-selected") as HTMLCollectionOf<HTMLElement>;
  for (let selectSelectedElementsIndex = 0; selectSelectedElementsIndex < selectSelectedElements.length; selectSelectedElementsIndex++) {
    if (elmnt == selectSelectedElements[selectSelectedElementsIndex]) {
      arrNo.push(selectSelectedElementsIndex);
    } else {
      selectSelectedElements[selectSelectedElementsIndex].classList.remove("select-arrow-active");
    }
  }
  for (let selectItemsElementsIndex = 0; selectItemsElementsIndex < selectItemsElements.length; selectItemsElementsIndex++) {
    if (arrNo.indexOf(selectItemsElementsIndex)) {
      selectItemsElements[selectItemsElementsIndex].classList.add("select-hide");
    }
  }
}

document.addEventListener("click", function () {
  closeAllSelect(this.activeElement as HTMLElement);
});

function checkDropdownValue(windowType: "edit" | "create", dropdownType: "gender" | "ethnicity" | "religion" | "civilstatus") {
  if (dropdownType == "gender") {
    const selectedGender = document.querySelector<HTMLDivElement>("body > div." + windowType + "-container > div > div.scroll-box > div:nth-child(1) > div > div.select-selected")?.innerHTML ?? "";
    const gender: { [key: string]: string } = {};

    // English

    gender["Select gender:"] = "";
    gender["Male"] = "Male";
    gender["Female"] = "Female";
    gender["Other"] = "Other";

    return gender[selectedGender];
  } else if (dropdownType == "ethnicity") {
    const selectedEthnicity = document.querySelector<HTMLDivElement>("body > div." + windowType + "-container > div > div.scroll-box > div:nth-child(2) > div > div.select-selected")?.innerHTML ?? "";
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
    const selectedReligion = document.querySelector<HTMLDivElement>("body > div." + windowType + "-container > div > div.scroll-box > div:nth-child(15) > div > div.select-selected")!.innerHTML ?? "";
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
  } else if (dropdownType == "civilstatus") {
    const selectedCivilstatus = document.querySelector<HTMLDivElement>("body > div." + windowType + "-container > div > div.scroll-box > div:nth-child(7) > div > div.select-selected")?.innerHTML ?? "";
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
type DropdownType = "gender" | "ethnicity" | "religion" | "civilstatus" | "language";

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
  } else if (dropdownType == "civilstatus") {
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

export { saveAsFile, getDropdownElementIndex, checkDropdownValue, apiCall, DropdownType };
