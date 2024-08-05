const channel = new BroadcastChannel("seekr-channel");

channel.addEventListener('message', (event) => {
  if (event.data.type === "theme") {
    const theme = event.data.theme;

    document.documentElement.setAttribute("data-theme", theme);
  }
});


const myBtn = document.querySelector(".create-mail-btn");
const myPara = document.querySelector(".email-headline");

if (myBtn && myPara) {
  myBtn.addEventListener("click", () => {
    fetch("https://www.developermail.com/api/v1/mailbox", {
      method: "PUT",
      headers: {
        "accept": "application/json"
      },
      body: ""
    })
      .then(response => response.json()) // Parse the response as JSON
      .then(data => {
        if (data.success && data.errors === null) {
          myPara.textContent = `Name: ${data.result.name}, Token: ${data.result.token}`;
        } else {
          myPara.textContent = "An error occurred.";
        }
      })
      .catch(error => {
        myPara.textContent = "An error occurred.";
        console.error(error);
      });
  });
}