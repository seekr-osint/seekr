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

      ${id}:
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
const button = document.querySelector("#myButton");
 
button!.addEventListener("click", function() {
   console.log(getValue("test1"));
 });

const d = getDropdown("test1")

d!.addEventListener("change", function() {
  console.log(getValue("test1"));
});

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
