const dropdowns: Set<string> = new Set([]);


class CustomOption extends HTMLElement {
  constructor() {
    super();
    const id = this.textContent || ""
    if (id == "") {
      console.log("error no dropdown name")
      return
    }
    this.setAttribute('id', id);
    dropdowns.add(id);

    const options = this.getAttribute('options') || "";
    if (options == "") {
      console.log("error empty dropdown options")
      return
    }

    const placeholder = this.getAttribute('placeholder') || "Select Item";
    this.attachShadow({ mode: 'open' });

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
        dropdown.classList.toggle("active");
      };
    }

    const internalDivs = this.shadowRoot!.querySelectorAll('.dropdown-option');
    internalDivs.forEach((internalDiv) => {
      internalDiv.addEventListener('click', () => {
        const optionName = internalDiv.textContent || '';
        const textBox = this.shadowRoot!.querySelector(".text-box") as HTMLInputElement;
        if (textBox) {
          textBox.value = optionName;                  
        }
        this.setAttribute('value', optionName);
      });
    });
  }
}

customElements.define('custom-dropdown', CustomOption);

function divTemplate(words: string): string {
  const wordArray = words.split(',').map(word => word.trim());

  let output = '';

  wordArray.forEach(word => {
    output += `<div class="dropdown-option">${word}</div>`;

  });

  return output;
}

const button = document.getElementById('myButton');

button!.addEventListener('click', function() {
  console.log(getValue("test1"));
});

function getValue(id: string): string {
  if (!dropdowns.has(id)) {
    console.log("Error: no dropdown with this name")
  }
  const customOptionElement = document.querySelector(`custom-dropdown#${id}`);
  
  return customOptionElement!.getAttribute("value") || ""
}


export { getValue };