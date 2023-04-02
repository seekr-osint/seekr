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






const myBtn = document.querySelector(".create-mail-btn");
const myPara = document.querySelector(".email-headline");

myBtn.addEventListener("click", () => {
  fetch("https://www.developermail.com/api/v1/mailbox", {
    method: "PUT",
    headers: {
      "accept": "application/json"
    },
    mode: "no-cors", // Add this line to disable CORS
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