import {delay, SaveAsFile, checkGender, getGenderElementIndex, checkReligion, getReligionElementIndex, checkCivilstatus, getCivilstatusElementIndex} from "./framework.js";

const element = document.getElementById("searchbar");


// Api rul 
var hostname = window.location.hostname;
var port = window.location.port;

var baseUrl = hostname + ":" + port;

var apiUrl = "http://" + baseUrl;



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


async function main() {
  const res = await fetch(apiUrl + "/")

  let data = await res.json();


  element.addEventListener("keyup", search_users);
  search_users();

  document.getElementById("exportbtn").onclick = function () {
    SaveAsFile(JSON.stringify(data), "data.json", "text/plain;charset=utf-8");
  }


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

        const hitbox_abbr = document.createElement("abbr"); // hitbox abbr
        hitbox_abbr.title = "View"
        hitbox_abbr.className = "hitbox-abbr";

        const hitbox_div = document.createElement("div"); // hitbox div
        hitbox_div.className = "hitbox";

        const p_icon_div = document.createElement("div"); // Icon div
        p_icon_div.className = "chip-icon";

        const p_icon = document.createElement("ion-icon"); // Person icon
        p_icon.className = "icon"
        p_icon.setAttribute("name", "person");

        const txt_div = document.createElement("div"); // Text container
        txt_div.className = "text-container";

        const name_p = document.createElement("p"); // Name paragraph
        name_p.className = "card-text";

        // View

        hitbox_div.onclick = async function () {
          document.querySelector('.main').style.display = "none";
          document.querySelector('.container').style.display = "flex";

          document.querySelector("#v-showid").innerHTML = obj.id;

          document.querySelector(".name-tag").value = obj.name;

          document.querySelector(".gender").innerHTML = "Gender: " + obj.gender;
          document.querySelector(".age").innerHTML = "Age: " + obj.age;
          document.querySelector(".bday").innerHTML = "Birthdate: " + obj.bday;
          document.querySelector(".address").innerHTML = "Address: " + obj.address;
          document.querySelector(".civilstatus").innerHTML = "Civil stand: " + obj.civilstatus;
          document.querySelector(".kids").innerHTML = "Kids: " + obj.kids;
          document.querySelector(".occupation").innerHTML = "Occupation: " + obj.occupation;
          document.querySelector(".prevoccupation").innerHTML = "Previous Occupation: " + obj.prevoccupation;
          document.querySelector(".education").innerHTML = "Education: " + obj.education;
          document.querySelector(".religion").innerHTML = "Religion: " + obj.religion;
          document.querySelector(".pets").innerHTML = "Pets: " + obj.pets;
          document.querySelector(".legal").innerHTML = "Legal: " + obj.legal;
          document.querySelector(".political").innerHTML = "Political: " + obj.political;
          document.querySelector(".notes").innerHTML = obj.notes;


          let allObjectsAtStart = document.querySelectorAll(".viewtag");

          allObjectsAtStart.forEach(object => {
            object.style.display = "flex";
          });


          // Get all the elements with the class "viewtag" and store them in a variable called "allObjects"
          let allObjects = document.getElementsByClassName("viewtag");

          // Loop through all the objects in the array
          for (let i = 0; i < allObjects.length; i++) {

            // Store the current object's HTML in a variable called "item"
            let item = allObjects[i].innerHTML;
            // Get the text from the object's HTML and store it in a variable called "tempText"
            let tempText = item.substring(item.indexOf(':') + 1).trim();

            // Check if the text is empty, null, or undefined
            if (tempText.length <= 0 || tempText == "" || tempText == " " || tempText == null || tempText == undefined || tempText == 0) {
              // Remove the object from the page
              // allObjects[i].remove();

              allObjects[i].style.display = "none";
              // i--;
            }
          }

          if (document.getElementById("notes").innerHTML.length <= 0) {
            document.getElementById("space-maker").style.display = "none";
          }

          // Hobbies

          document.querySelector('.v-hobby-base').style.display = "block";

          if (Object.keys(obj.hobbies).length >= 1) {
            const hobbyContainer = document.querySelector('.v-hobby-base');

            for (const [_, hobby] of Object.entries(obj.hobbies)) {
              if (hobby.hobby != "" && hobby.hobby != null && hobby.hobby != undefined) {
                document.querySelector('.v-hobby-space-maker').style.display = "block";
                const container = document.createElement("div");
                container.className = "v-hobby-container";

                const subContainer = document.createElement("div");
                subContainer.className = "hobby-subcontainer";

                const hobby_input = document.createElement("input");
                hobby_input.className = "form-input v-hobby";
                hobby_input.id = "v-hobby";
                hobby_input.type = "hobby";
                hobby_input.placeholder = "Enter hobby";
                hobby_input.spellcheck = "false";
                hobby_input.required = "true";
                hobby_input.value = hobby.hobby;
                hobby_input.disabled = "true";

                hobbyContainer.appendChild(container);
                container.appendChild(subContainer);
                subContainer.appendChild(hobby_input);
              }
            };
          } else {
            document.querySelector('.v-hobby-space-maker').style.display = "none";
          }

          // IPs

          document.querySelector('.v-ip-base').style.display = "block";

          if (Object.keys(obj.ips).length >= 1) {
            const ipContainer = document.querySelector('.v-ip-base');

            for (const [_, ip] of Object.entries(obj.ips)) {
              if (ip.ip != "" && ip.ip != null && ip.ip != undefined) {
                document.querySelector('.v-ip-space-maker').style.display = "block";
                const container = document.createElement("div");
                container.className = "v-ip-container";

                const subContainer = document.createElement("div");
                subContainer.className = "ip-subcontainer";

                const ip_input = document.createElement("input");
                ip_input.className = "form-input v-ip";
                ip_input.id = "v-ip";
                ip_input.type = "ip";
                ip_input.placeholder = "Enter IP";
                ip_input.spellcheck = "false";
                ip_input.required = "true";
                ip_input.value = ip.ip;
                ip_input.disabled = "true";

                ipContainer.appendChild(container);
                container.appendChild(subContainer);
                subContainer.appendChild(ip_input);
              }
            };
          } else {
            document.querySelector('.v-ip-space-maker').style.display = "none";
          }

          // Clubs

          document.querySelector('.v-club-base').style.display = "block";

          if (Object.keys(obj.clubs).length >= 1) {
            const clubContainer = document.querySelector('.v-club-base');

            for (const [_, club] of Object.entries(obj.clubs)) {
              if (club.club != "" && club.club != null && club.club != undefined) {
                document.querySelector('.v-club-space-maker').style.display = "block";
                const container = document.createElement("div");
                container.className = "v-club-container";

                const subContainer = document.createElement("div");
                subContainer.className = "club-subcontainer";

                const club_input = document.createElement("input");
                club_input.className = "form-input v-club";
                club_input.id = "v-club";
                club_input.type = "club";
                club_input.placeholder = "Enter club";
                club_input.spellcheck = "false";
                club_input.required = "true";
                club_input.value = club.club;
                club_input.disabled = "true";

                clubContainer.appendChild(container);
                container.appendChild(subContainer);
                subContainer.appendChild(club_input);
              }
            };
          } else {
            document.querySelector('.v-club-space-maker').style.display = "none";
          }

          // Phone

          document.querySelector('.v-phone-base').style.display = "block";

          if (Object.keys(obj.phone).length >= 1) {
            const phoneContainer = document.querySelector('.v-phone-base');

            for (const [_, phone] of Object.entries(obj.phone)) {
              if (phone.number != "" && phone.number != null && phone.number != undefined) {
                document.querySelector('.v-phone-space-maker').style.display = "block";
                const container = document.createElement("div");
                container.className = "v-phone-container";

                const subContainer = document.createElement("div");
                subContainer.className = "phone-subcontainer";

                const phone_input = document.createElement("input");
                phone_input.className = "form-input v-phone";
                phone_input.id = "v-phone";
                phone_input.type = "phone";
                phone_input.placeholder = "Enter phone number";
                phone_input.spellcheck = "false";
                phone_input.maxLength = "30";
                phone_input.required = "true";
                phone_input.value = phone.number;
                phone_input.disabled = "true";


                const infoBtn = document.createElement("div");
                infoBtn.className = "v-info-btn";

                const icon = document.createElement("ion-icon");
                icon.setAttribute("name", "information-outline");

                container.appendChild(subContainer);
                infoBtn.appendChild(icon);
                subContainer.appendChild(phone_input);
                subContainer.appendChild(infoBtn);
                phoneContainer.appendChild(container);

                infoBtn.onclick = function () {
                  const infoDiv = container.querySelector(".v-info-div");

                  if (!infoDiv) {
                    const infoDiv = document.createElement("div");
                    infoDiv.className = "v-info-div";

                    container.appendChild(infoDiv);

                    if (phone.valid == true) {
                      const abbrContainerValidity = document.createElement("abbr")
                      abbrContainerValidity.className = "validity-abbr";
                      abbrContainerValidity.title = "Valid Phone Number";

                      const iconDivValid = document.createElement("div");
                      iconDivValid.className = "valid-icon-div";

                      const iconValid = document.createElement("img");
                      iconValid.className = "valid-icon phone-icon";
                      iconValid.src = "./images/valid-phone.png";

                      infoDiv.appendChild(abbrContainerValidity);
                      abbrContainerValidity.appendChild(iconDivValid);
                      iconDivValid.appendChild(iconValid);
                    } else if (phone.valid == false) {
                      const abbrContainerValidity = document.createElement("abbr")
                      abbrContainerValidity.className = "validity-abbr";
                      abbrContainerValidity.title = "Invalid Phone Number";

                      const iconDivValid = document.createElement("div");
                      iconDivValid.className = "valid-icon-div";

                      const iconValid = document.createElement("img");
                      iconValid.className = "valid-icon phone-icon";
                      iconValid.src = "./images/invalid-phone.png";

                      infoDiv.appendChild(abbrContainerValidity);
                      abbrContainerValidity.appendChild(iconDivValid);
                      iconDivValid.appendChild(iconValid);
                    }

                    // This should almost never fail
                    if (phone.phoneinfoga.Country != "" && phone.phoneinfoga.Country != null && phone.phoneinfoga.Country != undefined) {
                      const abbrContainer = document.createElement("abbr")
                      abbrContainer.className = "phone-info-abbr";
                      abbrContainer.title = phone.phoneinfoga.Country;

                      const iconDiv = document.createElement("div");
                      iconDiv.className = "service-icon-div";

                      const icon = document.createElement("img");
                      icon.className = "country-icon";
                      icon.src = "./images/flags/" + phone.phoneinfoga.Country + ".png";

                      infoDiv.appendChild(abbrContainer);
                      abbrContainer.appendChild(iconDiv);
                      iconDiv.appendChild(icon);
                    }
                  } else {
                    container.removeChild(infoDiv);
                  }
                }
              }
            };
          } else {
            document.querySelector('.v-phone-space-maker').style.display = "none";
          }


          // Email

          document.querySelector('.v-email-base').style.display = "block";

          if (Object.keys(obj.email).length >= 1 && obj.email != {}) {

            const emailContainer = document.querySelector('.v-email-base');

            for (const [_, email] of Object.entries(obj.email)) {
              if (email.mail != "" && email.mail != null && email.mail != undefined) {
                document.querySelector('.v-email-space-maker').style.display = "block";
                const container = document.createElement("div");
                container.className = "v-email-container";

                const subContainer = document.createElement("div");
                subContainer.className = "email-subcontainer";

                const email_input = document.createElement("input");
                email_input.className = "form-input v-mail";
                email_input.id = "v-e-mail";
                email_input.type = "email";
                email_input.placeholder = "Enter email address";
                email_input.spellcheck = "false";
                email_input.maxLength = "30";
                email_input.required = "true";
                email_input.value = email.mail;
                email_input.disabled = "true";


                const infoBtn = document.createElement("div");
                infoBtn.className = "v-info-btn";

                const icon = document.createElement("ion-icon");
                icon.setAttribute("name", "information-outline");

                container.appendChild(subContainer);
                infoBtn.appendChild(icon);
                subContainer.appendChild(email_input);
                subContainer.appendChild(infoBtn);
                emailContainer.appendChild(container);

                infoBtn.onclick = function () {
                  const infoDiv = container.querySelector(".v-info-div");

                  if (!infoDiv) {
                    const infoDiv = document.createElement("div");
                    infoDiv.className = "v-info-div";

                    container.appendChild(infoDiv);

                    if (email.valid == true) {
                      const abbrContainerValidity = document.createElement("abbr")
                      abbrContainerValidity.className = "validity-abbr";
                      abbrContainerValidity.title = "Valid Email";

                      const iconDivValid = document.createElement("div");
                      iconDivValid.className = "valid-icon-div";

                      const iconValid = document.createElement("img");
                      iconValid.className = "valid-icon";
                      iconValid.src = "./images/valid.png";

                      infoDiv.appendChild(abbrContainerValidity);
                      abbrContainerValidity.appendChild(iconDivValid);
                      iconDivValid.appendChild(iconValid);
                    } else if (email.valid == false) {
                      const abbrContainerValidity = document.createElement("abbr")
                      abbrContainerValidity.className = "validity-abbr";
                      abbrContainerValidity.title = "Invalid Email";

                      const iconDivValid = document.createElement("div");
                      iconDivValid.className = "valid-icon-div";

                      const iconValid = document.createElement("img");
                      iconValid.className = "valid-icon";
                      iconValid.src = "./images/invalid.png";

                      infoDiv.appendChild(abbrContainerValidity);
                      abbrContainerValidity.appendChild(iconDivValid);
                      iconDivValid.appendChild(iconValid);
                    }

                    if (email.services != undefined && email.services != null) {
                      for (const [_, service] of Object.entries(email.services)) {
                        const abbrContainer = document.createElement("abbr")
                        abbrContainer.className = "service-abbr";
                        abbrContainer.title = service.name;

                        const iconDiv = document.createElement("div");
                        iconDiv.className = "service-icon-div";

                        const icon = document.createElement("img");
                        icon.className = "service-icon";
                        icon.src = service.icon;

                        
                        infoDiv.appendChild(abbrContainer);
                        abbrContainer.appendChild(iconDiv);
                        iconDiv.appendChild(icon);

                        iconDiv.onclick = function () {
                          if (service.link != "") {
                            window.open(service.link, "_blank");
                          }
                        }
                      };

                      container.appendChild(infoDiv);
                    }
                  } else {
                    container.removeChild(infoDiv);
                  }
                }
              }
            };
          } else {
            document.querySelector('.v-email-space-maker').style.display = "none";
          }


          // Accounts

          if (obj.accounts != {} && obj.accounts != null) {
            for (const [_, accObj] of Object.entries(obj.accounts)) {
              //let accObj = obj.accounts[i];

              // Creating elements

              const base_div = document.createElement("div"); // Outer div
              base_div.className = "acc-chip";

              const pfp_img = document.createElement("img"); // Pfp img
              pfp_img.className = "userPfp";

              if (accObj.profilePicture != null) {
                pfp_img.src = "data:image/png;base64," + accObj.profilePicture["1"].img;
              } else {
                pfp_img.src = "https://as2.ftcdn.net/v2/jpg/03/32/59/65/1000_F_332596535_lAdLhf6KzbW6PWXBWeIFTovTii1drkbT.jpg"
              }

              const info_div = document.createElement("div"); // Info div
              info_div.className = "info-container";

              const service_p = document.createElement("a");
              service_p.className = "serviceName";
              service_p.innerHTML = accObj.service;
              service_p.href = accObj.url;
              service_p.target = "_blank";

              const name_p = document.createElement("a");
              name_p.className = "userName";
              name_p.innerHTML = accObj.username;
              name_p.href = accObj.url;
              name_p.target = "_blank";


              document.querySelector(".accounts").appendChild(base_div);
              base_div.appendChild(pfp_img);
              base_div.appendChild(info_div);
              info_div.appendChild(service_p);
              info_div.appendChild(name_p);

              if (accObj.bio != null) {
                const bio_p = document.createElement("p");
                bio_p.className = "userBio";
                bio_p.innerHTML = accObj.bio["1"].bio;

                info_div.appendChild(bio_p);
              }
            }
          }
        }

        document.getElementById("savetxtbtn").onclick = async function () {
          let textToSave = "";

          let getId = document.getElementById("v-showid").innerHTML;

          const res = await fetch(apiUrl + "/people/" + getId)

          data = await res.json();

          // For each item in data: check if the field is empty, if not, add item to textToSave
          // For each item in data: check if the field is empty, if not, add item to textToSave
          if (data.name) {
            textToSave += `Name: ${data.name}\n`;
          }
          for (const [key, value] of Object.entries(data)) {
            if (key === "accounts" && typeof value === "object" && value != null) {
              let accountsText = "";
              let first = true;
              for (const [service, account] of Object.entries(value)) {
                if (first) {
                  accountsText += ` ${account.service}, ${account.username}, ${account.url}`;
                  first = false;
                } else {
                  accountsText += `\n          ${account.service}, ${account.username}, ${account.url}`;
                }
              }
              textToSave += `${key.charAt(0).toUpperCase()}${key.slice(1)}:${accountsText}\n`;
            } else if (key !== "id" && value && value != " " && value != null && value != undefined && value != 0 && key != "name") {
              textToSave += `${key.charAt(0).toUpperCase()}${key.slice(1)}: ${value}\n`;
            }
          }

          var textToSave1 = textToSave.replace(/<br>/g, "\n       ");


          SaveAsFile(textToSave1, data.name.toLowerCase().replace(/ /g, "") + ".txt", "text/plain;charset=utf-8");
        }


        const e_icon_div = document.createElement("div"); // Icon div
        e_icon_div.className = "chip-edit";

        const e_abbr = document.createElement("abbr");
        e_abbr.title = "Edit"

        e_icon_div.onclick = function () {
          document.querySelector('.main').style.display = "none";
          document.querySelector('.edit-container').style.display = "flex";

          document.querySelector("#e-showid").innerHTML = obj.id;

          document.querySelector(".e-name-tag").value = obj.name;

          if (obj.gender != "") {
            const genderSelect = document.querySelector(".edit-container > .components > .scroll-box > div:nth-child(1) > .gender-select");
            const selectItems = genderSelect.querySelector(".select-items");
            const selectSelected = genderSelect.querySelector(".select-selected");
  
            const genderIndex = getGenderElementIndex(obj.gender);
  
            const genderElement = selectItems.children[genderIndex];
  
            selectSelected.innerHTML = obj.gender;
            genderElement.className = "same-as-selected";
          }

          document.querySelector(".e-age").innerHTML = obj.age;
          document.querySelector(".e-bday").innerHTML = obj.bday;
          document.querySelector(".e-address").innerHTML = obj.address;

          // Phone

          if (Object.keys(obj.phone).length >= 1) {
            const phoneContainer = document.querySelector('.phone-base');

            for (const [_, phone] of Object.entries(obj.phone)) {
              const container = document.createElement("div");
              container.className = "phone-container";

              const subContainer = document.createElement("div");
              subContainer.className = "phone-subcontainer";

              const phone_input = document.createElement("input");
              phone_input.className = "form-input phone";
              phone_input.id = "e-mail";
              phone_input.type = "email";
              phone_input.placeholder = "Enter phone number";
              phone_input.spellcheck = "false";
              phone_input.maxLength = "30";
              phone_input.required = "true";
              phone_input.value = phone.number;

              const del_btn_div = document.createElement("div");
              del_btn_div.className = "del-btn";

              const del_btn = document.createElement("ion-icon");
              del_btn.name = "remove-outline";

              container.appendChild(subContainer);
              subContainer.appendChild(phone_input);
              phoneContainer.appendChild(container);
              subContainer.appendChild(del_btn_div);
              del_btn_div.appendChild(del_btn);


              del_btn.onclick = function () {
                container.remove();
              }
            };
          }

          document.getElementById("phone-add-btn").onclick = function () {
            const phone_base = document.querySelector(".phone-base");

            const phone_container = document.createElement("div");
            phone_container.className = "phone-container";

            const subContainer = document.createElement("div");
            subContainer.className = "phone-subcontainer";

            const phone_input = document.createElement("input");
            phone_input.className = "form-input e-phone";
            phone_input.id = "phone";
            phone_input.type = "tel";
            phone_input.placeholder = "Enter phone number";
            phone_input.spellcheck = "false";
            //phone_input.maxLength = "15"; // FIXME some formattings can have more then 15 chars.
            phone_input.required = "true";

            const del_btn_div = document.createElement("div");
            del_btn_div.className = "del-btn";

            const del_btn = document.createElement("ion-icon");
            del_btn.name = "remove-outline";

            phone_base.appendChild(phone_container);
            phone_container.appendChild(subContainer);
            subContainer.appendChild(phone_input);
            subContainer.appendChild(del_btn_div);
            del_btn_div.appendChild(del_btn);

            del_btn_div.onclick = function () {
              phone_container.remove();
            }
          }

          if (obj.civilstatus != "") {
            const civilstatusSelect = document.querySelector(".edit-container > .components > .scroll-box > div:nth-child(6) > .civilstatus-select");
            const selectItems = civilstatusSelect.querySelector(".select-items");
            const selectSelected = civilstatusSelect.querySelector(".select-selected");
  
            const civilstatusIndex = getCivilstatusElementIndex(obj.civilstatus);
  
            const civilstatusElement = selectItems.children[civilstatusIndex];
  
            selectSelected.innerHTML = obj.civilstatus;
            civilstatusElement.className = "same-as-selected";
          }

          document.querySelector(".e-kids").innerHTML = obj.kids;

          // Hobbies

          if (Object.keys(obj.hobbies).length >= 1) {
            const hobbyContainer = document.querySelector('.e-hobby-base');

            for (const [_, hobby] of Object.entries(obj.hobbies)) {
              const container = document.createElement("div");
              container.className = "hobby-container";

              const subContainer = document.createElement("div");
              subContainer.className = "hobby-subcontainer";

              const hobby_input = document.createElement("input");
              hobby_input.className = "form-input hobby";
              hobby_input.id = "e-mail";
              hobby_input.type = "email";
              hobby_input.placeholder = "Enter hobby";
              hobby_input.spellcheck = "false";
              hobby_input.value = hobby.hobby;

              const del_btn_div = document.createElement("div");
              del_btn_div.className = "del-btn";

              const del_btn = document.createElement("ion-icon");
              del_btn.name = "remove-outline";

              container.appendChild(subContainer);
              subContainer.appendChild(hobby_input);
              hobbyContainer.appendChild(container);
              subContainer.appendChild(del_btn_div);
              del_btn_div.appendChild(del_btn);


              del_btn.onclick = function () {
                container.remove();
              }
            };
          }

          document.getElementById("hobby-add-btn").onclick = function () {
            const hobby_base = document.querySelector(".e-hobby-base");

            const hobby_container = document.createElement("div");
            hobby_container.className = "hobby-container";

            const subContainer = document.createElement("div");
            subContainer.className = "hobby-subcontainer";

            const hobby_input = document.createElement("input");
            hobby_input.className = "form-input e-hobby";
            hobby_input.id = "hobby";
            hobby_input.type = "text";
            hobby_input.placeholder = "Enter hobby";
            hobby_input.spellcheck = "false";
            hobby_input.required = "true";

            const del_btn_div = document.createElement("div");
            del_btn_div.className = "del-btn";

            const del_btn = document.createElement("ion-icon");
            del_btn.name = "remove-outline";

            hobby_base.appendChild(hobby_container);
            hobby_container.appendChild(subContainer);
            subContainer.appendChild(hobby_input);
            subContainer.appendChild(del_btn_div);
            del_btn_div.appendChild(del_btn);

            del_btn_div.onclick = function () {
              hobby_container.remove();
            }
          }
          
          document.querySelector(".e-occupation").innerHTML = obj.occupation;
          document.querySelector(".e-prevoccupation").innerHTML = obj.prevoccupation;
          document.querySelector(".e-education").innerHTML = obj.education;
          
          if (obj.religion != "") {
            const religionSelect = document.querySelector(".edit-container > .components > .scroll-box > div:nth-child(13) > .religion-select");
            const selectItems = religionSelect.querySelector(".select-items");
            const selectSelected = religionSelect.querySelector(".select-selected");
  
            const religionIndex = getReligionElementIndex(obj.religion);
  
            const religionElement = selectItems.children[religionIndex];
  
            selectSelected.innerHTML = obj.religion;
            religionElement.className = "same-as-selected";
          }

          document.querySelector(".e-pets").innerHTML = obj.pets;

          // Clubs

          if (Object.keys(obj.clubs).length >= 1) {
            const clubContainer = document.querySelector('.e-club-base');

            for (const [_, club] of Object.entries(obj.clubs)) {
              const container = document.createElement("div");
              container.className = "club-container";

              const subContainer = document.createElement("div");
              subContainer.className = "club-subcontainer";

              const club_input = document.createElement("input");
              club_input.className = "form-input club";
              club_input.id = "e-club";
              club_input.type = "text";
              club_input.placeholder = "Enter club";
              club_input.spellcheck = "false";
              club_input.value = club.club;

              const del_btn_div = document.createElement("div");
              del_btn_div.className = "del-btn";

              const del_btn = document.createElement("ion-icon");
              del_btn.name = "remove-outline";

              container.appendChild(subContainer);
              subContainer.appendChild(club_input);
              clubContainer.appendChild(container);
              subContainer.appendChild(del_btn_div);
              del_btn_div.appendChild(del_btn);


              del_btn.onclick = function () {
                container.remove();
              }
            };
          }

          document.getElementById("hobby-add-btn").onclick = function () {
            const hobby_base = document.querySelector(".e-hobby-base");

            const hobby_container = document.createElement("div");
            hobby_container.className = "hobby-container";

            const subContainer = document.createElement("div");
            subContainer.className = "hobby-subcontainer";

            const hobby_input = document.createElement("input");
            hobby_input.className = "form-input e-hobby";
            hobby_input.id = "hobby";
            hobby_input.type = "text";
            hobby_input.placeholder = "Enter hobby";
            hobby_input.spellcheck = "false";
            hobby_input.required = "true";

            const del_btn_div = document.createElement("div");
            del_btn_div.className = "del-btn";

            const del_btn = document.createElement("ion-icon");
            del_btn.name = "remove-outline";

            hobby_base.appendChild(hobby_container);
            hobby_container.appendChild(subContainer);
            subContainer.appendChild(hobby_input);
            subContainer.appendChild(del_btn_div);
            del_btn_div.appendChild(del_btn);

            del_btn_div.onclick = function () {
              hobby_container.remove();
            }
          }

          document.querySelector(".e-legal").innerHTML = obj.legal;
          document.querySelector(".e-political").innerHTML = obj.political;
          document.querySelector(".e-notes").innerHTML = obj.notes;

          // IPs

          if (Object.keys(obj.ips).length >= 1) {
            const ipContainer = document.querySelector('.e-ip-base');

            for (const [_, ip] of Object.entries(obj.ips)) {
              console.log(ip.ip);
              const container = document.createElement("div");
              container.className = "ip-container";

              const subContainer = document.createElement("div");
              subContainer.className = "ip-subcontainer";

              const ip_input = document.createElement("input");
              ip_input.className = "form-input ip";
              ip_input.id = "e-ip";
              ip_input.type = "text";
              ip_input.placeholder = "Enter IP";
              ip_input.spellcheck = "false";
              ip_input.value = ip.ip;

              const del_btn_div = document.createElement("div");
              del_btn_div.className = "del-btn";

              const del_btn = document.createElement("ion-icon");
              del_btn.name = "remove-outline";

              container.appendChild(subContainer);
              subContainer.appendChild(ip_input);
              ipContainer.appendChild(container);
              subContainer.appendChild(del_btn_div);
              del_btn_div.appendChild(del_btn);


              del_btn.onclick = function () {
                container.remove();
              }
            };
          }

          document.getElementById("ip-add-btn").onclick = function () {
            const ip_base = document.querySelector(".e-ip-base");

            const ip_container = document.createElement("div");
            ip_container.className = "ip-container";

            const subContainer = document.createElement("div");
            subContainer.className = "ip-subcontainer";

            const ip_input = document.createElement("input");
            ip_input.className = "form-input e-ip";
            ip_input.id = "ip";
            ip_input.type = "text";
            ip_input.placeholder = "Enter IP";
            ip_input.spellcheck = "false";
            ip_input.required = "true";

            const del_btn_div = document.createElement("div");
            del_btn_div.className = "del-btn";

            const del_btn = document.createElement("ion-icon");
            del_btn.name = "remove-outline";

            ip_base.appendChild(ip_container);
            ip_container.appendChild(subContainer);
            subContainer.appendChild(ip_input);
            subContainer.appendChild(del_btn_div);
            del_btn_div.appendChild(del_btn);

            del_btn_div.onclick = function () {
              ip_container.remove();
            }
          }

          // Email

          if (Object.keys(obj.email).length >= 1) {
            const emailContainer = document.querySelector('.email-base');

            for (const [_, email] of Object.entries(obj.email)) {
              const container = document.createElement("div");
              container.className = "email-container";

              const subContainer = document.createElement("div");
              subContainer.className = "email-subcontainer";

              const email_input = document.createElement("input");
              email_input.className = "form-input e-mail";
              email_input.id = "e-mail";
              email_input.type = "email";
              email_input.placeholder = "Enter email address";
              email_input.spellcheck = "false";
              email_input.maxLength = "30";
              email_input.required = "true";
              email_input.value = email.mail;

              const del_btn_div = document.createElement("div");
              del_btn_div.className = "del-btn";

              const del_btn = document.createElement("ion-icon");
              del_btn.name = "remove-outline";

              container.appendChild(subContainer);
              subContainer.appendChild(email_input);
              emailContainer.appendChild(container);
              subContainer.appendChild(del_btn_div);
              del_btn_div.appendChild(del_btn);

              if (email.services != undefined && email.services != null && email.services != "") {
                const hidden_email_save = document.createElement("p");
                hidden_email_save.className = "hidden-email-save";

                hidden_email_save.innerHTML = JSON.stringify(email.services);
                container.appendChild(hidden_email_save);
              }


              del_btn.onclick = function () {
                container.remove();
              }
            };
          }



          document.getElementById("email-add-btn").onclick = function () {
            const email_base = document.querySelector(".email-base");

            const email_container = document.createElement("div");
            email_container.className = "email-container";

            const subContainer = document.createElement("div");
            subContainer.className = "email-subcontainer";

            const email_input = document.createElement("input");
            email_input.className = "form-input e-mail";
            email_input.id = "e-mail";
            email_input.type = "email";
            email_input.placeholder = "Enter email address";
            email_input.spellcheck = "false";
            email_input.maxLength = "30";
            email_input.required = "true";

            const del_btn_div = document.createElement("div");
            del_btn_div.className = "del-btn";

            const del_btn = document.createElement("ion-icon");
            del_btn.name = "remove-outline";

            email_base.appendChild(email_container);
            email_container.appendChild(subContainer);
            subContainer.appendChild(email_input);
            subContainer.appendChild(del_btn_div);
            del_btn_div.appendChild(del_btn);

            const hidden_email_save = document.createElement("p");
            hidden_email_save.className = "hidden-email-save";
            email_container.appendChild(hidden_email_save);

            del_btn_div.onclick = function () {
              email_container.remove();
            }
          }

          // Accounts

          if (obj.accounts != {} && obj.accounts != null) {
            for (const [_, accObj] of Object.entries(obj.accounts)) {
              //let accObj = obj.accounts[i];

              // Creating elements

              const base_div = document.createElement("div"); // Outer div
              base_div.className = "acc-chip";

              const pfp_img = document.createElement("img"); // Pfp img
              pfp_img.className = "userPfp";

              if (accObj.profilePicture != null) {
                pfp_img.src = "data:image/png;base64," + accObj.profilePicture["1"].img;
              } else {
                pfp_img.src = "https://as2.ftcdn.net/v2/jpg/03/32/59/65/1000_F_332596535_lAdLhf6KzbW6PWXBWeIFTovTii1drkbT.jpg"
              }

              const info_div = document.createElement("div"); // Info div
              info_div.className = "info-container";

              const icon_space = document.createElement("div");
              icon_space.className = "icon-space";

              const service_p = document.createElement("a");
              service_p.className = "serviceName";
              service_p.innerHTML = accObj.service;
              service_p.href = accObj.url;
              service_p.target = "_blank";

              const name_p = document.createElement("a");
              name_p.className = "userName";
              name_p.innerHTML = accObj.username;
              name_p.href = accObj.url;
              name_p.target = "_blank";

              document.querySelector(".e-accounts").appendChild(base_div);
              base_div.appendChild(pfp_img);
              base_div.appendChild(info_div);
              base_div.appendChild(icon_space);
              info_div.appendChild(service_p);
              info_div.appendChild(name_p);

              if (accObj.service.toLowerCase() == "github") { // If the service is github, add a deep investigation button
                const deep_btn = document.createElement("div");
                deep_btn.className = "deepInvBtn btn btn-secondary";
                deep_btn.id = "deepInvBtn";

                const deep_btn_txt = document.createElement("p");
                deep_btn_txt.className = "deepInvBtnTxt";
                deep_btn_txt.innerHTML = "Deep Investigation";

                base_div.appendChild(deep_btn);
                deep_btn.appendChild(deep_btn_txt);

                const del_btn_div = document.createElement("div");
                del_btn_div.className = "delAccBtn-deep btn btn-secondary";

                const del_btn = document.createElement("ion-icon");
                del_btn.name = "remove-outline";

                base_div.appendChild(del_btn_div);
                del_btn_div.appendChild(del_btn);


                // Deep investigation
                deep_btn.onclick = async function () {
                  if (icon_space.firstChild) {
                    icon_space.firstChild.remove();
                  }

                  deep_btn_txt.innerHTML = "This may take up to an hour...";

                  const loadingSpinner = document.createElement("div");
                  loadingSpinner.className = "neu";
                  loadingSpinner.id = "deepInvLoadingSpinner";
                  loadingSpinner.style.display = "flex";

                  const loadingSpinnerShape = document.createElement("div");
                  loadingSpinnerShape.className = "neu_shape";

                  const loadingSpinnerInner = document.createElement("div");
                  loadingSpinnerInner.className = "neu_inner";

                  const loadingSpinnerBall = document.createElement("div");
                  loadingSpinnerBall.className = "neu_ball";


                  icon_space.appendChild(loadingSpinner)
                  loadingSpinner.appendChild(loadingSpinnerShape);
                  loadingSpinnerShape.appendChild(loadingSpinnerInner);
                  loadingSpinnerInner.appendChild(loadingSpinnerBall);
                  loadingSpinnerInner.appendChild(loadingSpinnerBall.cloneNode());
                  loadingSpinnerInner.appendChild(loadingSpinnerBall.cloneNode());
                  loadingSpinnerInner.appendChild(loadingSpinnerBall.cloneNode());


                  const res = await fetch(apiUrl + "/deep/github/" + accObj.username)
                  let data = await res.json();

                  loadingSpinner.remove();

                  const deepInvResIcon = document.createElement("img");
                  deepInvResIcon.className = "deepInvResIcon";

                  icon_space.appendChild(deepInvResIcon);

                  if (data != null && data != {} && res.status == 200) {
                    deep_btn_txt.innerHTML = "Deep Investigation";

                    deepInvResIcon.src = "./images/checkmark.png";
                    deepInvResIcon.style.filter = "drop-shadow(0.3rem 0.3rem 0.2rem var(--greyLight-2)) drop-shadow(-0.2rem -0.2rem 0.5rem var(--white));"

                    for (const [i, _] of Object.entries(data)) {
                      let obj = data[i];

                      const email_base = document.querySelector(".email-base");

                      const email_container = document.createElement("div");
                      email_container.className = "email-container";

                      const subContainer = document.createElement("div");
                      subContainer.className = "email-subcontainer";

                      const email_input = document.createElement("input");
                      email_input.className = "form-input e-mail";
                      email_input.id = "e-mail";
                      email_input.type = "email";
                      email_input.placeholder = "Enter email address";
                      email_input.spellcheck = "false";
                      email_input.maxLength = "30";
                      email_input.required = "true";

                      email_input.value = obj.mail;

                      const del_btn_div = document.createElement("div");
                      del_btn_div.className = "del-btn";

                      const del_btn = document.createElement("ion-icon");
                      del_btn.name = "remove-outline";

                      const hidden_email_save = document.createElement("p");
                      hidden_email_save.className = "hidden-email-save";

                      hidden_email_save.innerHTML = JSON.stringify(obj.services);

                      email_base.appendChild(email_container);
                      email_container.appendChild(subContainer);
                      subContainer.appendChild(email_input);
                      subContainer.appendChild(del_btn_div);
                      del_btn_div.appendChild(del_btn);
                      email_container.appendChild(hidden_email_save);

                      del_btn_div.onclick = function () {
                        email_container.remove();
                      }
                    }
                  } else if (res.status == 403 && data["fatal"] == "rate limited") {
                    deepInvResIcon.src = "./images/limited.png";
                    deepInvResIcon.style.filter = "drop-shadow(0.3rem 0.3rem 0.15rem var(--greyLight-2)) drop-shadow(-0.2rem -0.2rem 0.5rem var(--white));"
                  } else {
                    deepInvResIcon.src = "./images/cross.png";
                    deepInvResIcon.style.filter = "drop-shadow(0.3rem 0.3rem 0.2rem var(--greyLight-2)) drop-shadow(-0.2rem -0.2rem 0.5rem var(--white));"
                  }
                }


                del_btn_div.onclick = function () {
                  fetch(apiUrl + "/people/" + document.querySelector("#e-showid").innerHTML + "/accounts/" + accObj.service + "-" + accObj.username + "/delete", {
                    method: "GET",
                    mode: "no-cors"
                  });

                  base_div.remove();
                  // TODO Add stuff here
                }
              } else {
                const del_btn_div = document.createElement("div");
                del_btn_div.className = "delAccBtn btn btn-secondary";

                const del_btn = document.createElement("ion-icon");
                del_btn.name = "remove-outline";

                base_div.appendChild(del_btn_div);
                del_btn_div.appendChild(del_btn);

                del_btn_div.onclick = function () {
                  fetch(apiUrl + "/people/" + document.querySelector("#e-showid").innerHTML + "/accounts/" + accObj.service + "-" + accObj.username + "/delete", {
                    method: "GET",
                    mode: "no-cors"
                  });

                  base_div.remove();

                  // TODO Add stuff here
                }
              }

              if (accObj.bio != null) {
                const bio_p = document.createElement("p");
                bio_p.className = "userBio";
                bio_p.innerHTML = accObj.bio["1"].bio;

                info_div.appendChild(bio_p);
              }
            }
          }
        }

        const e_icon = document.createElement("ion-icon"); // Edit icon
        e_icon.className = "icon"
        e_icon.setAttribute("name", "create-outline");

        const acc_icon_div = document.createElement("div"); // Icon div
        acc_icon_div.className = "chip-acc";

        const acc_abbr = document.createElement("abbr");
        acc_abbr.title = "Add Accounts"

        const acc_icon = document.createElement("ion-icon"); // Edit icon
        acc_icon.className = "icon"
        acc_icon.setAttribute("name", "person-circle-outline");

        const d_icon_div = document.createElement("div"); // Icon div
        d_icon_div.className = "chip-delete";

        const d_abbr = document.createElement("abbr");
        d_abbr.title = "Delete"

        const d_icon = document.createElement("ion-icon"); // Edit icon
        d_icon.className = "icon"
        d_icon.setAttribute("name", "trash-outline");

        d_icon_div.onclick = function () {
          fetch(apiUrl + "/people/" + obj.id + "/delete", {
            method: "GET",
            mode: "no-cors"
          }).then(function () {
            location.reload();
          });
        }

        acc_icon_div.onclick = function () {
          document.querySelector("#e-showid").innerHTML = obj.id;
          document.querySelector('.main').style.display = "none";
          document.querySelector('.acc-container').style.display = "flex";
        }

        document.getElementById("acc-search-chip").onclick = search;

        document.getElementById("acc-name-tag").onkeypress = function (event) {
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

        let isButtonEnabled = true;

        async function search() {
          if (document.getElementById("acc-name-tag").value == "") {
            return;
          }
          // Check if the button is enabled
          if (!isButtonEnabled) {
            return;
          }

          // Disable the button
          isButtonEnabled = false;

          document.getElementById("loading-spinner").style.display = "inline-block";

          // Set the flag to indicate that a request is in progress
          const response = await fetch(apiUrl + '/getAccounts/' + document.getElementById("acc-name-tag").value);
          const data = await response.json();

          const term_container = document.createElement("div");
          term_container.className = "term-container";

          const term_header = document.createElement("p");
          term_header.className = "term-header";
          term_header.innerHTML = document.getElementById("acc-name-tag").value;

          term_container.appendChild(term_header);

          if (data != null && Object.entries(data).length >= 1) {
            document.getElementById("acc-no-results").style.display = "none";
            document.getElementById("acc-scroll-box").style.display = "block";

            const row_div = document.createElement("div");
            row_div.className = "acc-row";

            document.getElementById("accounts").appendChild(row_div);


            for (const [i, _] of Object.entries(data)) {
              let accObj = data[i];

              const manage_acc_chip = document.createElement("div");
              manage_acc_chip.className = "manage-acc-chip"

              const outer_div = document.createElement("div");
              outer_div.className = "acc-chip";

              const user_pfp = document.createElement("img");
              user_pfp.className = "userPfp";

              if (accObj.profilePicture != null) {
                user_pfp.src = "data:image/png;base64," + accObj.profilePicture["1"].img;
              } else {
                user_pfp.src = "https://as2.ftcdn.net/v2/jpg/03/32/59/65/1000_F_332596535_lAdLhf6KzbW6PWXBWeIFTovTii1drkbT.jpg";
              }

              const info_container = document.createElement("div");
              info_container.className = "info-container";

              const service_name = document.createElement("p");
              service_name.className = "serviceName";
              service_name.innerHTML = accObj.service;

              const user_name = document.createElement("a");
              user_name.className = "userName";
              user_name.innerHTML = accObj.username;
              user_name.href = accObj.url;
              user_name.target = "_blank";

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
                user_bio.innerHTML = accObj.bio["1"].bio;

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
                // Check if accObj.service and accObj.username are also in accounts object at obj.accounts
                let getId = document.getElementById("e-showid").innerHTML

                const res = await fetch(apiUrl + "/people/" + getId)

                let data = await res.json();

                data.accounts[accObj.service + "-" + accObj.username] = accObj;

                fetch(apiUrl + "/person", {
                  method: 'POST',
                  body: JSON.stringify(data)
                });

                accept_p.innerHTML = "Accepted!";
              }

              reject_btn.onclick = async function () {
                let elementCount = term_container.childElementCount;

                if (elementCount > 2) {
                  manage_acc_chip.remove();
                } else {
                  row_div.remove();
                }
              }
            }
          } else {
            // No accounts found

            if (document.getElementById("accounts").childElementCount <= 0) {
              document.getElementById("acc-no-results").style.display = "flex";
              document.getElementById("acc-scroll-box").style.display = "none";
            }
          }

          document.getElementById("loading-spinner").style.display = "none";
          isButtonEnabled = true;
        }

        base_div.appendChild(hitbox_abbr);
        hitbox_abbr.appendChild(hitbox_div);
        hitbox_div.appendChild(p_icon_div);
        hitbox_div.appendChild(txt_div);
        txt_div.appendChild(name_p);
        base_div.appendChild(e_abbr);
        e_abbr.appendChild(e_icon_div);
        base_div.appendChild(acc_abbr);
        base_div.appendChild(d_abbr);
        acc_abbr.appendChild(acc_icon_div);
        d_abbr.appendChild(d_icon_div);
        p_icon_div.appendChild(p_icon);
        e_icon_div.appendChild(e_icon);
        acc_icon_div.appendChild(acc_icon);
        d_icon_div.appendChild(d_icon);


        name_p.innerHTML = `${obj.name}`
        x.appendChild(base_div);
      }
    }

    if (x.childElementCount <= 0) {
      document.getElementById("base-no-results").style.display = "flex";
    } else {
      document.getElementById("base-no-results").style.display = "none";
    }
  }



  document.getElementById("backbtn").onclick = function () { // back button in view ig
    document.querySelector('.main').style.display = "flex";
    document.querySelector('.container').style.display = "none";

    document.getElementById("space-maker").style.display = "block";

    var elements = document.getElementsByClassName("acc-chip");

    while (elements.length > 0) {
      elements[0].parentNode.removeChild(elements[0]);
    }

    var elements = document.getElementsByClassName("v-phone-container");

    while (elements.length > 0) {
      elements[0].parentNode.removeChild(elements[0]);
    }

    var elements = document.getElementsByClassName("v-email-container");

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

    var phoneElements = document.getElementsByClassName("phone-container");

    while (phoneElements.length > 0) {
      phoneElements[0].parentNode.removeChild(phoneElements[0]);
    }

    var mailElements = document.getElementsByClassName("email-container");

    while (mailElements.length > 0) {
      mailElements[0].parentNode.removeChild(mailElements[0]);
    }

    var hobbyElements = document.getElementsByClassName("hobby-container");

    while (hobbyElements.length > 0) {
      hobbyElements[0].parentNode.removeChild(hobbyElements[0]);
    }

    var clubElements = document.getElementsByClassName("club-container");

    while (clubElements.length > 0) {
      clubElements[0].parentNode.removeChild(clubElements[0]);
    }

    var ipElements = document.getElementsByClassName("ip-container");

    while (ipElements.length > 0) {
      ipElements[0].parentNode.removeChild(ipElements[0]);
    }

    const parentElement = document.querySelector(".e-accounts");
    parentElement.innerHTML = "";
  }

  document.getElementById("c-backbtn").onclick = function () {
    document.querySelector('.main').style.display = "flex";
    document.querySelector('.create-container').style.display = "none";
  }

  document.getElementById("acc-backbtn").onclick = function () { // account back button
    location.reload();
  }

  // Hobbies

  document.getElementById("c-club-add-btn").onclick = function () {
    const club_base = document.querySelector(".c-club-base");

    const club_container = document.createElement("div");
    club_container.className = "c-club-container";

    const subContainer = document.createElement("div");
    subContainer.className = "club-subcontainer";

    const club_input = document.createElement("input");
    club_input.className = "form-input e-club";
    club_input.id = "club";
    club_input.type = "text";
    club_input.placeholder = "Enter club";
    club_input.spellcheck = "false";
    club_input.required = "true";

    const del_btn_div = document.createElement("div");
    del_btn_div.className = "del-btn";

    const del_btn = document.createElement("ion-icon");
    del_btn.name = "remove-outline";

    club_base.appendChild(club_container);
    club_container.appendChild(subContainer);
    subContainer.appendChild(club_input);
    subContainer.appendChild(del_btn_div);
    del_btn_div.appendChild(del_btn);

    del_btn_div.onclick = function () {
      club_container.remove();
    }
  }

  // IPs

  document.getElementById("c-ip-add-btn").onclick = function () {
    const ip_base = document.querySelector(".c-ip-base");

    const ip_container = document.createElement("div");
    ip_container.className = "c-ip-container";

    const subContainer = document.createElement("div");
    subContainer.className = "ip-subcontainer";

    const ip_input = document.createElement("input");
    ip_input.className = "form-input e-ip";
    ip_input.id = "ip";
    ip_input.type = "text";
    ip_input.placeholder = "Enter IP";
    ip_input.spellcheck = "false";
    ip_input.required = "true";

    const del_btn_div = document.createElement("div");
    del_btn_div.className = "del-btn";

    const del_btn = document.createElement("ion-icon");
    del_btn.name = "remove-outline";

    ip_base.appendChild(ip_container);
    ip_container.appendChild(subContainer);
    subContainer.appendChild(ip_input);
    subContainer.appendChild(del_btn_div);
    del_btn_div.appendChild(del_btn);

    del_btn_div.onclick = function () {
      ip_container.remove();
    }
  }

  // Phone

  document.getElementById("c-phone-add-btn").onclick = function () {
    const phone_base = document.querySelector(".c-phone-base");

    const phone_container = document.createElement("div");
    phone_container.className = "c-phone-container";

    const subContainer = document.createElement("div");
    subContainer.className = "phone-subcontainer";

    const phone_input = document.createElement("input");
    phone_input.className = "form-input e-phone";
    phone_input.id = "phone";
    phone_input.type = "tel";
    phone_input.placeholder = "Enter phone number";
    phone_input.spellcheck = "false";
    phone_input.maxLength = "15";
    phone_input.required = "true";

    const del_btn_div = document.createElement("div");
    del_btn_div.className = "del-btn";

    const del_btn = document.createElement("ion-icon");
    del_btn.name = "remove-outline";

    phone_base.appendChild(phone_container);
    phone_container.appendChild(subContainer);
    subContainer.appendChild(phone_input);
    subContainer.appendChild(del_btn_div);
    del_btn_div.appendChild(del_btn);

    del_btn_div.onclick = function () {
      phone_container.remove();
    }
  }

  // Hobbies

  document.getElementById("c-hobby-add-btn").onclick = function () {
    const hobby_base = document.querySelector(".c-hobby-base");

    const hobby_container = document.createElement("div");
    hobby_container.className = "c-hobby-container";

    const subContainer = document.createElement("div");
    subContainer.className = "hobby-subcontainer";

    const hobby_input = document.createElement("input");
    hobby_input.className = "form-input e-hobby";
    hobby_input.id = "hobby";
    hobby_input.type = "tel";
    hobby_input.placeholder = "Enter hobby";
    hobby_input.spellcheck = "false";
    hobby_input.required = "true";

    const del_btn_div = document.createElement("div");
    del_btn_div.className = "del-btn";

    const del_btn = document.createElement("ion-icon");
    del_btn.name = "remove-outline";

    hobby_base.appendChild(hobby_container);
    hobby_container.appendChild(subContainer);
    subContainer.appendChild(hobby_input);
    subContainer.appendChild(del_btn_div);
    del_btn_div.appendChild(del_btn);

    del_btn_div.onclick = function () {
      hobby_container.remove();
    }
  }

  // IPs

  document.getElementById("c-ip-add-btn").onclick = function () {
    const ip_base = document.querySelector(".c-ip-base");

    const ip_container = document.createElement("div");
    ip_container.className = "c-ip-container";

    const subContainer = document.createElement("div");
    subContainer.className = "ip-subcontainer";

    const ip_input = document.createElement("input");
    ip_input.className = "form-input e-ip";
    ip_input.id = "ip";
    ip_input.type = "tel";
    ip_input.placeholder = "Enter IP";
    ip_input.spellcheck = "false";
    ip_input.required = "true";

    const del_btn_div = document.createElement("div");
    del_btn_div.className = "del-btn";

    const del_btn = document.createElement("ion-icon");
    del_btn.name = "remove-outline";

    ip_base.appendChild(ip_container);
    ip_container.appendChild(subContainer);
    subContainer.appendChild(ip_input);
    subContainer.appendChild(del_btn_div);
    del_btn_div.appendChild(del_btn);

    del_btn_div.onclick = function () {
      ip_container.remove();
    }
  }

  // Email

  document.getElementById("c-add-btn").onclick = function () {
    const email_base = document.querySelector(".c-email-base");

    const email_container = document.createElement("div");
    email_container.className = "c-email-container";

    const subContainer = document.createElement("div");
    subContainer.className = "c-email-subcontainer";

    const email_input = document.createElement("input");
    email_input.className = "form-input e-mail";
    email_input.id = "c-e-mail";
    email_input.type = "email";
    email_input.placeholder = "Enter email address";
    email_input.spellcheck = "false";
    email_input.maxLength = "30";



    email_input.autocomplete = "off";

    const del_btn_div = document.createElement("div");
    del_btn_div.className = "del-btn";

    const del_btn = document.createElement("ion-icon");
    del_btn.name = "remove-outline";

    email_base.appendChild(email_container);
    email_container.appendChild(email_input);
    email_container.appendChild(del_btn_div);
    del_btn_div.appendChild(del_btn);

    del_btn_div.onclick = function () {
      email_container.remove();
    }

    email_container.appendChild(subContainer);
    subContainer.appendChild(email_input);
    subContainer.appendChild(del_btn_div);
    del_btn_div.appendChild(del_btn);
  }

  // CREATE

  document.getElementById("c-savebtn").onclick = function () { // new document save button
    let totalIds = Object.keys(data).length;
    let preId = String(totalIds + 1);

    //A function to check if the data list includes that id already, if it does, it should add one until it doesnt exist
    function checkId(preId) {
      let idExists = false;

      for (let i = 0; i < totalIds; i++) {
        if (Object.keys(data)[i] == preId) {
          idExists = true;
          break;
        }
      }

      if (idExists) {
        preId = String(parseInt(preId) + 1);
        return checkId(preId);
      }
      return preId;
    }

    let id = checkId(preId);

    let name = document.querySelector(".c-name-tag").innerHTML;

    let gender = checkGender("create");

    let age = parseInt(document.querySelector(".c-age").innerHTML);

    if (age <= 0) {
      age *= -1
    }
    if (age > 120) {
      age = 120
    }

    let bday = document.querySelector(".c-bday").innerHTML;
    let address = document.querySelector(".c-address").innerHTML;

    let phoneContainers = document.querySelectorAll('.c-phone-container');
    let phoneNumbers = {};

    phoneContainers.forEach(function (container) {
      let phoneInput = container.querySelector('input');
      phoneNumbers[phoneInput.value] = {
        "number": phoneInput.value
      };
    });

    let civilstatus = checkCivilstatus("create");

    let kids = document.querySelector(".c-kids").innerHTML;

    let hobbyContainers = document.querySelectorAll('.c-hobby-container');
    let hobbies = {};

    hobbyContainers.forEach(function (container) {
      let hobbyInput = container.querySelector('input');
      hobbies[hobbyInput.value] = {
        "hobby": hobbyInput.value
      };
    });

    let occupation = document.querySelector(".c-occupation").innerHTML;
    let prevoccupation = document.querySelector(".c-prevoccupation").innerHTML;
    let education = document.querySelector(".c-education").innerHTML;

    let religion = checkReligion("create");

    let pets = document.querySelector(".c-pets").innerHTML;

    let clubContainers = document.querySelectorAll('.c-club-container');
    let clubs = {};

    clubContainers.forEach(function (container) {
      let clubInput = container.querySelector('input');
      clubs[clubInput.value] = {
        "club": clubInput.value
      };
    });

    let legal = document.querySelector(".c-legal").innerHTML;
    let political = document.querySelector(".c-political").innerHTML;
    let notes = document.querySelector(".c-notes").innerHTML;

    let emailContainers = document.querySelectorAll('.c-email-container');
    let emailAddresses = {};

    emailContainers.forEach(function (container) {
      let emailInput = container.querySelector('input');
      emailAddresses[emailInput.value] = {
        "mail": emailInput.value,
        "src": "manual"
      };
    });

    let ipContainers = document.querySelectorAll('.c-ip-container');
    let ips = {};

    ipContainers.forEach(function (container) {
      let ipInput = container.querySelector('input');
      ips[ipInput.value] = {
        "ip": ipInput.value
      };
    });

    const loadingSpinner = document.querySelector("#c-loading-spinner");
    loadingSpinner.style.display = "flex"

    fetch(apiUrl + '/person', {
      method: 'POST',
      body: JSON.stringify({ "id": id, "name": name, "gender": gender, "age": age, "bday": bday, "address": address, "phone": phoneNumbers, "civilstatus": civilstatus, "kids": kids, "hobbies": hobbies, "email": emailAddresses, "ips": ips, "occupation": occupation, "prevoccupation": prevoccupation, "education": education, "religion": religion, "pets": pets, "clubs": clubs, "legal": legal, "political": political, "notes": notes })
    }).then(function () {
      loadingSpinner.style.display = "none"
      location.reload();
    });
  }



  // EDIT

  document.getElementById("e-savebtn").onclick = async function () {
    let id = document.querySelector("#e-showid").innerHTML;

    let name = document.querySelector(".e-name-tag").value;

    let gender = checkGender("edit");
    
    let age = parseInt(document.querySelector(".e-age").innerHTML);

    if (age < 0) {
      age *= -1
    }
    if (age > 120) {
      age = 120
    }

    let bday = document.querySelector(".e-bday").innerHTML;
    let address = document.querySelector(".e-address").innerHTML;

    let phoneContainers = document.querySelectorAll('.phone-container');
    let phoneNumbers = {};

    phoneContainers.forEach(function (container) {
      let phoneInput = container.querySelector('input');
      phoneNumbers[phoneInput.value] = {
        "number": phoneInput.value
      };
    });

    let civilstatus = checkCivilstatus("edit");

    let kids = document.querySelector(".e-kids").innerHTML;

    let hobbyContainers = document.querySelectorAll('.hobby-container');
    let hobbies = {};

    hobbyContainers.forEach(function (container) {
      let hobbyInput = container.querySelector('input');
      hobbies[hobbyInput.value] = {
        "hobby": hobbyInput.value
      };
    });

    let occupation = document.querySelector(".e-occupation").innerHTML;
    let prevoccupation = document.querySelector(".e-prevoccupation").innerHTML;
    let education = document.querySelector(".e-education").innerHTML;

    let religion = checkReligion("edit");

    let pets = document.querySelector(".e-pets").innerHTML;

    let clubContainers = document.querySelectorAll('.club-container');
    let clubs = {};

    clubContainers.forEach(function (container) {
      let clubInput = container.querySelector('input');
      clubs[clubInput.value] = {
        "club": clubInput.value
      };
    });

    let legal = document.querySelector(".e-legal").innerHTML;
    let political = document.querySelector(".e-political").innerHTML;
    let notes = document.querySelector(".e-notes").innerHTML;

    let emailContainers = document.querySelectorAll('.email-container');
    let emailAddresses = {};

    emailContainers.forEach(function (container) {
      let hiddenElement = container.querySelector(".hidden-email-save");

      let hiddenElementVal = null;

      if (hiddenElement.innerHTML != "" && hiddenElement.innerHTML != null && hiddenElement.innerHTML != undefined) {
        hiddenElementVal = JSON.parse(hiddenElement.innerHTML);
      }

      let emailInput = container.querySelector('input');
      emailAddresses[emailInput.value] = {
        "mail": emailInput.value,
        "src": "manual",
        "services": hiddenElementVal
      };
    });
    
    let ipContainers = document.querySelectorAll('.ip-container');
    let ips = {};

    ipContainers.forEach(function (container) {
      let ipInput = container.querySelector('input');
      ips[ipInput.value] = {
        "ip": ipInput.value
      };
    });
    
    const loadingSpinner = document.querySelector("#e-loading-spinner");
    loadingSpinner.style.display = "flex"

    const res = await fetch(apiUrl + "/people/" + id)

    let data = await res.json();

    fetch(apiUrl + '/person', {
      method: 'POST',
      body: JSON.stringify({ "id": id, "name": name, "gender": gender, "age": age, "bday": bday, "address": address, "phone": phoneNumbers, "civilstatus": civilstatus, "kids": kids, "hobbies": hobbies, "email": emailAddresses, "ips": ips, "occupation": occupation, "prevoccupation": prevoccupation, "education": education, "religion": religion, "pets": pets, "clubs": clubs, "legal": legal, "political": political, "notes": notes, "accounts": data.accounts })
    }).then(function () {
      loadingSpinner.style.display = "none"
      location.reload();
    });
  }
}

main()
