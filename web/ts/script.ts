import { saveAsFile, checkDropdownValue, getDropdownElementIndex, apiCall, DropdownType } from "./framework.js";
import * as person from "../ts-gen/person.js";




function mkDropdown(name: DropdownType, value: string, child: string) { // NAME AND CHILD NOT USER CONTROLLED!!!!!! 

    if (value != "" && value != undefined) {
      const valueSelect = document.querySelector(".edit-container > .components > .scroll-box > div:nth-child(" + child + ") > ." + name + "-select") as HTMLElement; // not unsave because name and child are not user controlled
      const selectItems = valueSelect.querySelector(".select-items") as HTMLElement;
      const selectSelected = valueSelect.querySelector(".select-selected") as HTMLElement;

      const dropdownValueIndex: string = getDropdownElementIndex(name, value);

      if (dropdownValueIndex != "" && dropdownValueIndex != undefined) {
        const dropdownElement = selectItems.children[parseInt(dropdownValueIndex)];

        selectSelected.innerHTML = translateRawWord(value)!;
        //dropdownElement.className = "same-as-selected";
      }
    } 
   // else {
   //   console.log("value-" + name + ": " + value)
   // }

}
function init<T>(value: { [key: string]: T }) :  { [key: string]: T } {
  if (value === undefined) {
    value = {} as { [key: string]: T };
  }
  return value;
}
function mkList<T>(name: string, value: { [key: string]: T }, fieldname: keyof T) {
  value = init<T>(value);

    const listBase = document.querySelector(".e-"+ name + "-base") as HTMLDivElement;

    if (Object.keys(value).length >= 1) {
      for (const [_, item] of Object.entries(value)) {
        const itemVar = (item as T)

        const container = document.createElement("div");
        container.className = name + "-container";

        const subContainer = document.createElement("div");
        subContainer.className = ""+ name + "-subcontainer";

        const value_input = document.createElement("input") as HTMLInputElement;
        value_input.className = "form-input " + name;
        value_input.id = "e-" + name;
        value_input.placeholder = "Enter " + name;
        value_input.spellcheck = false;
        value_input.setAttribute("lng-tag", "enter_" + name)
        value_input.value = String(itemVar[fieldname]);

        const del_btn_div = document.createElement("div");
        del_btn_div.className = "del-btn";

        const del_btn = document.createElement("ion-icon") as IonIconElement;
        del_btn.name = "remove-outline";

        container.appendChild(subContainer);
        subContainer.appendChild(value_input);
        listBase.appendChild(container);
        subContainer.appendChild(del_btn_div);
        del_btn_div.appendChild(del_btn);

        del_btn.onclick = function () {
          container.remove();
        }

        // refreshTranslation();
      };
    }

    document.getElementById(name + "-add-btn")!.onclick = function () {
      const list_container = document.createElement("div");
      list_container.className = ""+ name + "-container";

      const subContainer = document.createElement("div");
      subContainer.className = ""+ name + "-subcontainer";

      const value_input = document.createElement("input");
      value_input.className = "form-input e-" + name;
      value_input.id = name;
      value_input.type = "text";
      value_input.placeholder = "Enter "+ name;
      value_input.setAttribute("lng-tag", "enter_" + name)
      value_input.spellcheck = false;
      value_input.required = true;

      const del_btn_div = document.createElement("div");
      del_btn_div.className = "del-btn";

      const del_btn = document.createElement("ion-icon") as IonIconElement;
      del_btn.name = "remove-outline";

      listBase.appendChild(list_container);
      list_container.appendChild(subContainer);
      subContainer.appendChild(value_input);
      subContainer.appendChild(del_btn_div);
      del_btn_div.appendChild(del_btn);

      del_btn_div.onclick = function () {
        list_container.remove();
      }

      // refreshTranslation();
    }
}

