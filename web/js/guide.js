// Create a new broadcast channel with the same name as in the first code block
const channel = new BroadcastChannel('dark-mode-channel');

// Listen for messages on the broadcast channel
channel.addEventListener('message', (event) => {
  if (event.data.type === 'dark-mode') {
    const isDarkMode = event.data.isDarkMode;
    localStorage.setItem('isDarkMode', isDarkMode);

    if (isDarkMode) {
      document.documentElement.setAttribute('data-theme', 'dark');
    } else {
      document.documentElement.setAttribute('data-theme', 'light');
    }
  }
});

var x, i, j, l, ll, selElmnt, a, b, c;
/*look for any elements with the class "country-select":*/
x = document.getElementsByClassName("country-select");
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













// The actual stuff

const countryDropdown = document.getElementById("country-select");
const selectedValue = document.querySelector(".country-select select").value;

const checkboxName = document.getElementById("checkbox_01");
const checkboxAddress = document.getElementById("checkbox_02");
const checkboxPhone = document.getElementById("checkbox_03");
const checkboxVIN = document.getElementById("checkbox_04");
const checkboxBusiness = document.getElementById("checkbox_05");
const checkboxIP = document.getElementById("checkbox_06");
const checkboxUsername = document.getElementById("checkbox_07");

const linkItems = document.querySelectorAll("link-list-holder li");


function filterAllItems() {
  console.log("filtering all items");
}


const list_elements = document.querySelectorAll(".link-list-holder li");

document.querySelector(".select-selected").addEventListener("DOMSubtreeModified", function() {
  const selectedCountry = document.querySelector(".select-selected").innerHTML;
  let currentCountry;

  for (let i = 0; i < list_elements.length; i++) {
    const element = list_elements[i];
    element.style.display = "flex";
  }

  if (selectedCountry !== "") {
    if (selectedCountry == "Select country:" || selectedCountry == "WorldWide") {
      currentCountry = "ww";
    } else if (selectedCountry == "United States of America") {
      currentCountry = "us";
    } else if (selectedCountry == "Canada") {
      currentCountry = "ca";
    } else if (selectedCountry == "United Kingdom") {
      currentCountry = "uk";
    }
    
    for (let i = 0; i < list_elements.length; i++) {
      const element = list_elements[i];

      if (!element.classList.contains(currentCountry)) {
        element.style.display = "none";
      }
    }
  }
});