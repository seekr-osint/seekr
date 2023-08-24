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

const mailDetailsFrame = document.getElementById("content")
const mailDetailsDate = document.getElementById("content-mail-date")
const mailDetailsSender = document.getElementById("content-mail-sender")
const mailDetailsHeader = document.getElementById("content-mail-header")
const mailDetailsBody = document.getElementById("content-mail-body")

const mailListFrame = document.getElementById("emails")
const newBtn = document.getElementById("create-mail-btn");
const delBtn = document.getElementById("del-mail-btn");
const refreshBtn = document.getElementById("refresh-btn");
const currentMail = document.getElementById("current-mail");
const mailToken = document.getElementById("mail-token");
const loadingSpinner = document.getElementById("mail-loading-spinner");
const refreshLoadingSpinner = document.getElementById("refresh-loading-spinner");


newBtn.addEventListener("click", () => {
  if (mailToken.innerHTML == "") {
    loadingSpinner.style.display = "flex";

    fetch("http://localhost:8569/developermail/api/mailbox", {
      method: "PUT",
      headers: {
        "accept": "application/json"
      },
      body: ""
    })
      .then(response => response.json())
      .then(data => {
        if (data.success && data.errors === null) {
          currentMail.value = data.result.name + "@developermail.com";

          mailToken.innerHTML = data.result.token;
        } else {
          currentMail.value = "An error occurred.";
        }
      })
      .then(() => {
        loadingSpinner.style.display = "none";
      })
      .catch(error => {
        currentMail.value = "An error occurred.";
        console.error(error);
      });
  }
});

delBtn.addEventListener("click", () => {
  if (mailToken.innerHTML != "") {
    loadingSpinner.style.display = "flex";

    fetch("http://localhost:8569/developermail/api/mailbox/" + currentMail.value.split("@")[0], {
      method: "DELETE",
      headers: {
        "accept": "application/json",
        "X-MailboxToken": mailToken.innerHTML
      },
      body: ""
    })
      .then(response => response.json())
      .then(data => {
        if (data.success && data.errors === null) {
          mailToken.innerHTML = "";

          currentMail.value = "Tempmail deleted!";

          setTimeout(() => {
            currentMail.value = "";
          }, 3000);
        } else {
          currentMail.value = "An error occurred.";
        }
      })
      .then(() => {
        loadingSpinner.style.display = "none";
      })
      .catch(error => {
        currentMail.value = "An error occurred.";
        console.error(error);
      });
  }
});

refreshBtn.addEventListener("click", () => {
  if (mailToken.innerHTML != "") {
    refreshLoadingSpinner.style.display = "flex";

    // Get message IDs

    fetch("http://localhost:8569/developermail/api/mailbox/" + currentMail.value.split("@")[0], {
      method: "GET",
      headers: {
        "accept": "application/json",
        "X-MailboxToken": mailToken.innerHTML
      }
    })
      .then(response => response.json())
      .then(data => {
        if (data.success && data.errors === null) {

          console.log(data.result);

          // Delete old mail chips
          while (mailListFrame.firstChild) {
            mailListFrame.removeChild(mailListFrame.firstChild);
          }

          // Get messages using IDs

          if (data.result.length > 0) {
            fetch("http://localhost:8569/developermail/api/mailbox/" + currentMail.value.split("@")[0] + "/messages", {
              method: "POST",
              headers: {
                "accept": "application/json",
                "X-MailboxToken": mailToken.innerHTML,
                "Content-Type": "application/json"
              },
              body: JSON.stringify(data.result)
            })
              .then(response => {
                if (response.ok) {
                  return response.json();
                } else {
                  throw new Error('Network response was not ok.');
                }
              })
              .then(data => {
                if (data.success && data.errors === null) {
                  for (let i = 0; i < data.result.length; i++) {
                    const mailChip = document.createElement("div");
                    mailChip.className = "mail-chip chip";

                    const mailDate = document.createElement("p");
                    mailDate.className = "mail-chip-detail mail-date";
                    mailDate.innerHTML = data.result[i].value.split("Date: ")[1].split("\r\nMessage-ID")[0].match(/\d{2}:\d{2}/)[0];

                    const mailSender = document.createElement("p");
                    mailSender.className = "mail-chip-detail mail-sender";
                    mailSender.innerHTML = data.result[i].value.split("From: ")[1].split("\r\nDate:")[0].match(/\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b/)[0];

                    const mailHeader = document.createElement("p");
                    mailHeader.className = "mail-chip-detail mail-header";
                    mailHeader.innerHTML = data.result[i].value.split("Subject: ")[1].split("From: ")[0];


                    mailListFrame.appendChild(mailChip);
                    mailChip.appendChild(mailDate);
                    mailChip.appendChild(mailSender);
                    mailChip.appendChild(mailHeader);

                    mailChip.addEventListener("click", () => {
                      console.log(data.result[i].value);

                      mailDetailsDate.innerHTML = data.result[i].value.split("Date: ")[1].split("\r\nMessage-ID")[0].match(/\d{2}:\d{2}/)[0];
                      mailDetailsSender.innerHTML = data.result[i].value.split("From: ")[1].split("\r\nDate:")[0].match(/\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b/)[0];
                      mailDetailsHeader.innerHTML = data.result[i].value.split("Subject: ")[1].split("From: ")[0];

                      console.log(data.result[i].value.split("Subject: ")[1].split("From: ")[0]);

                      let actualData = data.result[i].value

                      const myIframe = document.getElementById('content-mail-body-iframe');
                      myIframe.srcdoc = "<!DOCTYPE html>" + actualData.split("<html>")[1].split("</html>")[0];
                    });
                  }
                } else {
                  console.log("An error occurred.");
                }
              })
              .then(() => {
                refreshLoadingSpinner.style.display = "none";
              })
              .catch(error => {
                console.log("An error occurred.");
                console.error(error);
              });
          }
        } else {
          console.log("An error occurred.");
        }
      })
      .then(() => {
        refreshLoadingSpinner.style.display = "none";
      })
      .catch(error => {
        console.log("An error occurred.");
        console.error(error);
      });
  }
});

