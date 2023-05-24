declare function saveAs(blob: Blob, filename: string): void;

function saveAsFile(textContent: string, fileName: string) {
  // saveAsFile("text","filename.txt");
  const textEncoding: string = "text/plain;charset=utf-8";

  try {
    var blob = new Blob([textContent], {type: textEncoding});
    saveAs(blob, fileName);
  } catch(exception) {
    window.open("data:" + textEncoding + "," + encodeURIComponent(textContent), '_blank', '');
  }
}

const elements: { [key: string]: HTMLCollectionOf<Element> } = { // used to select a element in the dropdown
  "gender-select": document.getElementsByClassName("gender-select"),
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

    // Modify the label text and lng-tag attribute
    const labelText = selElmnt.options[0].innerHTML;
    const lngTagValue = labelText
      .toLowerCase()
      .replace(/\s/g, "_")
      .replace(":", "_colon");
    selectSelectedDiv.setAttribute("lng-tag", lngTagValue);

    selectSelectedDiv.innerHTML = labelText;

    translateElement(selectSelectedDiv);

    node.appendChild(selectSelectedDiv);
    let b = document.createElement("DIV");
    b.setAttribute("class", "select-items select-hide");
    for (let j = 1; j < selElmntLength; j++) {
      const c = document.createElement("DIV");
      const optionValue = selElmnt.options[j].innerHTML;
      const lngTagValue = optionValue.toLowerCase().replace(/\s/g, "_");
      c.setAttribute("lng-tag", lngTagValue);

      c.innerHTML = optionValue;
      
      translateElement(c);

      c.addEventListener("click", function (e) {
        if (this.parentNode && this.parentNode.parentNode && this.parentNode.parentNode.querySelectorAll("select")[0]) {
          const y = this.parentNode.parentNode.querySelectorAll("select")[0] as HTMLSelectElement;
          const h = this.parentNode.previousSibling as HTMLElement;

          for (let k = 0; k < y.length; k++) {
            //console.log(this.getAttribute("lng-tag").charAt(0).toUpperCase() + string.slice(1))
            // FIXME **** (bad swear word) this **** (bad swear word) this should not be used never do anything like this its totally bad and buggy.
            // value is basically .innerHTML translated to English
            let value = this.getAttribute("lng-tag")!.charAt(0).toUpperCase() + this.getAttribute("lng-tag")!.slice(1)
            if (y.options[k].innerHTML == value!) { //translateText(this.innerHTML)) {
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

document.addEventListener("click", function() {
  closeAllSelect(this.activeElement as HTMLElement);
});

function checkDropdownValue(windowType: "edit" | "create", dropdownType: "gender" | "religion" | "civilstatus") {
  if (dropdownType == "gender") {
    const selectedGender = document.querySelector<HTMLDivElement>("body > div." + windowType + "-container > div > div.scroll-box > div:nth-child(1) > div > div.select-selected")?.innerHTML ?? "";
    const gender: { [key: string]: string } = {};
  
    // English

    gender["Select gender:"] = "";
    gender["Male"] = "Male";
    gender["Female"] = "Female";
    gender["Other"] = "Other";

    // Translations

    if (gender[selectedGender] == undefined) {
      gender[translateText("select_gender_colon")!] = "";
      gender[translateText("male")!] = "Male";
      gender[translateText("female")!] = "Female";
      gender[translateText("other")!] = "Other";
    }

    return gender[selectedGender];
  } else if (dropdownType == "religion") {
    const selectedReligion = document.querySelector<HTMLDivElement>("body > div." + windowType + "-container > div > div.scroll-box > div:nth-child(14) > div > div.select-selected")?.innerHTML ?? "";
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

    // Translations

    if (religion[selectedReligion] == undefined) {
      religion[translateText("select_religion_colon")!] = "";
      religion[translateText("christianity")!] = "Christianity";
      religion[translateText("atheism")!] = "Atheism";
      religion[translateText("islam")!] = "Islam";
      religion[translateText("hinduism")!] = "Hinduism";
      religion[translateText("buddhism")!] = "Buddhism";
      religion[translateText("sikhism")!] = "Sikhism";
      religion[translateText("judaism")!] = "Judaism";
      religion[translateText("other")!] = "Other";
    }
  
    return religion[selectedReligion];
  } else if (dropdownType == "civilstatus") {
    const selectedCivilstatus = document.querySelector<HTMLDivElement>("body > div." + windowType + "-container > div > div.scroll-box > div:nth-child(6) > div > div.select-selected")?.innerHTML ?? "";
    const civilstatus: { [key: string]: string } = {};

    // English
  
    civilstatus["Select civil status:"] = "";
    civilstatus["Single"] = "Single";
    civilstatus["Married"] = "Married";
    civilstatus["Widowed"] = "Widowed";
    civilstatus["Divorced"] = "Divorced";
    civilstatus["Separated"] = "Separated";

    // Translations

    if (civilstatus[selectedCivilstatus] == undefined) {
      civilstatus[translateText("select_civil_status_colon")!] = "";
      civilstatus[translateText("single")!] = "Single";
      civilstatus[translateText("married")!] = "Married";
      civilstatus[translateText("widowed")!] = "Widowed";
      civilstatus[translateText("divorced")!] = "Divorced";
      civilstatus[translateText("separated")!] = "Separated";  
    }
  
    return civilstatus[selectedCivilstatus];
  }
}

function getDropdownElementIndex(dropdownType: "gender" | "religion" | "civilstatus", dropdownValue: string): string | undefined {
  if (dropdownType == "gender") {
    const genderIndex: { [key: string]: string } = {};

    genderIndex[""] = "";

    // English

    genderIndex["Male"] = "0";
    genderIndex["Female"] = "1";
    genderIndex["Other"] = "2";

    // Translation

    if (genderIndex[dropdownValue] == undefined) {
      genderIndex[translateText("male")!] = "0";
      genderIndex[translateText("female")!] = "1";
      genderIndex[translateText("other")!] = "2";
    }
  
    return genderIndex[dropdownValue];
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
    
    // Translation

    if (religionIndex[dropdownValue] == undefined) {
      religionIndex[translateText("christianity")!] = "0";
      religionIndex[translateText("atheism")!] = "1";
      religionIndex[translateText("islam")!] = "2";
      religionIndex[translateText("hinduism")!] = "3";
      religionIndex[translateText("buddhism")!] = "4";
      religionIndex[translateText("sikhism")!] = "5";
      religionIndex[translateText("judaism")!] = "6";
      religionIndex[translateText("other")!] = "7";
    }

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

    // Translation

    if (civilstatusIndex[dropdownValue] == undefined) {
      civilstatusIndex[translateText("single")!] = "0";
      civilstatusIndex[translateText("married")!] = "1";
      civilstatusIndex[translateText("widowed")!] = "2";
      civilstatusIndex[translateText("divorced")!] = "3";
      civilstatusIndex[translateText("separated")!] = "4";
    }
  
    return civilstatusIndex[dropdownValue];
  }
}


export {saveAsFile, getDropdownElementIndex, checkDropdownValue};
