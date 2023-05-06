const channel = new BroadcastChannel("theme-channel");

// Listen for messages on the broadcast channel
channel.addEventListener('message', (event) => {
  if (event.data.type === "theme") {
    const theme = event.data.theme;
    localStorage.setItem("theme", theme);

    document.documentElement.setAttribute("data-theme", theme);
  }
});


const myBtn = document.querySelector(".create-mail-btn");
const myPara = document.querySelector(".email-headline");

if (myBtn && myPara) {
  myBtn.addEventListener("click", () => {
    fetch("https://cors-anywhere.herokuapp.com/https://www.developermail.com/api/v1/mailbox", {
      method: "PUT",
      headers: {
        "accept": "application/json"
      },
      body: ""
    })
      .then(response => response.json()) // Parse the response as JSON
      .then(data => {
        // Check if the API call was successful and there are no errors
        if (data.success && data.errors === null) {
          // Set the text of the paragraph tag to the name and token values
          myPara.textContent = `Name: ${data.result.name}, Token: ${data.result.token}`;
        } else {
          // Set the text of the paragraph tag to indicate an error occurred
          myPara.textContent = "An error occurred.";
        }
      })
      .catch(error => {
        // Set the text of the paragraph tag to indicate an error occurred
        myPara.textContent = "An error occurred.";
        console.error(error);
      });
  });
}