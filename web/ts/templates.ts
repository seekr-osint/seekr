class CustomDropdown extends HTMLElement {
  constructor() {
    super();

    // Create shadow DOM and attach template
    const template = document.getElementById("dropdown-template") as HTMLTemplateElement;
    const shadowRoot = this.attachShadow({ mode: "open" });
    shadowRoot.appendChild(template.content.cloneNode(true));

    // Attach .css
    const linkElem = document.createElement("link");
    linkElem.setAttribute("rel", "stylesheet");
    linkElem.setAttribute("href", "./css/style.css");
    shadowRoot.appendChild(linkElem);

    // Get elements from shadow DOM
    const node = this.shadowRoot!.querySelector(".dropdown-select")!;
    const dropdownTitleElement = this.shadowRoot!.querySelector(".tag")!;
    const selElmnt = node.querySelector("select") as HTMLSelectElement;

    // Get dropdown title attribute and set initial values
    const dropdownTitleAttr = this.getAttribute("title")!;
    const dropdownTitle = dropdownTitleAttr.charAt(0).toUpperCase() + dropdownTitleAttr.slice(1) + ":";

    if (this.getAttribute("no-tag") != "") {
      dropdownTitleElement.textContent = dropdownTitle;
    }

    // Set initial option for select element
    const selElmntTag = selElmnt.children[0] as HTMLOptionElement;
    const selElmntTagText = "Select " + dropdownTitleAttr + ":";
    selElmntTag.textContent = selElmntTagText;

    // Create and append option elements to select element
    const options = Array.from(this.querySelectorAll("custom-option"));
    options.forEach(option => {
      const optionValue = option.innerHTML;
      const optionElement = document.createElement("option");
      optionElement.innerHTML = optionValue!;
      selElmnt.appendChild(optionElement);
    });

    // Create div elements for dropdown
    const selectSelectedDiv = document.createElement("div");
    selectSelectedDiv.className = "select-selected";
    const labelText = selElmnt.options[0].innerHTML;
    selectSelectedDiv.innerHTML = labelText;

    const b = document.createElement("div");
    b.className = "select-items select-hide";

    // Create option elements for dropdown
    for (let j = 1; j < selElmnt.length; j++) {
      const c = document.createElement("div");
      const optionValue = selElmnt.options[j].innerHTML;
      c.innerHTML = optionValue;

      // Add click event listener for dropdown options
      c.addEventListener("click", function () {
        const y = this.parentNode?.parentNode?.querySelector("select") as HTMLSelectElement;
        const h = this.parentNode?.previousSibling as HTMLElement;

        for (let k = 0; k < y.length; k++) {
          y.selectedIndex = k;
          h.innerHTML = this.innerHTML;
          const yl = this.parentNode?.querySelector(".same-as-selected") as HTMLElement | null;

          if (yl) {
            yl.classList.remove("same-as-selected");
            this.classList.add("same-as-selected");
            break;
          }
        }
        h.click();
      });

      b.appendChild(c);
    }

    // Append elements to dropdown
    node.appendChild(selectSelectedDiv);
    node.appendChild(b);

    // Add click event listener for showing/hiding dropdown options
    selectSelectedDiv.addEventListener("click", function (e) {
      e.stopPropagation();
      closeAllSelect(this);
      const s = this.nextSibling as HTMLElement;
      s.classList.toggle("select-hide");
      this.classList.toggle("select-arrow-active");
    });


    if (document.location.pathname == "/web/guide.html") {
      const linkElemGuide = document.createElement("link");
      linkElemGuide.setAttribute("rel", "stylesheet");
      linkElemGuide.setAttribute("href", "./css/guide.css");
      shadowRoot.appendChild(linkElemGuide);

      function preListHandler() {
        if (selectSelectedDiv && selectSelectedDiv.innerHTML !== "") {
          listHandler();
        }
      }

      // Create a MutationObserver instance
      const observer = new MutationObserver(preListHandler);

      // Define the configuration for the observer
      const config = { childList: true, subtree: true };

      // Start observing the parent node
      if (selectSelectedDiv.parentElement) {
        observer.observe(selectSelectedDiv.parentElement, config);
      }

      const checkboxName = document.getElementById("checkbox_01") as HTMLInputElement;
      const checkboxNameIcon = document.getElementById("checkbox_01_icon") as HTMLElement;

      const checkboxAddress = document.getElementById("checkbox_02") as HTMLInputElement;
      const checkboxAddressIcon = document.getElementById("checkbox_02_icon") as HTMLElement;

      const checkboxPhone = document.getElementById("checkbox_03") as HTMLInputElement;
      const checkboxPhoneIcon = document.getElementById("checkbox_03_icon") as HTMLElement;

      const checkboxVIN = document.getElementById("checkbox_04") as HTMLInputElement;
      const checkboxVINIcon = document.getElementById("checkbox_04_icon") as HTMLElement;

      const checkboxBusiness = document.getElementById("checkbox_05") as HTMLInputElement;
      const checkboxBusinessIcon = document.getElementById("checkbox_05_icon") as HTMLElement;

      const checkboxIP = document.getElementById("checkbox_06") as HTMLInputElement;
      const checkboxIPIcon = document.getElementById("checkbox_06_icon") as HTMLElement;

      const checkboxUsername = document.getElementById("checkbox_07") as HTMLInputElement;
      const checkboxUsernameIcon = document.getElementById("checkbox_07_icon") as HTMLElement;

      const checkboxDomain = document.getElementById("checkbox_08") as HTMLInputElement;
      const checkboxDomainIcon = document.getElementById("checkbox_08_icon") as HTMLElement;

      const list_elements = document.querySelectorAll(".link-list-holder li");


      function resetAll() {
        for (let i = 0; i < list_elements.length; i++) {
          const element = list_elements[i] as HTMLElement;
          element.style.display = "flex";
        }
      }

      function checkChecboxValue(checkboxType: string): boolean {
        if (checkboxType == "name") {
          return checkboxName.checked;
        } else if (checkboxType == "address") {
          return checkboxAddress.checked;
        } else if (checkboxType == "phone") {
          return checkboxPhone.checked;
        } else if (checkboxType == "vin") {
          return checkboxVIN.checked;
        } else if (checkboxType == "business") {
          return checkboxBusiness.checked;
        } else if (checkboxType == "ip") {
          return checkboxIP.checked;
        } else if (checkboxType == "username") {
          return checkboxUsername.checked;
        } else if (checkboxType == "domain") {
          return checkboxDomain.checked;
        } else {
          return false;
        }
      }

      function checkCountry(): string | undefined {
        if (document) {
          if (selectSelectedDiv) {
            const countries: { [key: string]: string } = {};

            // English

            countries["Select country:"] = "all";
            countries["WorldWide"] = "ww";
            countries["United States"] = "us";
            countries["Canada"] = "ca";
            countries["United Kingdom"] = "uk";
            countries["Sweden"] = "se";
            countries["Germany"] = "de";
            countries["France"] = "fr";
            countries["Italy"] = "it";
            countries["Russia"] = "ru";
            countries["Australia"] = "au";

            return countries[selectSelectedDiv.innerHTML]; // Error here
          }
        }
      }

      function listHandler() {
        type checkboxResultType = string | boolean;

        let listOfClasses: checkboxResultType[] = ["country", "name", "address", "phone", "vin", "business", "ip", "username", "domain"];

        resetAll();

        const selectedCountry = checkCountry();

        // Replace the first element with the country code
        if (selectedCountry != undefined) {
          listOfClasses[0] = selectedCountry;
        }

        if (checkChecboxValue("name") == false) {
          listOfClasses[1] = false;

          checkboxNameIcon.style.opacity = "0";
        } else {
          checkboxNameIcon.style.opacity = "1";
        }

        if (checkChecboxValue("address") == false) {
          listOfClasses[2] = false;    
          
          checkboxAddressIcon.style.opacity = "0";
        } else {
          checkboxAddressIcon.style.opacity = "1";
        }

        if (checkChecboxValue("phone") == false) {
          listOfClasses[3] = false;

          checkboxPhoneIcon.style.opacity = "0";
        } else {
          checkboxPhoneIcon.style.opacity = "1";
        }
        
        if (checkChecboxValue("vin") == false) {
          listOfClasses[4] = false;
          
          checkboxVINIcon.style.opacity = "0";
        } else {
          checkboxVINIcon.style.opacity = "1";
        }

        if (checkChecboxValue("business") == false) {
          listOfClasses[5] = false;

          checkboxBusinessIcon.style.opacity = "0";
        } else {
          checkboxBusinessIcon.style.opacity = "1";
        }

        if (checkChecboxValue("ip") == false) {
          listOfClasses[6] = false;

          checkboxIPIcon.style.opacity = "0";
        } else {
          checkboxIPIcon.style.opacity = "1";
        }

        if (checkChecboxValue("username") == false) {
          listOfClasses[7] = false;

          checkboxUsernameIcon.style.opacity = "0";
        } else {
          checkboxUsernameIcon.style.opacity = "1";
        }

        if (checkChecboxValue("domain") == false) {
          listOfClasses[8] = false;
          
          checkboxDomainIcon.style.opacity = "0";
        } else {
          checkboxDomainIcon.style.opacity = "1";
        }

        if (listOfClasses[1] == false && listOfClasses[2] == false && listOfClasses[3] == false && listOfClasses[4] == false && listOfClasses[5] == false && listOfClasses[6] == false && listOfClasses[7] == false && listOfClasses[8] == false) {
          for (let i = 0; i < list_elements.length; i++) {
            const element = list_elements[i] as HTMLElement;
        
            if (listOfClasses[0] == "all") {
              element.style.display = "flex";
            } else if (!element.classList.contains(listOfClasses[0].toString())) {
              element.style.display = "none";
            }
          }
        } else {
          for (let i = 0; i < list_elements.length; i++) {
            const element = list_elements[i] as HTMLElement;
            
            if (!element.classList.contains(listOfClasses[0].toString()) && listOfClasses[0] != "all") {
              element.style.display = "none";
            } else {
              let hasBeenChanged = false;

              for (let i = 1; i < listOfClasses.length; i++) {
                if (element.classList.contains(listOfClasses[i].toString())) {
                  element.style.display = "flex";
                  hasBeenChanged = true;
                }

                if (i == 8 && hasBeenChanged == false) {
                  element.style.display = "none";
                }
              }
            }
          }
        }
      }



      // Rest of your code for checkbox event listeners
      checkboxName.addEventListener('change', preListHandler);
      checkboxAddress.addEventListener('change', preListHandler);
      checkboxPhone.addEventListener('change', preListHandler);
      checkboxVIN.addEventListener('change', preListHandler);
      checkboxBusiness.addEventListener('change', preListHandler);
      checkboxIP.addEventListener('change', preListHandler);
      checkboxUsername.addEventListener('change', preListHandler);
      checkboxDomain.addEventListener('change', preListHandler);
    }
  }
}

customElements.define("custom-dropdown", CustomDropdown);

// Global event listener
document.addEventListener("click", function () {
  closeAllSelect(this.activeElement as HTMLElement);
});

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