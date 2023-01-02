let data;
let accData;
const element = document.getElementById("searchbar");

function delay(time) { // Because there is no default sleep function
  return new Promise(resolve => setTimeout(resolve, time));
}

async function main() {
  const res = await fetch("http://localhost:8080/people")

  data = await res.json();

  console.log(data);
  

  element.addEventListener("keyup", search_users);
  search_users();


  function search_users() {
    let input = document.getElementById('searchbar').value;
    input = input.toLowerCase();
    let x = document.querySelector('#list-holder');
    x.innerHTML = ""

    for (const [i, _] of Object.entries(data)) {
      let obj = data[i];

      if (obj.name.toLowerCase().includes(input)) {

        // Create Cards For Each Person

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

        // View

        v_icon_div.onclick = async function () {
          document.querySelector('.main').style.display = "none";
          document.querySelector('.container').style.display = "flex";

          document.querySelector(".name-tag").innerHTML = obj.name;

          document.querySelector(".maiden-name").innerHTML = obj.name;
          document.querySelector(".age").innerHTML = "Age: " + obj.age;
          document.querySelector(".bday").innerHTML = "Birthdate: " + obj.bday;
          document.querySelector(".address").innerHTML = "Address: " + obj.address;
          document.querySelector(".phone").innerHTML = "Phone: " + obj.phone;
          document.querySelector(".ssn").innerHTML = "SSN: " + obj.ssn;
          document.querySelector(".civilstatus").innerHTML = "Civil stand: " + obj.civilstatus;
          document.querySelector(".kids").innerHTML = "Kids: " + obj.kids;
          document.querySelector(".hobbies").innerHTML = "Hobbies: " + obj.hobbies;
          document.querySelector(".email").innerHTML = "E-Mail: " + obj.email;
          document.querySelector(".occupation").innerHTML = "Occupation: " + obj.occupation;
          document.querySelector(".prev-occupation").innerHTML = "Previous Occupation: " + obj.prevoccupation;
          document.querySelector(".education").innerHTML = "Education: " + obj.education;
          document.querySelector(".military").innerHTML = "Military stand: " + obj.military;
          document.querySelector(".religion").innerHTML = "Religion: " + obj.religion;
          document.querySelector(".pets").innerHTML = "Pets: " + obj.pets;
          document.querySelector(".club").innerHTML = "Club: " + obj.club;
          document.querySelector(".legal").innerHTML = "Legal: " + obj.legal;
          document.querySelector(".political").innerHTML = "Political: " + obj.political;
          document.querySelector(".notes").innerHTML = "Notes: " + obj.notes;

          // Accounts

          if (obj.accounts != null) {
            for (const [i, _] of Object.entries(obj.accounts)) {
              let accObj = obj.accounts[i];
  
              console.log(accObj);

              // Creating elements

              const base_div = document.createElement("div"); // Outer div
              base_div.className = "acc-chip";

              const pfp_img = document.createElement("img"); // Pfp img
              pfp_img.className = "userPfp";

              if (accObj.profilePicture != null) {
                pfp_img.src = "data:image/png;base64," + accObj.profilePicture[0];
                console.log("profilePicture exsits")
              } else {
                pfp_img.src = "https://as2.ftcdn.net/v2/jpg/03/32/59/65/1000_F_332596535_lAdLhf6KzbW6PWXBWeIFTovTii1drkbT.jpg"
              }

              const info_div = document.createElement("div"); // Info div
              info_div.className = "info-container";

              const service_p = document.createElement("p");
              service_p.className = "serviceName";
              service_p.innerHTML = accObj.service;

              const name_p = document.createElement("p");
              name_p.className = "userName";
              name_p.innerHTML = accObj.username;


              document.querySelector(".accounts").appendChild(base_div);
              base_div.appendChild(pfp_img);
              base_div.appendChild(info_div);
              info_div.appendChild(service_p);
              info_div.appendChild(name_p);

              if (accObj.bio != null) {
                const bio_p = document.createElement("p");
                bio_p.className = "userBio";
                bio_p.innerHTML = accObj.bio[0];

                info_div.appendChild(bio_p);
              }
            }
          }
        }

        const v_icon = document.createElement("ion-icon"); // View icon
        v_icon.className = "icon";
        v_icon.setAttribute("name", "eye-outline");

        const e_icon_div = document.createElement("div"); // Icon div
        e_icon_div.className = "chip-edit";

        e_icon_div.onclick = function () {
          document.querySelector('.main').style.display = "none";
          document.querySelector('.edit-container').style.display = "flex";

          document.querySelector("#e-showid").innerHTML = obj.id;

          document.querySelector(".e-name-tag").innerHTML = obj.name;

          document.querySelector(".e-maiden-name").innerHTML = obj.maidenname;
          document.querySelector(".e-age").innerHTML = obj.age;
          document.querySelector(".e-bday").innerHTML = obj.bday;
          document.querySelector(".e-address").innerHTML = obj.address;
          document.querySelector(".e-phone").innerHTML = obj.phone;
          document.querySelector(".e-ssn").innerHTML = obj.ssn;
          document.querySelector(".e-civilstatus").innerHTML = obj.civilstatus;
          document.querySelector(".e-kids").innerHTML = obj.kids;
          document.querySelector(".e-hobbies").innerHTML = obj.hobbies;
          document.querySelector(".e-email").innerHTML = obj.email;
          document.querySelector(".e-occupation").innerHTML = obj.occupation;
          document.querySelector(".e-prev-occupation").innerHTML = obj.prevoccupation;
          document.querySelector(".e-education").innerHTML = obj.education;
          document.querySelector(".e-military").innerHTML = obj.military;
          document.querySelector(".e-religion").innerHTML = obj.religion;
          document.querySelector(".e-pets").innerHTML = obj.pets;
          document.querySelector(".e-club").innerHTML = obj.club;
          document.querySelector(".e-legal").innerHTML = obj.legal;
          document.querySelector(".e-political").innerHTML = obj.political;
          document.querySelector(".e-notes").innerHTML = obj.notes;
        }

        const e_icon = document.createElement("ion-icon"); // Edit icon
        e_icon.className = "icon"
        e_icon.setAttribute("name", "create-outline");

        const d_icon_div = document.createElement("div"); // Icon div
        d_icon_div.className = "chip-edit";

        const d_icon = document.createElement("ion-icon"); // Edit icon
        d_icon.className = "icon"
        d_icon.setAttribute("name", "trash-outline");

        d_icon_div.onclick = function () {
          const headers = new Headers();
          headers.append('Access-Control-Allow-Origin', '*');

          fetch("http://localhost:8080/people/3", {
            method: "DELETE"
          });
        }


        base_div.appendChild(p_icon_div);
        base_div.appendChild(txt_div);
        txt_div.appendChild(name_p);

        base_div.appendChild(v_icon_div);
        base_div.appendChild(e_icon_div);
        base_div.appendChild(d_icon_div);
        p_icon_div.appendChild(p_icon);
        v_icon_div.appendChild(v_icon);
        e_icon_div.appendChild(e_icon);
        d_icon_div.appendChild(d_icon);


        name_p.innerHTML = `${obj.name}`
        x.appendChild(base_div)
      }



      document.getElementById("acc-name-tag").onkeypress = function(event) {
  // Check if the pressed key is the Enter key
  if (event.key === "Enter") {
        event.preventDefault();
    // Execute the search function
    search();
  }
  if (event.key == " ") {
    event.preventDefault();
  }
};
      document.getElementById("acc-searchbtn").onclick = search; 
let isButtonEnabled = true;
        async function search() {
            // Check if the button is enabled
  if (!isButtonEnabled) {
    return;
  }

  // Disable the button
  isButtonEnabled = false;

  // Code for the action you want to execute goes here
  console.log("Searching...");

  // Enable the button after a delay
          event.stopPropagation();
          console.log("search");

  // Set the flag to indicate that a request is in progress
        const response = await fetch('http://localhost:8080/getAccounts/' + document.getElementById("acc-name-tag").textContent);
        const data = await response.json();
      
        console.log(data);
    
        const term_container = document.createElement("div");
        term_container.className = "term-container";
    
        const term_header = document.createElement("p");
        term_header.className = "term-header";
        term_header.innerHTML = document.getElementById("acc-name-tag").innerHTML;
    
        term_container.appendChild(term_header);
    
    
        for (const [i, _] of Object.entries(data)) {
          let accObj = data[i];
    
          const row_div = document.createElement("div");
          row_div.className = "acc-row";
    
          const manage_acc_chip = document.createElement("div");
          manage_acc_chip.className = "manage-acc-chip"
    
          const outer_div = document.createElement("div");
          outer_div.className = "acc-chip";
    
          const user_pfp = document.createElement("img");
          user_pfp.className = "userPfp";
    
          if (obj.profilePicture != null) {
            user_pfp.src = "data:image/png;base64," + accObj.profilePicture[0];
          } else {
            user_pfp.src = "https://as2.ftcdn.net/v2/jpg/03/32/59/65/1000_F_332596535_lAdLhf6KzbW6PWXBWeIFTovTii1drkbT.jpg";
          }
    
          const info_container = document.createElement("div");
          info_container.className = "info-container";
    
          const service_name = document.createElement("p");
          service_name.className = "serviceName";
          service_name.innerHTML = accObj.service;
    
          const user_name = document.createElement("p");
          user_name.className = "userName";
          user_name.innerHTML = accObj.username;
    
          document.getElementById("accounts").appendChild(row_div);
          row_div.appendChild(term_container);
          term_container.appendChild(manage_acc_chip);
          manage_acc_chip.appendChild(outer_div);
          outer_div.appendChild(user_pfp);
          outer_div.appendChild(info_container);
          info_container.appendChild(service_name);
          info_container.appendChild(user_name);
    
          if (accObj.bio != null) {
            const user_bio = document.createElement("p");
            user_bio.className = "userBio";
            user_bio.innerHTML = accObj.bio[0];
    
            info_container.appendChild(user_bio);
          }
    
          const btn_container = document.createElement("div");
          btn_container.className = "manage-btn-container";
    
          const reject_btn = document.createElement("div");
          reject_btn.id = "acc-rejectbtn";
          reject_btn.className = "btn btn-secondary";
    
          const reject_p = document.createElement("p");
          reject_p.innerHTML = "Reject";
    
          const accept_btn = document.createElement("div");
          accept_btn.id = "acc-acceptbtn";
          accept_btn.className = "btn btn-secondary";
    
          const accept_p = document.createElement("p");
          accept_p.innerHTML = "Accept";
    
          manage_acc_chip.appendChild(btn_container);
          btn_container.appendChild(reject_btn);
          btn_container.appendChild(accept_btn);
          reject_btn.appendChild(reject_p);
          accept_btn.appendChild(accept_p);

          
    
          accept_btn.onclick = async function () {
            console.log(accObj.id);

            let midSave = [];

            midSave.push(accObj.id);

            document.getElementById("acc-midsave").innerHTML = midSave;
          }
        }

  setTimeout(function() {
    isButtonEnabled = true;
  }, 1000); // 1000ms = 1 second
      }
    }
  }




  document.getElementById("backbtn").onclick = function () { // back button in view ig
    document.querySelector('.main').style.display = "flex";
    document.querySelector('.container').style.display = "none";

    var elements = document.getElementsByClassName("acc-chip");

    while (elements.length > 0) {
      elements[0].parentNode.removeChild(elements[0]);
    }
  }

  document.getElementById("newbtn").onclick = function () {
    document.querySelector('.main').style.display = "none";
    document.querySelector('.create-container').style.display = "flex";
  }

  document.getElementById("e-backbtn").onclick = function () {
    document.querySelector('.main').style.display = "flex";
    document.querySelector('.edit-container').style.display = "none";
  }

  document.getElementById("e-backbtn").onclick = function () {
    document.querySelector('.main').style.display = "flex";
    document.querySelector('.edit-container').style.display = "none";
  }

  document.getElementById("c-accbtn").onclick = function () { // account button
    document.querySelector('.create-container').style.display = "none";
    document.querySelector('.acc-container').style.display = "flex";
  }

  document.getElementById("c-backbtn").onclick = function () {
    document.querySelector('.main').style.display = "flex";
    document.querySelector('.create-container').style.display = "none";
  }

  document.getElementById("acc-backbtn").onclick = function () { // account back button
    document.querySelector('.create-container').style.display = "flex";
    document.querySelector('.acc-container').style.display = "none";
  }



  document.getElementById("acc-savebtn").onclick = function () { // account menu save button
    


    document.getElementById("c-savebtn-p").innerHTML = "Saved!";
    delay(1000).then(() => document.getElementById("c-savebtn-p").innerHTML = "Save");
    document.querySelector('.create-container').style.display = "flex";
    document.querySelector('.acc-container').style.display = "none";
  }

  // CREATE

  document.getElementById("c-savebtn").onclick = function () { // new document save button
    console.log("Save data to db (new)");

    let totalIds = Object.keys(data).length;

    let id = String(totalIds + 1);

    let name = document.querySelector(".c-name-tag").innerHTML;

    let maidenname = document.querySelector(".c-maiden-name").innerHTML;
    let age = parseInt(document.querySelector(".c-age").innerHTML);
    let bday = document.querySelector(".c-bday").innerHTML;
    let address = document.querySelector(".c-address").innerHTML;
    let phone = document.querySelector(".c-phone").innerHTML;
    let ssn = document.querySelector(".c-ssn").innerHTML;
    let civilstatus = document.querySelector(".c-civilstatus").innerHTML;
    let kids = document.querySelector(".c-kids").innerHTML;
    let hobbies = document.querySelector(".c-hobbies").innerHTML;
    let email = document.querySelector(".c-email").innerHTML;
    let occupation = document.querySelector(".c-occupation").innerHTML;
    let prevoccupation = document.querySelector(".c-prev-occupation").innerHTML;
    let education = document.querySelector(".c-education").innerHTML;
    let military = document.querySelector(".c-military").innerHTML;
    let religion = document.querySelector(".c-religion").innerHTML;
    let pets = document.querySelector(".c-pets").innerHTML;
    let club = document.querySelector(".c-club").innerHTML;
    let legal = document.querySelector(".c-legal").innerHTML;
    let political = document.querySelector(".c-political").innerHTML;
    let notes = document.querySelector(".c-notes").innerHTML;

    fetch('http://localhost:8080/people', {
      method: 'POST',
      body: JSON.stringify({ "id": id, "name": name, "age": age, "bday": bday, "address": address, "phone": phone, "civilstatus": civilstatus, "kids": kids, "hobbies": hobbies, "email": email, "occupation": occupation, "prevoccupation": prevoccupation, "military": military, "club": club, "legal": legal, "political": political, "notes": notes })
    });

    document.getElementById("c-savebtn-p").innerHTML = "Saved!";
    delay(1000).then(() => document.getElementById("c-savebtn-p").innerHTML = "Save");
    document.querySelector('.main').style.display = "flex";
    document.querySelector('.create-container').style.display = "none";
  }

  // EDIT

  document.getElementById("e-savebtn").onclick = function () {
    console.log("Save data to db (edit)");

    let id = document.querySelector("#e-showid").innerHTML;

    let name = document.querySelector(".e-name-tag").innerHTML;

    let maidenname = document.querySelector(".c-maiden-name").innerHTML;
    let age = parseInt(document.querySelector(".c-age").innerHTML);
    let bday = document.querySelector(".c-bday").innerHTML;
    let address = document.querySelector(".c-address").innerHTML;
    let phone = document.querySelector(".c-phone").innerHTML;
    let ssn = document.querySelector(".c-ssn").innerHTML;
    let civilstatus = document.querySelector(".c-civilstatus").innerHTML;
    let kids = document.querySelector(".c-kids").innerHTML;
    let hobbies = document.querySelector(".c-hobbies").innerHTML;
    let email = document.querySelector(".c-email").innerHTML;
    let occupation = document.querySelector(".c-occupation").innerHTML;
    let prevoccupation = document.querySelector(".c-prev-occupation").innerHTML;
    let education = document.querySelector(".c-education").innerHTML;
    let military = document.querySelector(".c-military").innerHTML;
    let religion = document.querySelector(".c-religion").innerHTML;
    let pets = document.querySelector(".c-pets").innerHTML;
    let club = document.querySelector(".c-club").innerHTML;
    let legal = document.querySelector(".c-legal").innerHTML;
    let political = document.querySelector(".c-political").innerHTML;
    let notes = document.querySelector(".c-notes").innerHTML;


    fetch('http://localhost:8080/people', {
      method: 'POST',
      body: JSON.stringify({ "id": id, "maidenname": maidenname, "name": name, "age": age, "bday": bday, "address": address, "phone": phone, "ssn": ssn, "civilstatus": civilstatus, "kids": kids, "hobbies": hobbies, "email": email, "occupation": occupation, "prevoccupation": prevoccupation, "education": education, "military": military, "religion": religion, "pets": pets, "club": club, "legal": legal, "political": political, "notes": notes })
    });

    document.getElementById("e-savebtn-p").innerHTML = "Saved!";
    delay(1000).then(() => document.getElementById("e-savebtn-p").innerHTML = "Save");
    document.querySelector('.main').style.display = "flex";
    document.querySelector('.edit-container').style.display = "none";
  }
}

main()
