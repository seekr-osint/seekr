import {jsonData} from "./data.js";

console.log(jsonData)

let data = JSON.parse(jsonData);
var i;

const element = document.getElementById("searchbar");

element.addEventListener("keyup", search_animal);

function search_animal() {
  let input = document.getElementById('searchbar').value
  input = input.toLowerCase();
  let x = document.querySelector('#list-holder');
  x.innerHTML = ""
  

  for (const [i, _] of Object.entries(data)) {
    let obj = data[i];

    if (obj.name.toLowerCase().includes(input)) {
      const elem = document.createElement("li")
      elem.innerHTML = `${obj.name} - ${obj.id} - ${obj.age}`
      x.appendChild(elem)
    }
  }
}