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
    dropdownTitleElement.textContent = dropdownTitle;

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