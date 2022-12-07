import {jsonData} from "../data.js";

// Uncomment to test server
// If not commented, it wont work
// const response = await fetch('http://127.0.0.1:8080/persons', {
//   method: "GET",
//   mode: 'no-cors',
// });
// const myJson = await response.json();

// console.log(myJson);

let data = JSON.parse(jsonData);
var i;

const element = document.getElementById("searchbar");

element.addEventListener("keyup", search_users);

search_users();


function search_users() {
  let input = document.getElementById('searchbar').value
  input = input.toLowerCase();
  let x = document.querySelector('#list-holder');
  x.innerHTML = ""
  

  for (const [i, _] of Object.entries(data)) {
    let obj = data[i];

    if (obj.name.toLowerCase().includes(input)) {
      const base_div = document.createElement("div"); // Outer div
      base_div.className = "chip";

      const p_icon_div = document.createElement("div"); // Icon div
      p_icon_div.className = "chip-icon";

      const p_icon = document.createElement("ion-icon"); // Person icon
      p_icon.className = "icon"
      p_icon.setAttribute("name", "person");

      const txt_div = document.createElement("div"); // Text container
      txt_div.className = "text-container";

      const name_p = document.createElement("p"); // Name paragraph

      const v_icon_div = document.createElement("div"); // Icon div
      v_icon_div.className = "chip-view";

      const v_icon = document.createElement("ion-icon"); // View icon
      v_icon.className = "icon"
      v_icon.setAttribute("name", "eye-outline");

      const e_icon_div = document.createElement("div"); // Icon div
      e_icon_div.className = "chip-edit";

      const e_icon = document.createElement("ion-icon"); // Edit icon
      e_icon.className = "icon"
      e_icon.setAttribute("name", "create-outline");
      

      
      base_div.appendChild(p_icon_div);
      base_div.appendChild(txt_div);
      txt_div.appendChild(name_p);

      base_div.appendChild(v_icon_div);
      base_div.appendChild(e_icon_div);
      p_icon_div.appendChild(p_icon);
      v_icon_div.appendChild(v_icon);
      e_icon_div.appendChild(e_icon);



      name_p.innerHTML = `${obj.name}`
      x.appendChild(base_div)
    }
  }
}