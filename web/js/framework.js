function delay(time) { // Because there is no default sleep function
  return new Promise(resolve => setTimeout(resolve, time));
}

function SaveAsFile(t, f, m) {
  // SaveAsFile("text","filename.txt","text/plain;charset=utf-8");

  try {
    var b = new Blob([t],{type:m});
    saveAs(b, f);
  } catch (e) {
    window.open("data:"+m+"," + encodeURIComponent(t), '_blank','');
  }
}


var x, i, j, l, ll, selElmnt, a, b, c;
/*look for any elements with the class "gender-select":*/
x = document.getElementsByClassName("gender-select");
l = x.length;
for (i = 0; i < l; i++) {
  selElmnt = x[i].getElementsByTagName("select")[0];
  ll = selElmnt.length;
  /*for each element, create a new DIV that will act as the selected item:*/
  a = document.createElement("DIV");
  a.setAttribute("class", "select-selected");
  a.innerHTML = selElmnt.options[selElmnt.selectedIndex].innerHTML;
  x[i].appendChild(a);
  /*for each element, create a new DIV that will contain the option list:*/
  b = document.createElement("DIV");
  b.setAttribute("class", "select-items select-hide");
  for (j = 1; j < ll; j++) {
    /*for each option in the original select element,
    create a new DIV that will act as an option item:*/
    c = document.createElement("DIV");
    c.innerHTML = selElmnt.options[j].innerHTML;
    c.addEventListener("click", function(e) {
        /*when an item is clicked, update the original select box,
        and the selected item:*/
        var y, i, k, s, h, sl, yl;
        s = this.parentNode.parentNode.getElementsByTagName("select")[0];
        sl = s.length;
        h = this.parentNode.previousSibling;
        for (i = 0; i < sl; i++) {
          if (s.options[i].innerHTML == this.innerHTML) {
            s.selectedIndex = i;
            h.innerHTML = this.innerHTML;
            y = this.parentNode.getElementsByClassName("same-as-selected");
            yl = y.length;
            for (k = 0; k < yl; k++) {
              y[k].removeAttribute("class");
            }
            this.setAttribute("class", "same-as-selected");
            break;
          }
        }
        h.click();
    });
    b.appendChild(c);
  }
  x[i].appendChild(b);
  a.addEventListener("click", function(e) {
      /*when the select box is clicked, close any other select boxes,
      and open/close the current select box:*/
      e.stopPropagation();
      closeAllSelect(this);
      this.nextSibling.classList.toggle("select-hide");
      this.classList.toggle("select-arrow-active");
    });
}
function closeAllSelect(elmnt) {
  /*a function that will close all select boxes in the document,
  except the current select box:*/
  var x, y, i, xl, yl, arrNo = [];
  x = document.getElementsByClassName("select-items");
  y = document.getElementsByClassName("select-selected");
  xl = x.length;
  yl = y.length;
  for (i = 0; i < yl; i++) {
    if (elmnt == y[i]) {
      arrNo.push(i)
    } else {
      y[i].classList.remove("select-arrow-active");
    }
  }
  for (i = 0; i < xl; i++) {
    if (arrNo.indexOf(i)) {
      x[i].classList.add("select-hide");
    }
  }
}
/*if the user clicks anywhere outside the select box,
then close all select boxes:*/
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

function checkGender() {
  const selectedGender = document.querySelector("body > div.edit-container > div > div.scroll-box > div:nth-child(1) > div > div.select-selected").innerHTML;
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


export {delay, SaveAsFile, checkGender, getGenderElementIndex};