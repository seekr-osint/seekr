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


export {delay, SaveAsFile, checkGender, getGenderElementIndex, checkReligion, getReligionElementIndex, checkCivilstatus, getCivilstatusElementIndex};