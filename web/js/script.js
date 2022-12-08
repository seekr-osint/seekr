import {jsonData} from "./data.js";

async function foo() {
  let obj;

  const res = await fetch("http://localhost:8080/persons")

  obj = await res.json();

  // console.log(obj)
  return obj;
}



// fetch("http://localhost:8080/persons").then(function(response) {
//   return response.json();
// }).then(function(data) {
//   console.log(data);
// }).catch(function(err) {
//   console.log('Fetch Error :-S', err);
// });

// foo()
console.log(foo());

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

function delay(time) { // Because there is no default sleep function
  return new Promise(resolve => setTimeout(resolve, time));
}

function search_users() {
  let input = document.getElementById('searchbar').value;
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
      name_p.className = "card-text";

      const v_icon_div = document.createElement("div"); // Icon div
      v_icon_div.className = "chip-view";
      
      v_icon_div.onclick = function() {
        document.querySelector('.main').style.display = "none";
        document.querySelector('.container').style.display = "flex";

        document.querySelector(".name-tag").innerHTML = obj.name;

        document.querySelector(".age").innerHTML = "Age: " + obj.age;
        document.querySelector(".bday").innerHTML = "Birthdate: " + obj.bday;
        document.querySelector(".address").innerHTML = "Address: " + obj.address;
        document.querySelector(".phone").innerHTML = "Phone: " + obj.phone;
        document.querySelector(".civilstatus").innerHTML = "Civil stand: " + obj.civilstatus;
        document.querySelector(".kids").innerHTML = "Kids: " + obj.kids;
        document.querySelector(".hobbies").innerHTML = "Hobbies: " + obj.hobbies;
        document.querySelector(".email").innerHTML = "E-Mail: " + obj.email;
        document.querySelector(".occupation").innerHTML = "Occupation: " + obj.occupation;
        document.querySelector(".prev-occupation").innerHTML = "Previous Occupation: " + obj.prevoccupation;
        document.querySelector(".military").innerHTML = "Military: " + obj.military;
        document.querySelector(".club").innerHTML = "Club: " + obj.club;
        document.querySelector(".legal").innerHTML = "Legal: " + obj.legal;
        document.querySelector(".political").innerHTML = "Political: " + obj.political;
        document.querySelector(".notes").innerHTML = "Notes: " + obj.notes;
      }

      const v_icon = document.createElement("ion-icon"); // View icon
      v_icon.className = "icon"
      v_icon.setAttribute("name", "eye-outline");

      const e_icon_div = document.createElement("div"); // Icon div
      e_icon_div.className = "chip-edit";

      e_icon_div.onclick = function() {
        document.querySelector('.main').style.display = "none";
        document.querySelector('.edit-container').style.display = "flex";

        document.querySelector(".e-name-tag").innerHTML = obj.name;

        document.querySelector(".e-age").innerHTML = obj.age;
        document.querySelector(".e-bday").innerHTML = obj.bday;
        document.querySelector(".e-address").innerHTML = obj.address;
        document.querySelector(".e-phone").innerHTML = obj.phone;
        document.querySelector(".e-civilstatus").innerHTML = obj.civilstatus;
        document.querySelector(".e-kids").innerHTML = obj.kids;
        document.querySelector(".e-hobbies").innerHTML = obj.hobbies;
        document.querySelector(".e-email").innerHTML = obj.email;
        document.querySelector(".e-occupation").innerHTML = obj.occupation;
        document.querySelector(".e-prev-occupation").innerHTML = obj.prevoccupation;
        document.querySelector(".e-military").innerHTML = obj.military;
        document.querySelector(".e-club").innerHTML = obj.club;
        document.querySelector(".e-legal").innerHTML = obj.legal;
        document.querySelector(".e-political").innerHTML = obj.political;
        document.querySelector(".e-notes").innerHTML = obj.notes;
      }

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

document.getElementById("backbtn").onclick = function() {
  document.querySelector('.main').style.display = "flex";
  document.querySelector('.container').style.display = "none";
}

document.getElementById("e-backbtn").onclick = function() {
  document.querySelector('.main').style.display = "flex";
  document.querySelector('.edit-container').style.display = "none";
}

document.getElementById("e-savebtn").onclick = function() {
  console.log("Save data to db");

  console.log(document.querySelector(".e-age").innerHTML);
  console.log(document.querySelector(".e-bday").innerHTML);
  console.log(document.querySelector(".e-address").innerHTML);
  console.log(document.querySelector(".e-phone").innerHTML);
  console.log(document.querySelector(".e-civilstatus").innerHTML);
  console.log(document.querySelector(".e-kids").innerHTML);
  console.log(document.querySelector(".e-hobbies").innerHTML);
  console.log(document.querySelector(".e-email").innerHTML);
  console.log(document.querySelector(".e-occupation").innerHTML);
  console.log(document.querySelector(".e-prev-occupation").innerHTML);
  console.log(document.querySelector(".e-military").innerHTML);
  console.log(document.querySelector(".e-club").innerHTML);
  console.log(document.querySelector(".e-legal").innerHTML);
  console.log(document.querySelector(".e-political").innerHTML);
  console.log(document.querySelector(".e-notes").innerHTML);


  document.getElementById("e-savebtn-p").innerHTML = "Saved!";
  delay(1000).then(() => document.getElementById("e-savebtn-p").innerHTML = "Save");
  


  // TODO Add saving to db
}

