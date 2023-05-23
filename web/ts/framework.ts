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

const elements: { [key: string]: HTMLCollectionOf<Element> } = {
  "gender-select": document.getElementsByClassName("gender-select"),
  "religion-select": document.getElementsByClassName("religion-select"),
  "civilstatus-select": document.getElementsByClassName("civilstatus-select"),
  "country-select": document.getElementsByClassName("country-select"),
  "language-select": document.getElementsByClassName("language-select")
};

for (const [className, nodeList] of Object.entries(elements)) {
  for (let i = 0; i < nodeList.length; i++) {
    const x = nodeList[i] as HTMLElement;
    const selElmnt = x.getElementsByTagName("select")[0] as HTMLSelectElement;
    const ll = selElmnt.length;
    let a = document.createElement("DIV");
    a.setAttribute("class", "select-selected");

    // Modify the label text and lng-tag attribute
    const labelText = selElmnt.options[0].innerHTML;
    const lngTagValue = labelText
      .toLowerCase()
      .replace(/\s/g, "_")
      .replace(":", "_colon");
    a.setAttribute("lng-tag", lngTagValue);

    a.innerHTML = labelText;
    x.appendChild(a);
    let b = document.createElement("DIV");
    b.setAttribute("class", "select-items select-hide");
    for (let j = 1; j < ll; j++) {
      const c = document.createElement("DIV");
      const optionValue = selElmnt.options[j].innerHTML;
      const lngTagValue = optionValue.toLowerCase().replace(/\s/g, "_");
      c.setAttribute("lng-tag", lngTagValue);
      c.innerHTML = optionValue;
      c.addEventListener("click", function (e) {
        if (
          this.parentNode &&
          this.parentNode.parentNode &&
          this.parentNode.parentNode.querySelectorAll("select")[0]
        ) {
          const y = this.parentNode.parentNode.querySelectorAll(
            "select"
          )[0] as HTMLSelectElement;
          const h = this.parentNode.previousSibling as HTMLElement;

          for (let k = 0; k < y.length; k++) {
            if (y.options[k].innerHTML == this.innerHTML) {
              y.selectedIndex = k;
              h.innerHTML = this.innerHTML;
              let yl = this.parentNode.querySelector(
                ".same-as-selected"
              ) as HTMLSelectElement;
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
    x.appendChild(b);
    a.addEventListener("click", function (e) {
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
  const x = document.getElementsByClassName("select-items") as HTMLCollectionOf<HTMLElement>;
  const y = document.getElementsByClassName("select-selected") as HTMLCollectionOf<HTMLElement>;
  for (let i = 0; i < y.length; i++) {
    if (elmnt == y[i]) {
      arrNo.push(i);
    } else {
      y[i].classList.remove("select-arrow-active");
    }
  }
  for (let i = 0; i < x.length; i++) {
    if (arrNo.indexOf(i)) {
      x[i].classList.add("select-hide");
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

    // German

    gender["Wähle Geschlecht:"] = "";
    gender["Männlich"] = "Male";
    gender["Weiblich"] = "Female";
    gender["Sonstiges"] = "Other";
  
    return gender[selectedGender];
  } else if (dropdownType == "religion") {
    const selectedReligion = document.querySelector<HTMLDivElement>("body > div." + windowType + "-container > div > div.scroll-box > div:nth-child(14) > div > div.select-selected")?.innerHTML ?? "";
    const religion: { [key: string]: string } = {};
  
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
    const selectedCivilstatus = document.querySelector<HTMLDivElement>("body > div." + windowType + "-container > div > div.scroll-box > div:nth-child(6) > div > div.select-selected")?.innerHTML ?? "";
    const civilstatus: { [key: string]: string } = {};
  
    civilstatus["Select civil status:"] = "";
    civilstatus["Single"] = "Single";
    civilstatus["Married"] = "Married";
    civilstatus["Widowed"] = "Widowed";
    civilstatus["Divorced"] = "Divorced";
    civilstatus["Separated"] = "Separated";
  
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

    // German

    genderIndex["Männlich"] = "0";
    genderIndex["Weiblich"] = "1";
    genderIndex["Sonstiges"] = "2";
  
    return genderIndex[dropdownValue];
  } else if (dropdownType == "religion") {
    const religionIndex: { [key: string]: string } = {};

    religionIndex[""] = "";
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
    civilstatusIndex["Single"] = "0";
    civilstatusIndex["Married"] = "1";
    civilstatusIndex["Widowed"] = "2";
    civilstatusIndex["Divorced"] = "3";
    civilstatusIndex["Separated"] = "4";
  
    return civilstatusIndex[dropdownValue];
  }
}


export {saveAsFile, getDropdownElementIndex, checkDropdownValue};