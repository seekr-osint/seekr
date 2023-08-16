const channel = new BroadcastChannel("seekr-channel");

// Listen for messages on the broadcast channel
channel.addEventListener('message', (event) => {
  if (event.data.type === "theme") {
    const theme = event.data.theme;

    document.documentElement.setAttribute("data-theme", theme);
  }
});


// The actual stuff

const countryDropdown = document.getElementById("country-select");
const selectedSelect = document.querySelector(".country-select select");

const checkboxName = document.getElementById("checkbox_01") as HTMLInputElement;
const checkboxNameIcon = document.getElementById("checkbox_01_icon") as HTMLElement;

const checkboxAddress = document.getElementById("checkbox_02") as HTMLInputElement;
const checkboxAddressIcon = document.getElementById("checkbox_02_icon") as HTMLElement;

const checkboxPhone = document.getElementById("checkbox_03") as HTMLInputElement;
const checkboxPhoneIcon = document.getElementById("checkbox_03_icon") as HTMLElement;

const checkboxVIN = document.getElementById("checkbox_04") as HTMLInputElement;
const checkboxVINIcon = document.getElementById("checkbox_04_icon") as HTMLElement;

const checkboxBusiness = document.getElementById("checkbox_05") as HTMLInputElement;
const checkboxBusinessIcon = document.getElementById("checkbox_05_icon") as HTMLElement;

const checkboxIP = document.getElementById("checkbox_06") as HTMLInputElement;
const checkboxIPIcon = document.getElementById("checkbox_06_icon") as HTMLElement;

const checkboxUsername = document.getElementById("checkbox_07") as HTMLInputElement;
const checkboxUsernameIcon = document.getElementById("checkbox_07_icon") as HTMLElement;

const checkboxDomain = document.getElementById("checkbox_08") as HTMLInputElement;
const checkboxDomainIcon = document.getElementById("checkbox_08_icon") as HTMLElement;

const list_elements = document.querySelectorAll(".link-list-holder li");


function resetAll() {
  for (let i = 0; i < list_elements.length; i++) {
    const element = list_elements[i] as HTMLElement;
    element.style.display = "flex";
  }
}

function checkChecboxValue(checkboxType: string): boolean {
  if (checkboxType == "name") {
    return checkboxName.checked;
  } else if (checkboxType == "address") {
    return checkboxAddress.checked;
  } else if (checkboxType == "phone") {
    return checkboxPhone.checked;
  } else if (checkboxType == "vin") {
    return checkboxVIN.checked;
  } else if (checkboxType == "business") {
    return checkboxBusiness.checked;
  } else if (checkboxType == "ip") {
    return checkboxIP.checked;
  } else if (checkboxType == "username") {
    return checkboxUsername.checked;
  } else if (checkboxType == "domain") {
    return checkboxDomain.checked;
  } else {
    return false;
  }
}

function checkCountry(): "all" | "ww" | "us" | "ca" | "uk" | "se" | "de" | undefined {
  if (document) {
    const selectedCountry = document.querySelector(".select-selected");

    if (selectedCountry) {
      const countries: { [key: string]: "all" | "ww" | "us" | "ca" | "uk" | "se" | "de" } = {};

      // English

      countries["Select country:"] = "all";
      countries["WorldWide"] = "ww";
      countries["United States"] = "us";
      countries["Canada"] = "ca";
      countries["United Kingdom"] = "uk";
      countries["Sweden"] = "se";
      countries["Germany"] = "de";

      return countries[selectedCountry.innerHTML]; // Error here
    }
  }
}

function listHandler() {
  type checkboxResultType = string | boolean;

  let listOfClasses: checkboxResultType[] = ["country", "name", "address", "phone", "vin", "business", "ip", "username", "domain"];

  resetAll();

  const selectedCountry = checkCountry();

  // Replace the first element with the country code
  if (selectedCountry != undefined) {
    listOfClasses[0] = selectedCountry;
  }

  if (checkChecboxValue("name") == false) {
    listOfClasses[1] = false;

    checkboxNameIcon.style.opacity = "0";
  } else {
    checkboxNameIcon.style.opacity = "1";
  }

  if (checkChecboxValue("address") == false) {
    listOfClasses[2] = false;    
    
    checkboxAddressIcon.style.opacity = "0";
  } else {
    checkboxAddressIcon.style.opacity = "1";
  }

  if (checkChecboxValue("phone") == false) {
    listOfClasses[3] = false;

    checkboxPhoneIcon.style.opacity = "0";
  } else {
    checkboxPhoneIcon.style.opacity = "1";
  }
  
  if (checkChecboxValue("vin") == false) {
    listOfClasses[4] = false;
    
    checkboxVINIcon.style.opacity = "0";
  } else {
    checkboxVINIcon.style.opacity = "1";
  }

  if (checkChecboxValue("business") == false) {
    listOfClasses[5] = false;

    checkboxBusinessIcon.style.opacity = "0";
  } else {
    checkboxBusinessIcon.style.opacity = "1";
  }

  if (checkChecboxValue("ip") == false) {
    listOfClasses[6] = false;

    checkboxIPIcon.style.opacity = "0";
  } else {
    checkboxIPIcon.style.opacity = "1";
  }

  if (checkChecboxValue("username") == false) {
    listOfClasses[7] = false;

    checkboxUsernameIcon.style.opacity = "0";
  } else {
    checkboxUsernameIcon.style.opacity = "1";
  }

  if (checkChecboxValue("domain") == false) {
    listOfClasses[8] = false;
    
    checkboxDomainIcon.style.opacity = "0";
  } else {
    checkboxDomainIcon.style.opacity = "1";
  }

  if (listOfClasses[1] == false && listOfClasses[2] == false && listOfClasses[3] == false && listOfClasses[4] == false && listOfClasses[5] == false && listOfClasses[6] == false && listOfClasses[7] == false && listOfClasses[8] == false) {
    for (let i = 0; i < list_elements.length; i++) {
      const element = list_elements[i] as HTMLElement;
  
      if (listOfClasses[0] == "all") {
        element.style.display = "flex";
      } else if (!element.classList.contains(listOfClasses[0].toString())) {
        element.style.display = "none";
      }
    }
  } else {
    for (let i = 0; i < list_elements.length; i++) {
      const element = list_elements[i] as HTMLElement;
      
      if (!element.classList.contains(listOfClasses[0].toString()) && listOfClasses[0] != "all") {
        element.style.display = "none";
      } else {
        let hasBeenChanged = false;

        for (let i = 1; i < listOfClasses.length; i++) {
          if (element.classList.contains(listOfClasses[i].toString())) {
            element.style.display = "flex";
            hasBeenChanged = true;
          }

          if (i == 8 && hasBeenChanged == false) {
            element.style.display = "none";
          }
        }
      }
    }
  }
}

function preListHandler() {
  const selectedCountry = document.querySelector(".select-selected");

  if (selectedCountry && selectedCountry.innerHTML != "") {
    listHandler();
  }
}

document.querySelector(".select-selected")!.addEventListener("DOMSubtreeModified", preListHandler);

checkboxName.addEventListener('change', preListHandler);

checkboxAddress.addEventListener('change', preListHandler);

checkboxPhone.addEventListener('change', preListHandler);

checkboxVIN.addEventListener('change', preListHandler);

checkboxBusiness.addEventListener('change', preListHandler);

checkboxIP.addEventListener('change', preListHandler);

checkboxUsername.addEventListener('change', preListHandler);

checkboxDomain.addEventListener('change', preListHandler);

export {};