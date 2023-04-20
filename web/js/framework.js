function delay(time) { // Because there is no default sleep function
  return new Promise(resolve => setTimeout(resolve, time));
}

function SaveAsFile(t, f, m) {
  // SaveAsFile("text","filename.txt","text/plain;charset=utf-8");

  try {
    var b = new Blob([t], {type: m});
    saveAs(b, f);
  } catch(e) {
    window.open("data:" + m + "," + encodeURIComponent(t), '_blank', '');
  }
}


const elements = {
  "gender-select": document.getElementsByClassName("gender-select"),
  "religion-select": document.getElementsByClassName("religion-select"),
  "civilstatus-select": document.getElementsByClassName("civilstatus-select")
};

for (const [className, nodeList] of Object.entries(elements)) {
  for (let i = 0; i < nodeList.length; i++) {
    const x = nodeList[i];
    const selElmnt = x.getElementsByTagName("select")[0];
    const ll = selElmnt.length;
    let a = document.createElement("DIV");
    a.setAttribute("class", "select-selected");
    a.innerHTML = selElmnt.options[selElmnt.selectedIndex].innerHTML;
    x.appendChild(a);
    let b = document.createElement("DIV");
    b.setAttribute("class", "select-items select-hide");
    for (let j = 1; j < ll; j++) {
      const c = document.createElement("DIV");
      c.innerHTML = selElmnt.options[j].innerHTML;
      c.addEventListener("click", function (e) {
        const y = this.parentNode.parentNode.getElementsByTagName("select")[0];
        const h = this.parentNode.previousSibling;
        for (let k = 0; k < y.length; k++) {
          if (y.options[k].innerHTML == this.innerHTML) {
            y.selectedIndex = k;
            h.innerHTML = this.innerHTML;
            let yl = this.parentNode.getElementsByClassName("same-as-selected");
            for (let l = 0; l < yl.length; l++) {
              yl[l].removeAttribute("class");
            }
            this.setAttribute("class", "same-as-selected");
            break;
          }
        }
        h.click();
      });
      b.appendChild(c);
    }
    x.appendChild(b);
    a.addEventListener("click", function (e) {
      e.stopPropagation();
      closeAllSelect(this);
      this.nextSibling.classList.toggle("select-hide");
      this.classList.toggle("select-arrow-active");
    });
  }
}

function closeAllSelect(elmnt) {
  const x = document.getElementsByClassName("select-items");
  const y = document.getElementsByClassName("select-selected");
  for (let i = 0; i < y.length; i++) {
    if (elmnt == y[i]) {
      continue;
    }
    y[i].classList.remove("select-arrow-active");
  }
  for (let i = 0; i < x.length; i++) {
    if (elmnt == y[i]) {
      continue;
    }
    x[i].classList.add("select-hide");
  }
}

document.addEventListener("click", closeAllSelect);

document.querySelectorAll("span").forEach(function (element) {
  element.addEventListener('paste', function (e) {
    // Prevent the default action
    e.preventDefault();

    // Get the copied text from the clipboard
    const text = e.clipboardData
      ? (e.originalEvent || e).clipboardData.getData('text/plain')
      : // For IE
      window.clipboardData
        ? window.clipboardData.getData('Text')
        : '';

    if (document.queryCommandSupported('insertText')) {
      document.execCommand('insertText', false, text);
    } else {
      // Insert text at the current position of caret
      const range = document.getSelection().getRangeAt(0);
      range.deleteContents();

      const textNode = document.createTextNode(text);
      range.insertNode(textNode);
      range.selectNodeContents(textNode);
      range.collapse(false);

      const selection = window.getSelection();
      selection.removeAllRanges();
      selection.addRange(range);
    }
  });
});

[document.getElementById("c-name-tag"), document.getElementById("acc-name-tag")].forEach(item => {
  item.addEventListener('paste', function (e) {
    // Prevent the default action
    e.preventDefault();

    // Get the copied text from the clipboard
    const text = e.clipboardData
      ? (e.originalEvent || e).clipboardData.getData('text/plain')
      : // For IE
      window.clipboardData
        ? window.clipboardData.getData('Text')
        : '';

    if (document.queryCommandSupported('insertText')) {
      document.execCommand('insertText', false, text);
    } else {
      // Insert text at the current position of caret
      const range = document.getSelection().getRangeAt(0);
      range.deleteContents();

      const textNode = document.createTextNode(text);
      range.insertNode(textNode);
      range.selectNodeContents(textNode);
      range.collapse(false);

      const selection = window.getSelection();
      selection.removeAllRanges();
      selection.addRange(range);
    }
  });
});

function checkGender(type) {
  const selectedGender = document.querySelector("body > div." + type + "-container > div > div.scroll-box > div:nth-child(1) > div > div.select-selected").innerHTML;
  const gender = new Map();

  gender["Select gender:"] = "";
  gender["Male"] = "Male";
  gender["Female"] = "Female";
  gender["Other"] = "Other";

  return gender[selectedGender];
}

function getGenderElementIndex(gender) {
  const genderIndex = new Map();

  genderIndex[""] = "";
  genderIndex["Male"] = "0";
  genderIndex["Female"] = "1";
  genderIndex["Other"] = "2";

  return genderIndex[gender];
}

function checkCivilstatus(type) {
  const selectedCivilstatus = document.querySelector("body > div." + type + "-container > div > div.scroll-box > div:nth-child(6) > div > div.select-selected").innerHTML;
  const civilstatus = new Map();

  civilstatus["Select civil status:"] = "";
  civilstatus["Single"] = "Single";
  civilstatus["Married"] = "Married";
  civilstatus["Widowed"] = "Widowed";
  civilstatus["Divorced"] = "Divorced";
  civilstatus["Separated"] = "Separated";

  return civilstatus[selectedCivilstatus];
}

function getCivilstatusElementIndex(civilstatus) {
  const civilstatusIndex = new Map();

  civilstatusIndex[""] = "";
  civilstatusIndex["Single"] = "0";
  civilstatusIndex["Married"] = "1";
  civilstatusIndex["Widowed"] = "2";
  civilstatusIndex["Divorced"] = "3";
  civilstatusIndex["Separated"] = "4";

  return civilstatusIndex[civilstatus];
}


function checkReligion(type) {
  const selectedReligion = document.querySelector("body > div." + type + "-container > div > div.scroll-box > div:nth-child(13) > div > div.select-selected").innerHTML;
  const religion = new Map();

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
}

function getReligionElementIndex(religion) {
  const religionIndex = new Map();

  religionIndex[""] = "";
  religionIndex["Christianity"] = "0";
  religionIndex["Atheism"] = "1";
  religionIndex["Islam"] = "2";
  religionIndex["Hinduism"] = "3";
  religionIndex["Buddhism"] = "4";
  religionIndex["Sikhism"] = "5";
  religionIndex["Judaism"] = "6";
  religionIndex["Other"] = "7";

  return religionIndex[religion];
}


export {delay, SaveAsFile, checkGender, getGenderElementIndex, checkReligion, getReligionElementIndex, checkCivilstatus, getCivilstatusElementIndex};