class Person extends person.Person {
  Post(loadingSpinner?: HTMLDivElement): void {
    const requestOptions = {
      method: "POST",
      body: JSON.stringify(this),
    };

    fetch(apiCall("/person"), requestOptions)
      .then(() => {
        if (loadingSpinner) {
          loadingSpinner.style.display = "none";
        }
        location.reload();
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }

  Edit(): void {
    let obj = this;
    mainContainer.style.display = "none";
    editContainer.style.display = "flex";

    editShowID.innerHTML = obj.id;

    editNameTag.value = obj.name;

    mkDropdown("gender",obj.gender, "1");
    //if (obj.gender != "") {
    //  const genderSelect = document.querySelector(".edit-container > .components > .scroll-box > div:nth-child(1) > .gender-select") as HTMLElement;
    //  const selectItems = genderSelect.querySelector(".select-items") as HTMLElement;
    //  const selectSelected = genderSelect.querySelector(".select-selected") as HTMLElement;

    //  const genderIndex: string = getDropdownElementIndex("gender", obj.gender);

    //  if (genderIndex != "" && genderIndex != undefined) {
    //    const genderElement = selectItems.children[parseInt(genderIndex)];

    //    selectSelected.innerHTML = translateRawWord(obj.gender)!;
    //    genderElement.className = "same-as-selected";
    //  }
    //}

    mkDropdown("ethnicity",obj.ethnicity, "2")

    editAge.innerHTML = obj.age.toString() || "";
    editBday.innerHTML = obj.bday;
    editAddress.innerHTML = obj.address;

    // Phone


    obj.phone = init<typeof obj.phone[string]>(obj.phone);
    const phoneBase = document.querySelector(".phone-base") as HTMLDivElement;

    if (Object.keys(obj.phone).length >= 1) {
      for (const [_, phone] of Object.entries(obj.phone)) {
        const phoneVar = (phone as { number: string, valid: boolean, phoneinfoga: { Country: string } })

        const container = document.createElement("div");
        container.className = "phone-container";

        const subContainer = document.createElement("div");
        subContainer.className = "phone-subcontainer";

        const phone_input = document.createElement("input");
        phone_input.className = "form-input phone";
        phone_input.id = "e-phone";
        phone_input.type = "tel";
        phone_input.placeholder = "Enter phone number";
        phone_input.spellcheck = false;
        phone_input.setAttribute("lng-tag", "enter_phone_number")
        phone_input.required = true;
        phone_input.value = phoneVar.number;

        const del_btn_div = document.createElement("div");
        del_btn_div.className = "del-btn";

        const del_btn = document.createElement("ion-icon") as IonIconElement;
        del_btn.name = "remove-outline";

        container.appendChild(subContainer);
        subContainer.appendChild(phone_input);
        phoneBase.appendChild(container);
        subContainer.appendChild(del_btn_div);
        del_btn_div.appendChild(del_btn);


        del_btn.onclick = function () {
          container.remove();
        }

        refreshTranslation();
      };
    }
    document.getElementById("phone-add-btn")!.onclick = function () {
      const phone_container = document.createElement("div");
      phone_container.className = "phone-container";

      const subContainer = document.createElement("div");
      subContainer.className = "phone-subcontainer";

      const phone_input = document.createElement("input");
      phone_input.className = "form-input e-phone";
      phone_input.id = "phone";
      phone_input.type = "tel";
      phone_input.placeholder = "Enter phone number";
      phone_input.spellcheck = false;
      phone_input.setAttribute("lng-tag", "enter_phone_number")
      //phone_input.maxLength = "15"; // FIXME some formattings can have more then 15 chars.
      phone_input.required = true;

      const del_btn_div = document.createElement("div");
      del_btn_div.className = "del-btn";

      const del_btn = document.createElement("ion-icon") as IonIconElement;
      del_btn.name = "remove-outline";

      phoneBase.appendChild(phone_container);
      phone_container.appendChild(subContainer);
      subContainer.appendChild(phone_input);
      subContainer.appendChild(del_btn_div);
      del_btn_div.appendChild(del_btn);

      del_btn_div.onclick = function () {
        phone_container.remove();
      }

      refreshTranslation();
    }



    mkDropdown("civilstatus",obj.civilstatus, "7");

    editKids.innerHTML = obj.kids;

    mkList<typeof obj.hobbies[string]>("hobby", obj.hobbies, "hobby");

    // const hobbyBase = document.querySelector(".e-hobby-base") as HTMLDivElement;

    // if (Object.keys(obj.hobbies).length >= 1) {
    //   for (const [_, hobby] of Object.entries(obj.hobbies)) {
    //     const hobbyVar = (hobby as { hobby: string })

    //     const container = document.createElement("div");
    //     container.className = "hobby-container";

    //     const subContainer = document.createElement("div");
    //     subContainer.className = "hobby-subcontainer";

    //     const hobby_input = document.createElement("input") as HTMLInputElement;
    //     hobby_input.className = "form-input hobby";
    //     hobby_input.id = "e-hobby";
    //     hobby_input.placeholder = "Enter hobby";
    //     hobby_input.spellcheck = false;
    //     hobby_input.setAttribute("lng-tag", "enter_hobby")
    //     hobby_input.value = hobbyVar.hobby;

    //     const del_btn_div = document.createElement("div");
    //     del_btn_div.className = "del-btn";

    //     const del_btn = document.createElement("ion-icon") as IonIconElement;
    //     del_btn.name = "remove-outline";

    //     container.appendChild(subContainer);
    //     subContainer.appendChild(hobby_input);
    //     hobbyBase.appendChild(container);
    //     subContainer.appendChild(del_btn_div);
    //     del_btn_div.appendChild(del_btn);

    //     del_btn.onclick = function () {
    //       container.remove();
    //     }

    //     refreshTranslation();
    //   };
    // }

    // document.getElementById("hobby-add-btn")!.onclick = function () {
    //   const hobby_container = document.createElement("div");
    //   hobby_container.className = "hobby-container";

    //   const subContainer = document.createElement("div");
    //   subContainer.className = "hobby-subcontainer";

    //   const hobby_input = document.createElement("input");
    //   hobby_input.className = "form-input e-hobby";
    //   hobby_input.id = "hobby";
    //   hobby_input.type = "text";
    //   hobby_input.placeholder = "Enter hobby";
    //   hobby_input.setAttribute("lng-tag", "enter_hobby")
    //   hobby_input.spellcheck = false;
    //   hobby_input.required = true;

    //   const del_btn_div = document.createElement("div");
    //   del_btn_div.className = "del-btn";

    //   const del_btn = document.createElement("ion-icon") as IonIconElement;
    //   del_btn.name = "remove-outline";

    //   hobbyBase.appendChild(hobby_container);
    //   hobby_container.appendChild(subContainer);
    //   subContainer.appendChild(hobby_input);
    //   subContainer.appendChild(del_btn_div);
    //   del_btn_div.appendChild(del_btn);

    //   del_btn_div.onclick = function () {
    //     hobby_container.remove();
    //   }

    //   refreshTranslation();
    // }

    editOccupation.innerHTML = obj.occupation;
    editPrevOccupation.innerHTML = obj.prevoccupation;
    editEducation.innerHTML = obj.education;

    mkDropdown("religion",obj.religion, "15");
    editPets.innerHTML = obj.pets;


    mkList<typeof obj.clubs[string]>("club", obj.clubs, "club");

    // const clubBase = document.querySelector(".e-club-base") as HTMLDivElement;

    // if (Object.keys(obj.clubs).length >= 1) {
    //   for (const [_, club] of Object.entries(obj.clubs)) {
    //     const clubVar = (club as { club: string })

    //     const container = document.createElement("div");
    //     container.className = "club-container";

    //     const subContainer = document.createElement("div");
    //     subContainer.className = "club-subcontainer";

    //     const club_input = document.createElement("input");
    //     club_input.className = "form-input club";
    //     club_input.id = "e-club";
    //     club_input.type = "text";
    //     club_input.placeholder = "Enter club";
    //     club_input.spellcheck = false;
    //     club_input.setAttribute("lng-tag", "enter_club")
    //     club_input.value = clubVar.club;

    //     const del_btn_div = document.createElement("div");
    //     del_btn_div.className = "del-btn";

    //     const del_btn = document.createElement("ion-icon") as IonIconElement;
    //     del_btn.name = "remove-outline";

    //     container.appendChild(subContainer);
    //     subContainer.appendChild(club_input);
    //     clubBase.appendChild(container);
    //     subContainer.appendChild(del_btn_div);
    //     del_btn_div.appendChild(del_btn);

    //     del_btn.onclick = function () {
    //       container.remove();
    //     }

    //     refreshTranslation();
    //   };
    // }

    // document.getElementById("club-add-btn")!.onclick = function () {
    //   const club_container = document.createElement("div");
    //   club_container.className = "club-container";

    //   const subContainer = document.createElement("div");
    //   subContainer.className = "club-subcontainer";

    //   const club_input = document.createElement("input");
    //   club_input.className = "form-input e-club";
    //   club_input.id = "club";
    //   club_input.type = "text";
    //   club_input.placeholder = "Enter club";
    //   club_input.spellcheck = false;
    //   club_input.setAttribute("lng-tag", "enter_club")
    //   club_input.required = true;

    //   const del_btn_div = document.createElement("div");
    //   del_btn_div.className = "del-btn";

    //   const del_btn = document.createElement("ion-icon") as IonIconElement;
    //   del_btn.name = "remove-outline";

    //   clubBase.appendChild(club_container);
    //   club_container.appendChild(subContainer);
    //   subContainer.appendChild(club_input);
    //   subContainer.appendChild(del_btn_div);
    //   del_btn_div.appendChild(del_btn);

    //   del_btn_div.onclick = function () {
    //     club_container.remove();
    //   }

    //   refreshTranslation();
    // }

    editLegal.innerHTML = obj.legal;
    editPolitical.innerHTML = obj.political;

    // Sources


    mkList<typeof obj.sources[string]>("source", obj.sources, "url");

    // const sourceBase = document.querySelector(".e-source-base") as HTMLDivElement;

    // if (Object.keys(obj.sources).length >= 1) {
    //   for (const [_, url] of Object.entries(obj.sources)) {
    //     const sourceVar = (url as { url: string })

    //     const container = document.createElement("div");
    //     container.className = "source-container";

    //     const subContainer = document.createElement("div");
    //     subContainer.className = "source-subcontainer";

    //     const source_input = document.createElement("input");
    //     source_input.className = "form-input source";
    //     source_input.id = "e-source";
    //     source_input.type = "text";
    //     source_input.placeholder = "Enter source";
    //     source_input.spellcheck = false;
    //     source_input.setAttribute("lng-tag", "enter_source")
    //     source_input.value = sourceVar.url;

    //     const del_btn_div = document.createElement("div");
    //     del_btn_div.className = "del-btn";

    //     const del_btn = document.createElement("ion-icon") as IonIconElement;
    //     del_btn.name = "remove-outline";

    //     container.appendChild(subContainer);
    //     subContainer.appendChild(source_input);
    //     sourceBase.appendChild(container);
    //     subContainer.appendChild(del_btn_div);
    //     del_btn_div.appendChild(del_btn);

    //     del_btn.onclick = function () {
    //       container.remove();
    //     }

    //     refreshTranslation();
    //   };
    // }

    // document.getElementById("source-add-btn")!.onclick = function () {
    //   const source_container = document.createElement("div");
    //   source_container.className = "source-container";

    //   const subContainer = document.createElement("div");
    //   subContainer.className = "source-subcontainer";

    //   const source_input = document.createElement("input");
    //   source_input.className = "form-input e-source";
    //   source_input.id = "source";
    //   source_input.type = "text";
    //   source_input.placeholder = "Enter source";
    //   source_input.spellcheck = false;
    //   source_input.setAttribute("lng-tag", "enter_source")
    //   source_input.required = true;

    //   const del_btn_div = document.createElement("div");
    //   del_btn_div.className = "del-btn";

    //   const del_btn = document.createElement("ion-icon") as IonIconElement;
    //   del_btn.name = "remove-outline";

    //   sourceBase.appendChild(source_container);
    //   source_container.appendChild(subContainer);
    //   subContainer.appendChild(source_input);
    //   subContainer.appendChild(del_btn_div);
    //   del_btn_div.appendChild(del_btn);

    //   del_btn_div.onclick = function () {
    //     source_container.remove();
    //   }

    //   refreshTranslation();
    // }

    editNotes.innerHTML = obj.notes;

    // IPs


    //mkList<typeof obj.ips[string]>("ip", obj.ips, "ip");
    
    obj.ips = init<typeof obj.ips[string]>(obj.ips);
    const ipBase = document.querySelector(".e-ip-base") as HTMLDivElement;

    if (Object.keys(obj.ips).length >= 1) {
      for (const [_, ip] of Object.entries(obj.ips)) {
        const ipVar = (ip as { ip: string })

        const container = document.createElement("div");
        container.className = "ip-container";

        const subContainer = document.createElement("div");
        subContainer.className = "ip-subcontainer";

        const ip_input = document.createElement("input");
        ip_input.className = "form-input ip";
        ip_input.id = "e-ip";
        ip_input.type = "text";
        ip_input.placeholder = "Enter IP";
        ip_input.spellcheck = false;
        ip_input.setAttribute("lng-tag", "enter_ip")
        ip_input.value = ipVar.ip;

        const del_btn_div = document.createElement("div");
        del_btn_div.className = "del-btn";

        const del_btn = document.createElement("ion-icon") as IonIconElement;
        del_btn.name = "remove-outline";

        container.appendChild(subContainer);
        subContainer.appendChild(ip_input);
        ipBase.appendChild(container);
        subContainer.appendChild(del_btn_div);
        del_btn_div.appendChild(del_btn);

        del_btn.onclick = function () {
          container.remove();
        }

        refreshTranslation();
      };
    }

    document.getElementById("ip-add-btn")!.onclick = function () {
      const ip_container = document.createElement("div");
      ip_container.className = "ip-container";

      const subContainer = document.createElement("div");
      subContainer.className = "ip-subcontainer";

      const ip_input = document.createElement("input");
      ip_input.className = "form-input e-ip";
      ip_input.id = "ip";
      ip_input.type = "text";
      ip_input.placeholder = "Enter IP";
      ip_input.spellcheck = false;
      ip_input.setAttribute("lng-tag", "enter_ip")
      ip_input.required = true;

      const del_btn_div = document.createElement("div");
      del_btn_div.className = "del-btn";

      const del_btn = document.createElement("ion-icon") as IonIconElement;
      del_btn.name = "remove-outline";

      ipBase.appendChild(ip_container);
      ip_container.appendChild(subContainer);
      subContainer.appendChild(ip_input);
      subContainer.appendChild(del_btn_div);
      del_btn_div.appendChild(del_btn);

      del_btn_div.onclick = function () {
        ip_container.remove();
      }

      refreshTranslation();
    }

    // Email


    obj.email = init<typeof obj.email[string]>(obj.email);
    const emailBase = document.querySelector(".email-base") as HTMLDivElement;

    if (Object.keys(obj.email).length >= 1) {
      for (const [_, email] of Object.entries(obj.email)) {
        const emailVar = (email as { mail: string, services: {} });

        const container = document.createElement("div");
        container.className = "email-container";

        const subContainer = document.createElement("div");
        subContainer.className = "email-subcontainer";

        const email_input = document.createElement("input");
        email_input.className = "form-input e-mail";
        email_input.id = "e-mail";
        email_input.type = "email";
        email_input.placeholder = "Enter email address";
        email_input.spellcheck = false;
        email_input.setAttribute("lng-tag", "enter_email_address")
        email_input.required = true;
        email_input.value = emailVar.mail;

        const del_btn_div = document.createElement("div");
        del_btn_div.className = "del-btn";

        const del_btn = document.createElement("ion-icon") as IonIconElement;
        del_btn.name = "remove-outline";

        container.appendChild(subContainer);
        subContainer.appendChild(email_input);
        emailBase.appendChild(container);
        subContainer.appendChild(del_btn_div);
        del_btn_div.appendChild(del_btn);

        if (emailVar.services != undefined && emailVar.services != null && emailVar.services != "") {
          const hidden_email_save = document.createElement("p");
          hidden_email_save.className = "hidden-email-save";

          hidden_email_save.innerHTML = JSON.stringify(emailVar.services);
          container.appendChild(hidden_email_save);
        }

        del_btn.onclick = function () {
          container.remove();
        }

        refreshTranslation();
      };
    }



    document.getElementById("email-add-btn")!.onclick = function () {
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
      email_input.spellcheck = false;
      email_input.setAttribute("lng-tag", "enter_email_address")
      email_input.required = true;

      const del_btn_div = document.createElement("div");
      del_btn_div.className = "del-btn";

      const del_btn = document.createElement("ion-icon") as IonIconElement;
      del_btn.name = "remove-outline";

      emailBase.appendChild(email_container);
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

      refreshTranslation();
    }

    // Accounts


    obj.accounts= init<typeof obj.accounts[string]>(obj.accounts);
    if (Object.keys(obj.accounts).length != 0 && obj.accounts != null) {
      for (const [_, accObj] of Object.entries(obj.accounts)) {
        const accVar = (accObj as { service: string, id: string, username: string, url: string, profilePicture: { [key: number]: { img: string, img_hash: number } }, bio: { [key: number]: { bio: string } } });

        //let accObj = obj.accounts[i];

        // Creating elements

        const base_div = document.createElement("div"); // Outer div
        base_div.className = "acc-chip";

        const pfp_img = document.createElement("img"); // Pfp img
        pfp_img.className = "userPfp";

        if (accVar.profilePicture != null) {
          pfp_img.src = "data:image/png;base64," + accVar.profilePicture["1"].img;
        } else {
          pfp_img.src = "https://as2.ftcdn.net/v2/jpg/03/32/59/65/1000_F_332596535_lAdLhf6KzbW6PWXBWeIFTovTii1drkbT.jpg"
        }

        const info_div = document.createElement("div"); // Info div
        info_div.className = "info-container";

        const icon_space = document.createElement("div");
        icon_space.className = "icon-space";

        const service_p = document.createElement("a");
        service_p.className = "serviceName";
        service_p.innerHTML = accVar.service;
        service_p.href = accVar.url;
        service_p.target = "_blank";

        const name_p = document.createElement("a");
        name_p.className = "userName";
        name_p.innerHTML = accVar.username;
        name_p.href = accVar.url;
        name_p.target = "_blank";

        document.querySelector(".e-accounts")!.appendChild(base_div);
        base_div.appendChild(pfp_img);
        base_div.appendChild(info_div);
        base_div.appendChild(icon_space);
        info_div.appendChild(service_p);
        info_div.appendChild(name_p);

        if (accVar.service.toLowerCase() == "github") { // If the service is github, add a deep investigation button
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

          const del_btn = document.createElement("ion-icon") as IonIconElement;
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


            const res = await fetch(apiCall("/deep/github/" + accVar.username))
            let data = await res.json();

            loadingSpinner.remove();

            const deepInvResIcon = document.createElement("ion-icon") as IonIconElement;
            deepInvResIcon.className = "deepInvResIcon";

            icon_space.appendChild(deepInvResIcon);

            if (data != null && data != "{}" && res.status == 200) {
              deep_btn_txt.innerHTML = "Deep Investigation";

              deepInvResIcon.name = "checkmark-outline";
              deepInvResIcon.style.filter = "drop-shadow(0.3rem 0.3rem 0.2rem var(--greyLight-2)) drop-shadow(-0.2rem -0.2rem 0.5rem var(--white));"

              for (const [i, _] of Object.entries(data)) {
                let obj = data[i];

                const email_container = document.createElement("div");
                email_container.className = "email-container";

                const subContainer = document.createElement("div");
                subContainer.className = "email-subcontainer";

                const email_input = document.createElement("input");
                email_input.className = "form-input e-mail";
                email_input.id = "e-mail";
                email_input.type = "email";
                email_input.placeholder = "Enter email address";
                email_input.spellcheck = false;
                email_input.setAttribute("lng-tag", "enter_email_address")
                email_input.value = obj.mail;

                const del_btn_div = document.createElement("div");
                del_btn_div.className = "del-btn";

                const del_btn = document.createElement("ion-icon") as IonIconElement;
                del_btn.name = "remove-outline";

                const hidden_email_save = document.createElement("p");
                hidden_email_save.className = "hidden-email-save";

                hidden_email_save.innerHTML = JSON.stringify(obj.services);

                emailBase.appendChild(email_container);
                email_container.appendChild(subContainer);
                subContainer.appendChild(email_input);
                subContainer.appendChild(del_btn_div);
                del_btn_div.appendChild(del_btn);
                email_container.appendChild(hidden_email_save);

                del_btn_div.onclick = function () {
                  email_container.remove();
                }

                refreshTranslation();
              }
            } else if (res.status == 403 && data["fatal"] == "rate limited") {
              deepInvResIcon.name = "timer-outline";
              deepInvResIcon.style.filter = "drop-shadow(0.3rem 0.3rem 0.15rem var(--greyLight-2)) drop-shadow(-0.2rem -0.2rem 0.5rem var(--white));"
            } else {
              deepInvResIcon.name = "close-outline";
              deepInvResIcon.style.filter = "drop-shadow(0.3rem 0.3rem 0.2rem var(--greyLight-2)) drop-shadow(-0.2rem -0.2rem 0.5rem var(--white));"
            }
          }


          del_btn_div.onclick = function () {
            fetch(apiCall("/people/" + document.querySelector("#e-showid")!.innerHTML + "/accounts/" + accVar.service + "-" + accVar.username + "/delete"), {
              method: "GET",
              mode: "no-cors"
            });

            base_div.remove();
            // TODO Add stuff here
          }
        } else {
          const del_btn_div = document.createElement("div");
          del_btn_div.className = "delAccBtn btn btn-secondary";

          const del_btn = document.createElement("ion-icon") as IonIconElement;
          del_btn.name = "remove-outline";

          base_div.appendChild(del_btn_div);
          del_btn_div.appendChild(del_btn);

          del_btn_div.onclick = function () {
            fetch(apiCall("/people/" + document.querySelector("#e-showid")!.innerHTML + "/accounts/" + accVar.service + "-" + accVar.username + "/delete"), {
              method: "GET",
              mode: "no-cors"
            });

            base_div.remove();

            // TODO Add stuff here
          }
        }

        if (accVar.bio != null) {
          const bio_p = document.createElement("p");
          bio_p.className = "userBio";
          bio_p.innerHTML = accVar.bio["1"].bio;

          info_div.appendChild(bio_p);
        }
      }
    }
    return
  }
}






const searchBar = document.getElementById("searchbar");


// Listen for messages on the broadcast channel
const channel = new BroadcastChannel("seekr-channel");

channel.addEventListener("message", (event) => {
  if (event.data.type === "theme") {
    const theme = event.data.theme;

    document.documentElement.setAttribute("data-theme", theme);
  } else if (event.data.type === "language") {
    translate()
  }
});

if (!localStorage.getItem("language")) {
  setLanguage("en"); // Default language
}

// Interface for IonIcons
interface IonIconElement extends HTMLElement {
  name: string;
}

// Variables for HTML elements

// General

const mainContainer = document.querySelector(".main") as HTMLDivElement;
const container = document.querySelector(".container") as HTMLDivElement;
const editContainer = document.querySelector(".edit-container") as HTMLDivElement;
const createContainer = document.querySelector(".create-container") as HTMLDivElement;

// View

const viewShowId = document.querySelector("#v-showid") as HTMLParagraphElement;
const viewNameTag = document.querySelector(".name-tag") as HTMLInputElement;

const viewGender = document.querySelector(".gender") as HTMLParagraphElement;
const viewEthnicity = document.querySelector(".ethnicity") as HTMLParagraphElement;
const viewAge = document.querySelector(".age") as HTMLParagraphElement;
const viewBday = document.querySelector(".bday") as HTMLParagraphElement;
const viewAddress = document.querySelector(".address") as HTMLParagraphElement;
const viewCivilStatus = document.querySelector(".civilstatus") as HTMLParagraphElement;
const viewKids = document.querySelector(".kids") as HTMLParagraphElement;
const viewOccupation = document.querySelector(".occupation") as HTMLParagraphElement;
const viewPrevOccupation = document.querySelector(".prevoccupation") as HTMLParagraphElement;
const viewEducation = document.querySelector(".education") as HTMLParagraphElement;
const viewReligion = document.querySelector(".religion") as HTMLParagraphElement;
const viewPets = document.querySelector(".pets") as HTMLParagraphElement;
const viewLegal = document.querySelector(".legal") as HTMLParagraphElement;
const viewPolitical = document.querySelector(".political") as HTMLParagraphElement;
const viewNotes = document.getElementById("notes") as HTMLDivElement;

// Edit

const editShowID = document.querySelector("#e-showid") as HTMLParagraphElement;
const editNameTag = document.querySelector(".e-name-tag") as HTMLInputElement;

const editAge = document.querySelector(".e-age") as HTMLInputElement;
const editBday = document.querySelector(".e-bday") as HTMLInputElement;
const editAddress = document.querySelector(".e-address") as HTMLInputElement;
const editKids = document.querySelector(".e-kids") as HTMLInputElement;
const editOccupation = document.querySelector(".e-occupation") as HTMLInputElement;
const editPrevOccupation = document.querySelector(".e-prevoccupation") as HTMLInputElement;
const editEducation = document.querySelector(".e-education") as HTMLInputElement;
const editPets = document.querySelector(".e-pets") as HTMLInputElement;
const editLegal = document.querySelector(".e-legal") as HTMLInputElement;
const editPolitical = document.querySelector(".e-political") as HTMLInputElement;
const editNotes = document.getElementById("e-notes") as HTMLDivElement;

// Create

const createNameTag = document.querySelector(".c-name-tag") as HTMLInputElement;

// Accs

const accNameTag = document.getElementById("acc-name-tag") as HTMLInputElement;
const accContainer = document.querySelector(".acc-container") as HTMLDivElement;
const accSearchChip = document.getElementById("acc-search-chip") as HTMLDivElement;
const accNoResults = document.getElementById("acc-no-results") as HTMLDivElement;
const accScrollBox = document.getElementById("acc-scroll-box") as HTMLDivElement;
const accAccounts = document.getElementById("accounts") as HTMLDivElement;

const accLoadingSpinner = document.getElementById("loading-spinner") as HTMLDivElement;

async function getData(): Promise<object> {
  const res = await fetch(apiCall("/"));

  let data = await res.json();

  return data; 
}

searchBar!.addEventListener("keyup", searchEntries);

document.getElementById("savemdbtn")!.onclick = async function () {
  const getId = document.getElementById("v-showid") as HTMLParagraphElement;
  const getName = document.getElementById("name-tag") as HTMLInputElement;

  const request = await fetch(apiCall("/people/" + getId!.innerHTML + "/markdown"));
  const textToSave = await request.json();


  saveAsFile(textToSave.markdown, getName!.value.toLowerCase().replace(/ /g, "") + ".md");
}

// NEW BUTTON
document.getElementById("newbtn")!.onclick = async function () {

  const data = await getData() as object[];

  let totalIds = Object.keys(data).length;
  let preId = String(totalIds + 1);


  function checkId(preId: string): string {
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

  let obj = new Person();
  obj.id = checkId(preId);
  obj.name = "";
  // obj = replaceUndefinedWithEmpty<typeof obj>(obj);
  obj.Edit();
  // mainContainer.style.display = "none";
  // createContainer.style.display = "flex";
}



document.getElementById("exportbtn")!.onclick = async function () {
  const data = await getData() as object[];

  saveAsFile(JSON.stringify(data), "data.json");
}

function createCards(obj: Person) {
  let x = document.querySelector('#list-holder')!;

  // Basic

  const base_div = document.createElement("div");
  base_div.className = "chip";

  const hitbox_abbr = document.createElement("abbr");
  hitbox_abbr.title = "View"
  hitbox_abbr.className = "hitbox-abbr";

  const hitbox_div = document.createElement("div");
  hitbox_div.className = "hitbox";

  const p_icon_div = document.createElement("div");
  p_icon_div.className = "chip-icon";

  const p_icon = document.createElement("ion-icon");
  p_icon.className = "icon"
  p_icon.setAttribute("name", "person");

  const txt_div = document.createElement("div");
  txt_div.className = "text-container";

  const name_p = document.createElement("p");
  name_p.className = "card-text";

  // Edit

  const e_icon_div = document.createElement("div");
  e_icon_div.className = "chip-edit";

  const e_abbr = document.createElement("abbr");
  e_abbr.title = "Edit"

  const e_icon = document.createElement("ion-icon");
  e_icon.className = "icon"
  e_icon.setAttribute("name", "create-outline");

  // Accounts

  const acc_icon_div = document.createElement("div");
  acc_icon_div.className = "chip-acc";

  const acc_abbr = document.createElement("abbr");
  acc_abbr.title = "Add Accounts"

  const acc_icon = document.createElement("ion-icon");
  acc_icon.className = "icon"
  acc_icon.setAttribute("name", "person-circle-outline");

  // Delete

  const d_icon_div = document.createElement("div");
  d_icon_div.className = "chip-delete";

  const d_abbr = document.createElement("abbr");
  d_abbr.title = "Delete"

  const d_icon = document.createElement("ion-icon");
  d_icon.className = "icon"
  d_icon.setAttribute("name", "trash-outline");

  d_icon_div.onclick = function () {
    fetch(apiCall("/people/" + obj.id + "/delete"), {
      method: "GET",
      mode: "no-cors"
    }).then(function () {
      location.reload();
    });
  }

  acc_icon_div.onclick = function () {
    editShowID.innerHTML = obj.id;
    mainContainer.style.display = "none";
    accContainer.style.display = "flex";
  }

  accSearchChip.onclick = search;

  accNameTag.onkeypress = function (event) {
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
    if (accNameTag.value == "") {
      return;
    }
    // Check if the button is enabled
    if (!isButtonEnabled) {
      return;
    }

    // Disable the button
    isButtonEnabled = false;

    accLoadingSpinner.style.display = "inline-block";

    // Set the flag to indicate that a request is in progress
    const response = await fetch(apiCall('/getAccounts/' + accNameTag.value));
    const data = await response.json();

    const term_container = document.createElement("div");
    term_container.className = "term-container";

    const term_header = document.createElement("p");
    term_header.className = "term-header";
    term_header.textContent = accNameTag.value;

    term_container.appendChild(term_header);

    if (data != null && Object.entries(data).length >= 1) {
      accNoResults.style.display = "none";
      accScrollBox.style.display = "block";

      const row_div = document.createElement("div");
      row_div.className = "acc-row";

      accAccounts.appendChild(row_div);


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
        reject_p.className = "not-selectable";

        const accept_btn = document.createElement("div");
        accept_btn.id = "acc-acceptbtn";
        accept_btn.className = "btn btn-secondary";

        const accept_p = document.createElement("p");
        accept_p.innerHTML = "Accept";
        accept_p.className = "not-selectable";

        manage_acc_chip.appendChild(btn_container);
        btn_container.appendChild(reject_btn);
        btn_container.appendChild(accept_btn);
        reject_btn.appendChild(reject_p);
        accept_btn.appendChild(accept_p);




        accept_btn.onclick = async function () {
          // Check if accObj.service and accObj.username are also in accounts object at obj.accounts
          let getId = editShowID.innerHTML

          const res = await fetch(apiCall("/people/" + getId));

          let data = await res.json() as Person;

          data.accounts[accObj.service + "-" + accObj.username] = accObj;
          fetch(apiCall("/person"), {
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

      if (accAccounts.childElementCount <= 0) {
        accNoResults.style.display = "flex";
        accScrollBox.style.display = "none";
      }
    }

    accLoadingSpinner.style.display = "none";
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

  hitbox_div.onclick = async function () {
    mainContainer.style.display = "none";
    container.style.display = "flex";

    viewShowId.innerHTML = obj.id;

    viewNameTag.value = obj.name;

    viewGender.innerHTML = translateRawWord("Gender:")! + " " + translateRawWord(obj.gender);
    viewEthnicity.innerHTML = translateRawWord("Ethnicity:")! + " " + translateRawWord(obj.ethnicity);
    viewAge.innerHTML = translateRawWord("Age:")! + " " + obj.age;
    viewBday.innerHTML = translateRawWord("Birthday:")! + " " + obj.bday;
    viewAddress.innerHTML = translateRawWord("Address:")! + " " + obj.address;
    viewCivilStatus.innerHTML = translateRawWord("Civil status:")! + " " + translateRawWord(obj.civilstatus);
    viewKids.innerHTML = translateRawWord("Kids:") + " " + (obj.kids || "");
    viewOccupation.innerHTML = translateRawWord("Occupation:")! + " " + obj.occupation;
    viewPrevOccupation.innerHTML = translateRawWord("Previous Occupation:")! + obj.prevoccupation;
    viewEducation.innerHTML = translateRawWord("Education:")! + " " + obj.education;
    viewReligion.innerHTML = translateRawWord("Religion:")! + " " + translateRawWord(obj.religion);
    viewPets.innerHTML = translateRawWord("Pets:")! + " " + obj.pets;
    viewLegal.innerHTML = translateRawWord("Legal:")! + " " + obj.legal;
    viewPolitical.innerHTML = translateRawWord("Political:")! + " " + obj.political;
    viewNotes.innerHTML = (obj.notes || "");


    const allObjectsAtStart = document.querySelectorAll<HTMLElement>(".viewtag");

    allObjectsAtStart.forEach((object) => {
      object.style.display = "flex";
    });


    // Get all the elements with the class "viewtag" and store them in a variable called "allObjects"
    let allObjects = document.getElementsByClassName("viewtag");

    // Loop through all the objects in the array
    for (let i = 0; i < allObjects.length; i++) {

      // Store the current object's HTML in a variable called "item"
      let item = allObjects[i] as HTMLElement;
      // Get the text from the object's HTML and store it in a variable called "tempText"
      let tempText = item.innerHTML.substring(item.innerHTML.indexOf(':') + 1).trim();

      // Check if the text is empty, null, or undefined
      if (tempText.length <= 0 || tempText.replace(" ", "") == "" || tempText == null || tempText == undefined || tempText == "0") {
        // Remove the object from the page
        // allObjects[i].remove();

        item.style.display = "none";
        // i--;
      }
    }

    if (viewNotes.innerHTML.length <= 0) {
      document.getElementById("space-maker")!.style.display = "none";
    }

    // Hobbies

    const viewHobbyBase = document.querySelector(".v-hobby-base") as HTMLDivElement;
    const viewHobbySpacemaker = document.querySelector(".v-hobby-space-maker") as HTMLDivElement;

    viewHobbyBase.style.display = "block";

    if (Object.keys(obj.hobbies).length >= 1) {
      const hobbyContainer = document.querySelector(".v-hobby-base") as HTMLDivElement;

      for (const [_, hobby] of Object.entries(obj.hobbies)) {
        const hobbyVar = (hobby as { hobby: string })

        if (hobbyVar.hobby != "" && hobbyVar.hobby != null && hobbyVar.hobby != undefined) {
          viewHobbySpacemaker.style.display = "block";
          const container = document.createElement("div");
          container.className = "v-hobby-container";

          const subContainer = document.createElement("div");
          subContainer.className = "hobby-subcontainer";

          const hobby_input = document.createElement("input") as HTMLInputElement;
          hobby_input.className = "form-input v-hobby";
          hobby_input.id = "v-hobby";
          hobby_input.type = "hobby";
          hobby_input.required = true;
          hobby_input.value = hobbyVar.hobby;
          hobby_input.disabled = true;

          hobbyContainer.appendChild(container);
          container.appendChild(subContainer);
          subContainer.appendChild(hobby_input);
        };
      };
    } else {
      viewHobbySpacemaker.style.display = "none";
    }

    // IPs

    const viewIpBase = document.querySelector(".v-ip-base") as HTMLDivElement;
    const viewIpSpacemaker = document.querySelector(".v-ip-space-maker") as HTMLDivElement;

    viewIpBase.style.display = "block";

    if (Object.keys(obj.ips).length >= 1) {
      const ipContainer = document.querySelector(".v-ip-base") as HTMLDivElement;

      for (const [_, ip] of Object.entries(obj.ips)) {
        const ipVar = (ip as { ip: string })

        if (ipVar.ip != "" && ipVar.ip != null && ipVar.ip != undefined) {
          viewIpSpacemaker.style.display = "block";
          const container = document.createElement("div");
          container.className = "v-ip-container";

          const subContainer = document.createElement("div");
          subContainer.className = "ip-subcontainer";

          const ip_input = document.createElement("input");
          ip_input.className = "form-input v-ip";
          ip_input.id = "v-ip";
          ip_input.type = "ip";
          ip_input.spellcheck = false;
          ip_input.value = ipVar.ip;
          ip_input.disabled = true;

          ipContainer.appendChild(container);
          container.appendChild(subContainer);
          subContainer.appendChild(ip_input);
        }
      };
    } else {
      viewIpSpacemaker.style.display = "none";
    }

    // Clubs

    const viewClubBase = document.querySelector(".v-club-base") as HTMLDivElement;
    const viewClubSpacemaker = document.querySelector(".v-club-space-maker") as HTMLDivElement;

    viewClubBase.style.display = "block";

    if (Object.keys(obj.clubs).length >= 1) {
      const clubContainer = document.querySelector(".v-club-base") as HTMLDivElement;

      for (const [_, club] of Object.entries(obj.clubs)) {
        const clubVar = (club as { club: string })

        if (clubVar.club != "" && clubVar.club != null && clubVar.club != undefined) {
          viewClubSpacemaker.style.display = "block";
          const container = document.createElement("div");
          container.className = "v-club-container";

          const subContainer = document.createElement("div");
          subContainer.className = "club-subcontainer";

          const club_input = document.createElement("input");
          club_input.className = "form-input v-club";
          club_input.id = "v-club";
          club_input.type = "club";
          club_input.spellcheck = false;
          club_input.value = clubVar.club;
          club_input.disabled = true;

          clubContainer.appendChild(container);
          container.appendChild(subContainer);
          subContainer.appendChild(club_input);
        }
      };
    } else {
      viewClubSpacemaker.style.display = "none";
    }

    // Sources

    const viewSourceBase = document.querySelector(".v-source-base") as HTMLDivElement;
    const viewSourceSpacemaker = document.querySelector(".v-source-space-maker") as HTMLDivElement;

    viewSourceBase.style.display = "block";

    if (Object.keys(obj.sources).length >= 1) {
      const sourceContainer = document.querySelector(".v-source-base") as HTMLDivElement;

      for (const [_, url] of Object.entries(obj.sources)) {
        const sourceVar = (url as { url: string })

        if (sourceVar.url != "" && sourceVar.url != null && sourceVar.url != undefined) {
          viewSourceSpacemaker.style.display = "block";
          const container = document.createElement("div");
          container.className = "v-source-container";

          const subContainer = document.createElement("div");
          subContainer.className = "source-subcontainer";

          const source_field = document.createElement("a");
          source_field.className = "source-link v-source";
          source_field.id = "v-source";
          source_field.type = "source";
          source_field.spellcheck = false;
          source_field.innerHTML = sourceVar.url;
          source_field.href = sourceVar.url;
          source_field.target = "_blank";

          sourceContainer.appendChild(container);
          container.appendChild(subContainer);
          subContainer.appendChild(source_field);
        }
      };
    } else {
      viewSourceSpacemaker.style.display = "none";
    }

    // Phone


    const viewPhoneBase = document.querySelector(".v-phone-base") as HTMLDivElement;
    const viewPhoneSpacemaker = document.querySelector(".v-phone-space-maker") as HTMLDivElement;

    viewPhoneBase.style.display = "block";

    if (Object.keys(obj.phone).length >= 1) {
      const phoneContainer = document.querySelector(".v-phone-base") as HTMLDivElement;

      for (const [_, phone] of Object.entries(obj.phone)) {
        const phoneVar = (phone as { number: string, valid: boolean, phoneinfoga: { Country: string } })

        if (phoneVar.number != "" && phoneVar.number != null && phoneVar.number != undefined) {
          viewPhoneSpacemaker.style.display = "block";
          const container = document.createElement("div");
          container.className = "v-phone-container";

          const subContainer = document.createElement("div");
          subContainer.className = "phone-subcontainer";

          const phone_input = document.createElement("input");
          phone_input.className = "form-input v-phone";
          phone_input.id = "v-phone";
          phone_input.type = "phone";
          phone_input.value = phoneVar.number;
          phone_input.disabled = true;


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
              // use a function and parse the arg valid/invalid and use the literal string everywhere
              if (phoneVar.valid == true) {
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
              } else if (phoneVar.valid == false) {
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
              if (phoneVar.phoneinfoga.Country != "" && phoneVar.phoneinfoga.Country != null && phoneVar.phoneinfoga.Country != undefined) {
                const abbrContainer = document.createElement("abbr")
                abbrContainer.className = "phone-info-abbr";
                abbrContainer.title = phoneVar.phoneinfoga.Country;

                const iconDiv = document.createElement("div");
                iconDiv.className = "service-icon-div";

                const icon = document.createElement("img");
                icon.className = "country-icon";
                icon.src = "./images/flags/" + phoneVar.phoneinfoga.Country + ".png";

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
      viewPhoneSpacemaker.style.display = "none";
    }


    // Email

    const viewEmailBase = document.querySelector(".v-email-base") as HTMLDivElement;
    const viewEmailSpacemaker = document.querySelector(".v-email-space-maker") as HTMLDivElement;

    viewEmailBase.style.display = "block";

    if (Object.keys(obj.email).length >= 1) {
      for (const [_, email] of Object.entries(obj.email)) {
        const emailVar = (email as person.Email)

        if (emailVar.mail != "" && emailVar.mail != null && emailVar.mail != undefined) {
          viewEmailSpacemaker.style.display = "block";
          const container = document.createElement("div");
          container.className = "v-email-container";

          const subContainer = document.createElement("div");
          subContainer.className = "email-subcontainer";

          const email_input = document.createElement("input");
          email_input.className = "form-input v-mail";
          email_input.id = "v-e-mail";
          email_input.type = "email";
          email_input.required = true;
          email_input.value = emailVar.mail;
          email_input.disabled = true;


          const infoBtn = document.createElement("div");
          infoBtn.className = "v-info-btn";

          const icon = document.createElement("ion-icon");
          icon.setAttribute("name", "information-outline");

          container.appendChild(subContainer);
          infoBtn.appendChild(icon);
          subContainer.appendChild(email_input);
          subContainer.appendChild(infoBtn);
          viewEmailBase.appendChild(container);

          infoBtn.onclick = function () {
            const infoDiv = container.querySelector(".v-info-div");

            if (!infoDiv) {
              const infoDiv = document.createElement("div");
              infoDiv.className = "v-info-div";

              container.appendChild(infoDiv);

              if (emailVar.valid == true) {
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
              } else if (emailVar.valid == false) {
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

              if (emailVar.services != undefined && emailVar.services != null) {
                for (const [_, service] of Object.entries(emailVar.services)) {
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
      viewEmailSpacemaker.style.display = "none";
    }


    // Accounts

    if (Object.keys(obj.accounts).length != 0 && obj.accounts != null) {
      for (const [_, accObj] of Object.entries(obj.accounts)) {
        const accVar = (accObj as { service: string, id: string, username: string, url: string, profilePicture: { [key: number]: { img: string, img_hash: number } }, bio: { [key: number]: { bio: string } } });

        //let accObj = obj.accounts[i];

        // Creating elements

        const base_div = document.createElement("div"); // Outer div
        base_div.className = "acc-chip";

        const pfp_img = document.createElement("img"); // Pfp img
        pfp_img.className = "userPfp";

        if (accVar.profilePicture != null) {
          pfp_img.src = "data:image/png;base64," + accVar.profilePicture[1]!.img;
        } else {
          pfp_img.src = "https://as2.ftcdn.net/v2/jpg/03/32/59/65/1000_F_332596535_lAdLhf6KzbW6PWXBWeIFTovTii1drkbT.jpg"
        }

        const info_div = document.createElement("div"); // Info div
        info_div.className = "info-container";

        const service_p = document.createElement("a");
        service_p.className = "serviceName";
        service_p.innerHTML = accVar.service;
        service_p.href = accVar.url;
        service_p.target = "_blank";

        const name_p = document.createElement("a");
        name_p.className = "userName";
        name_p.innerHTML = accVar.username;
        name_p.href = accVar.url;
        name_p.target = "_blank";


        document.querySelector(".accounts")!.appendChild(base_div);
        base_div.appendChild(pfp_img);
        base_div.appendChild(info_div);
        info_div.appendChild(service_p);
        info_div.appendChild(name_p);

        if (accVar.bio != null) {
          const bio_p = document.createElement("p");
          bio_p.className = "userBio";
          bio_p.innerHTML = accVar.bio[1].bio;

          info_div.appendChild(bio_p);
        }
      }
    }
  }
  // edit button
  e_icon_div.onclick = function () {
    let obj2 = new Person(obj);
    obj2.Edit()
  }
}


const editSaveBtn = document.querySelector("#e-savebtn")! as HTMLDivElement;

editSaveBtn.onclick = async function () {
  const editShowID = document.querySelector("#e-showid")! as HTMLParagraphElement;

  const editNameTag = document.querySelector(".e-name-tag") as HTMLInputElement;
  const editAge = document.querySelector(".e-age") as HTMLParagraphElement;
  const editBday = document.querySelector(".e-bday") as HTMLParagraphElement;
  const editAddress = document.querySelector(".e-address") as HTMLParagraphElement;
  const editPhoneContainers = document.querySelectorAll(".phone-container") as NodeListOf<HTMLDivElement>;
  const editKids = document.querySelector(".e-kids")! as HTMLParagraphElement;
  const editHobbyContainers = document.querySelectorAll(".hobby-container") as NodeListOf<HTMLDivElement>;
  const editOccupation = document.querySelector(".e-occupation") as HTMLParagraphElement;
  const editPrevOccupation = document.querySelector(".e-prevoccupation") as HTMLParagraphElement;
  const editEducation = document.querySelector(".e-education") as HTMLParagraphElement;
  const editPets = document.querySelector(".e-pets") as HTMLParagraphElement;
  const editClubContainers = document.querySelectorAll(".club-container") as NodeListOf<HTMLDivElement>;
  const editLegal = document.querySelector(".e-legal") as HTMLParagraphElement;
  const editPolitical = document.querySelector(".e-political") as HTMLParagraphElement;
  const editSourceContainers = document.querySelectorAll(".source-container") as NodeListOf<HTMLDivElement>;
  const editNotes = document.querySelector(".e-notes") as HTMLDivElement;
  const editEmailContainers = document.querySelectorAll(".email-container") as NodeListOf<HTMLDivElement>;
  const editIPContainers = document.querySelectorAll(".ip-container") as NodeListOf<HTMLDivElement>;

  let id = editShowID.innerHTML;

  let name = editNameTag.value;

  let gender = checkDropdownValue("edit", "gender");

  let ethnicity = checkDropdownValue("edit", "ethnicity");

  let age = parseInt(editAge.innerHTML);

  if (age < 0) {
    age *= -1;
  }
  if (age > 128) {
    age = 128;
  }

  let bday = editBday.innerHTML;
  let address = editAddress.innerHTML;

  let phoneNumbers: { [key: string]: { number: string } } = {};

  editPhoneContainers.forEach((container: HTMLDivElement) => {
    const phoneInput: HTMLInputElement | null = container.querySelector('input[type="tel"]')!;

    const phoneNumber: string = phoneInput.value.toString();

    phoneNumbers[phoneNumber] = {
      "number": phoneNumber
    };
  });

  let civilstatus = checkDropdownValue("edit", "civilstatus");

  let kids = editKids.innerHTML;

  let hobbies: { [key: string]: { hobby: string } } = {};

  editHobbyContainers.forEach(function (container) {
    let hobbyInput = container.querySelector("input")!;
    hobbies[hobbyInput.value] = {
      "hobby": hobbyInput.value
    };
  });

  let occupation = editOccupation.innerHTML;
  let prevoccupation = editPrevOccupation.innerHTML;
  let education = editEducation.innerHTML;

  let religion = checkDropdownValue("edit", "religion");

  let pets = editPets.innerHTML;

  let clubs: { [key: string]: { club: string } } = {};

  editClubContainers.forEach(function (container) {
    let clubInput = container.querySelector("input")!;
    clubs[clubInput.value] = {
      "club": clubInput.value
    };
  });

  let legal = editLegal.innerHTML;
  let political = editPolitical.innerHTML;

  let sources: { [key: string]: { url: string } } = {};

  editSourceContainers.forEach(function (container) {
    let sourceInput = container.querySelector("input")!;
    sources[sourceInput.value] = {
      "url": sourceInput.value
    };
  });

  let notes = editNotes.innerHTML;

  let emailAddresses: { [key: string]: { mail: string, src: string, services: string } } = {};

  editEmailContainers.forEach(function (container) {
    let hiddenElement = container.querySelector(".hidden-email-save")!;

    // FIXME this is beautiful
    let hiddenElementVal = null;

    if (hiddenElement.innerHTML != "" && hiddenElement.innerHTML != null && hiddenElement.innerHTML != undefined) {
      hiddenElementVal = JSON.parse(hiddenElement.innerHTML);
    }

    let emailInput = container.querySelector("input")!;
    emailAddresses[emailInput.value] = {
      "mail": emailInput.value,
      "src": "manual",
      "services": hiddenElementVal
    };
  });

  let ips: { [key: string]: { ip: string } } = {};

  editIPContainers.forEach(function (container) {
    let ipInput = container.querySelector("input")!;
    ips[ipInput.value] = {
      "ip": ipInput.value
    };
  });

  const loadingSpinner = document.querySelector("#e-loading-spinner")! as HTMLDivElement;
  loadingSpinner.style.display = "flex"

  const res = await fetch(apiCall("/people/" + id))

  let data = await res.json() as Person;
  let obj = new Person();

  if (data == null) {
    obj.accounts = {};
  } else {
    obj.accounts = data.accounts;
  }

  obj.id = id;
  obj.name = name;
  obj.gender = gender || '';
  obj.ethnicity = ethnicity || '';
  obj.age = age;
  obj.bday = bday;
  obj.address = address;
  obj.phone = phoneNumbers as { [key: string]: person.PhoneNumber };
  obj.civilstatus = civilstatus || '';
  obj.kids = kids;
  obj.hobbies = hobbies;
  obj.email = emailAddresses as unknown as { [key: string]: person.Email };
  obj.ips = ips;
  obj.occupation = occupation;
  obj.prevoccupation = prevoccupation;
  obj.education = education;
  obj.religion = religion || '';
  obj.pets = pets;
  obj.clubs = clubs;
  obj.legal = legal;
  obj.political = political;
  obj.sources = sources;
  obj.notes = notes;

  obj.Post(loadingSpinner);

}

document.getElementById("backbtn")!.onclick = function () {
  mainContainer.style.display = "flex";
  container.style.display = "none";

  document.getElementById("space-maker")!.style.display = "block";

  var elements = document.getElementsByClassName("acc-chip");

  while (elements.length > 0) {
    elements[0].parentNode!.removeChild(elements[0]);
  }

  var elements = document.getElementsByClassName("v-phone-container");

  while (elements.length > 0) {
    elements[0].parentNode!.removeChild(elements[0]);
  }

  var elements = document.getElementsByClassName("v-email-container");

  while (elements.length > 0) {
    elements[0].parentNode!.removeChild(elements[0]);
  }
}

// Back Button

document.getElementById("e-backbtn")!.onclick = function () {
  mainContainer.style.display = "flex";
  editContainer.style.display = "none";

  var phoneElements = document.getElementsByClassName("phone-container");

  while (phoneElements.length > 0) {
    phoneElements[0].parentNode!.removeChild(phoneElements[0]);
  }

  var mailElements = document.getElementsByClassName("email-container");

  while (mailElements.length > 0) {
    mailElements[0].parentNode!.removeChild(mailElements[0]);
  }

  var hobbyElements = document.getElementsByClassName("hobby-container");

  while (hobbyElements.length > 0) {
    hobbyElements[0].parentNode!.removeChild(hobbyElements[0]);
  }

  var clubElements = document.getElementsByClassName("club-container");

  while (clubElements.length > 0) {
    clubElements[0].parentNode!.removeChild(clubElements[0]);
  }

  var ipElements = document.getElementsByClassName("ip-container");

  while (ipElements.length > 0) {
    ipElements[0].parentNode!.removeChild(ipElements[0]);
  }

  var sourceElements = document.getElementsByClassName("source-container");

  while (sourceElements.length > 0) {
    sourceElements[0].parentNode!.removeChild(sourceElements[0]);
  }

  const parentElement = document.querySelector(".e-accounts") as HTMLDivElement;
  parentElement.innerHTML = "";
}

document.getElementById("acc-backbtn")!.onclick = function () { // account back button
  location.reload();
}

runOnStart();


async function runOnStart() {
  const noResults = document.getElementById("base-no-results") as HTMLDivElement;

  const exportBtn = document.getElementById("exportbtn") as HTMLDivElement;

  let x = document.querySelector("#list-holder") as HTMLDivElement;

  const data = await getData() as object[];

  if (Object.entries(data).length == 0) {
    exportBtn.style.display = "none";
  } else {
    for (const [i, _] of Object.entries(data)) {
      //let obj = data[Number(i)] as any;
      let obj = data[Number(i)] as Person;

      createCards(obj);
    }
  }

  if (x.childElementCount <= 0) {
    noResults.style.display = "flex";

    x.style.display = "none";
  } else {
    noResults.style.display = "none";

    x.style.display = "block";
  }
}

async function searchEntries() {
  const inputElement = document.getElementById("searchbar") as HTMLInputElement;
  let input = inputElement.value.toLowerCase();

  const noResults = document.getElementById("base-no-results") as HTMLDivElement;
  let x = document.querySelector('#list-holder') as HTMLDivElement;

  let amountOfResults = x.childElementCount;

  for (let i = 0; i < x.childElementCount; i++) {
    let card = x.children[i];

    let name = card.querySelector(".hitbox-abbr > .hitbox > div.text-container > p.card-text")!;

    if (!name.innerHTML.toLowerCase().includes(input)) {
      card.classList.add("dont-show");

      amountOfResults--;
    } else {
      card.classList.remove("dont-show");

      amountOfResults++;
    }
  }

  if (amountOfResults <= 0) {
    noResults.style.display = "flex";

    x.style.display = "none";
  } else {
    noResults.style.display = "none";

    x.style.display = "flex";
  }
}

export { };
