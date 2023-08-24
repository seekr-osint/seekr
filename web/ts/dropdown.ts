const dropdowns: Set<string> = new Set([]);


class CustomDropdown extends HTMLElement {
  constructor() {
    super();
    const id = this.getAttribute("dropdown-id") || ""
    if (id == "") {
      console.error("error no dropdown name")
      return
    } else if (dropdowns.has(id)) {
      console.error("ID" + id + "already exsists");
      return
    }
    dropdowns.add(id);

    const options = this.getAttribute("options") || "";
    if (options == "") {
      console.error("error empty dropdown options")
      dropdowns.delete(id)
      return
    }

    const placeholder = this.getAttribute("placeholder") || "Select Item";
    this.attachShadow({ mode: "open" });

    this.shadowRoot!.innerHTML = `
      <link rel="stylesheet" href="./css/style.css">
      <link rel="stylesheet" href="./css/dropdown.css">

      <div class="dropdown">
        <input class="text-box" type="text" placeholder="${placeholder}" readonly>
        <div class="options">
          ${divTemplate(options)}
        </div>
      </div>
    `;

    const dropdown = this.shadowRoot!.querySelector(".dropdown") as HTMLElement;
    if (dropdown) {
      dropdown.onclick = function() {
      const allDropdowns = document.querySelectorAll("custom-dropdown");
      console.log(allDropdowns);

      allDropdowns.forEach((currentDropdown) => {
      const shadowRoot = currentDropdown.shadowRoot;
      if (shadowRoot) {
        const dropdownElement = shadowRoot.querySelector(".dropdown");
        if (dropdownElement && dropdownElement != dropdown) { // don't remove on teh dropdown the user clicked on
          dropdownElement.classList.remove("active");
        }
      }
      });
        dropdown.classList.toggle("active");
      };
    }

    const internalDivs = this.shadowRoot!.querySelectorAll(".dropdown-option");
    internalDivs.forEach((internalDiv) => {
      internalDiv.addEventListener("click", () => {
        const optionName = internalDiv.textContent || "";
        const textBox = this.shadowRoot!.querySelector(".text-box") as HTMLInputElement;
        if (textBox) {
          textBox.value = optionName;                  
        }
        this.setAttribute("value", optionName);
        this.dispatchEvent(new Event("change"));
      });
    });
  }
}

customElements.define("custom-dropdown", CustomDropdown);

function divTemplate(words: string): string {
  const wordArray = words.split(",").map(word => word.trim());

  let output = "";

  wordArray.forEach(word => {
    output += `<div class="dropdown-option">${word}</div>`;

  });

  return output;
}

function getDropdown(id: string): Element | null {
  if (!dropdowns.has(id)) {
    console.error("Error: no dropdown with this name: " + id)
    return null
  }
  return document.querySelector(`custom-dropdown[dropdown-id="${id}"]`);
}

function getValue(id: string): string {
  if (!dropdowns.has(id)) {
    console.error("Error: no dropdown with this name: " + id)
    return ""
  }
  const customOptionElement = getDropdown(id);
  
  return customOptionElement!.getAttribute("value") || ""
}


export { getValue, getDropdown };
