function getElementId(containerType: string): string {
    const container = document.querySelector(`.${containerType}-container`);
    if (container) {
      return container.id;
    } else {
      return ""; // FIXME error
    }
  }
  
function checkGender(containerType: string): string {
    const selectedGender = getElementId(containerType);
    const gender: {[key: string]: string} = {};
  
    gender["Select gender:"] = "";
    gender["Male"] = "Male";
    gender["Female"] = "Female";
    gender["Other"] = "Other";
  
    return gender[selectedGender];
  }
function getGenderElementIndex(gender: string): string {
  const genderIndex: {[key: string]: string} = {};

  genderIndex[""] = "";
  genderIndex["Male"] = "0";
  genderIndex["Female"] = "1";
  genderIndex["Other"] = "2";

  return genderIndex[gender];
}

function checkCivilstatus(containerType: string): string {
  const selectedCivilstatus = getReligionElementIndex(containerType)
  const civilstatus: {[key: string]: string} = {};

  civilstatus["Select civil status:"] = "";
  civilstatus["Single"] = "Single";
  civilstatus["Married"] = "Married";
  civilstatus["Widowed"] = "Widowed";
  civilstatus["Divorced"] = "Divorced";
  civilstatus["Separated"] = "Separated";

  return civilstatus[selectedCivilstatus];
}

function getCivilstatusElementIndex(civilstatus: string): string {
  const civilstatusIndex: {[key: string]: string} = {};

  civilstatusIndex[""] = "";
  civilstatusIndex["Single"] = "0";
  civilstatusIndex["Married"] = "1";
  civilstatusIndex["Widowed"] = "2";
  civilstatusIndex["Divorced"] = "3";
  civilstatusIndex["Separated"] = "4";

  return civilstatusIndex[civilstatus];
}

function checkReligion(containerType: string): string {
  const selectedReligion = getElementId(containerType)
  const religion: {[key: string]: string} = {};

  religion["Select religion:"] = "";
  religion["Christianity"] = "Christianity";
  religion["Atheism"] = "Atheism";
  religion["Islam"] = "Islam";
  religion["Hinduism"] = "Hinduism";
  religion["Buddhism"] = "Buddhism";
  religion["Sikhism"] = "Sikhism";
  religion["Judaism"] = "Judaism";
  religion["Other"] = "Other";

  return religion[selectedReligion];
}

function getReligionElementIndex(religion: string): string {
  const religionIndex: {[key: string]: string} = {};

  religionIndex[""] = "";
  religionIndex["Christianity"] = "0";
  religionIndex["Atheism"] = "1";
  religionIndex["Islam"] = "2";
  religionIndex["Hinduism"] = "3";
  religionIndex["Buddhism"] = "4";
  religionIndex["Sikhism"] = "5";
  religionIndex["Judaism"] = "6";
  religionIndex["Other"] = "7";

  return religionIndex[religion];
}

