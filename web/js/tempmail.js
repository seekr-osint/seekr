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
                    mailHeader.innerHTML = data.result[i].value.split("Subject: ")[1].split("\r\nTo: ")[0];


                    mailListFrame.appendChild(mailChip);
                    mailChip.appendChild(mailDate);
                    mailChip.appendChild(mailSender);
                    mailChip.appendChild(mailHeader);

                    mailChip.addEventListener("click", () => {
                      mailDetailsDate.innerHTML = data.result[i].value.split("Date: ")[1].split("\r\nMessage-ID")[0].match(/\d{2}:\d{2}/)[0];
                      mailDetailsSender.innerHTML = data.result[i].value.split("From: ")[1].split("\r\nDate:")[0].match(/\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b/)[0];
                      mailDetailsHeader.innerHTML = data.result[i].value.split("Subject: ")[1].split("\r\nTo: ")[0];
                      mailDetailsBody.innerHTML = data.result[i].value.split("Content-Type: text/plain; charset=\"UTF-8\"\r\nContent-Transfer-Encoding: quoted-printable\r\n\r\n")[1].split("\r\n\r\n--")[0];
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


// {
//   "key": "638176942571524819",
//     "value": "Received: by mail-wm1-f49.google.com with SMTP id 5b1f17b1804b1-3f17e584462so20288345e9.2\r\n        for <z-tsmyw5@developermail.com>; Fri, 21 Apr 2023 10:17:35 -0700 (PDT)\r\nDKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;\r\n        d=gmail.com; s=20221208; t=1682097454; x=1684689454;\r\n        h=to:subject:message-id:date:from:mime-version:from:to:cc:subject\r\n         :date:message-id:reply-to;\r\n        bh=MGtAwkOPhhOK6xhGXRsv5NlFTcVWDQz1Orv8i2rlbkE=;\r\n        b=QB/rjB2l4YQmRuhq3hl1Wd5GecPv3okJCqDxG0oI5LPa7dEqPcH8S6g2yEPZRCvQtT\r\n         0NOO5pFYIQ1dxDCd+1VWUrN2mSIxwyY0YJQR5AvY3hiRaYvPJqFmECiFjSSi6umAT0ho\r\n         1eRb/QjtuoBiKvTt5D7uR0SKMAubpTZFM5Aiih8rWVgvdzSZMb96Wn1CIC1dMeobn6Ye\r\n         bj17iX7TwyOfej5akLhT2K6cx64+/9WSP9DWhHuxk6Dvgb6V/TKm3qc6xPZiZJDv2VKQ\r\n         bs+SJdkjCe+vMvq1Ee6ZLXWRhd1QkjUPF+FvcA+mHpwbgU0cd6N0kb7u/lHf85I5Tan+\r\n         4lWA==\r\nX-Google-DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;\r\n        d=1e100.net; s=20221208; t=1682097454; x=1684689454;\r\n        h=to:subject:message-id:date:from:mime-version:x-gm-message-state\r\n         :from:to:cc:subject:date:message-id:reply-to;\r\n        bh=MGtAwkOPhhOK6xhGXRsv5NlFTcVWDQz1Orv8i2rlbkE=;\r\n        b=TjyrN2qPfyCTyC7J/x2BveiPKZE66Gg5bpqzlHN0FpkXTJRXjmBsfiEj6tUXRf0Dre\r\n         Pw71HMqyRytSqNcHX4rZ0yR7kgf959m5u4Iba2n/pV/QLdEMJtQWJc8LajGcFb1bF0f1\r\n         qsN9rWegP6GyO8fDU09LXgHFqr+/CxTCGuF5ZUQdy4r6clNeve+8BPMbtJcNQDQ+rT7a\r\n         BQLQlOhjtk8QJq0wTIS6ssack7teXHvGWTa4kUp6HCTVLHDZLiO/mqd+8JXy3/NpsTb+\r\n         1cdYaOWLPaZNw5u+H87/6f2hYL7HOFJy9CFshha+zEeXxETqempsgkyDSFl07YLyAIBl\r\n         h+CQ==\r\nX-Gm-Message-State: AAQBX9e5DIHZrK2Z5bcOYriiKNYIZrqwyhD98F2q2aivgAdQXeAw82XT\r\n\tFKLEM8VkPdk7aaV8aS+yJ9qFxWKlf1OrkO0OnKy9eDRA\r\nX-Google-Smtp-Source: AKy350ZCj9rTseJsPkP9ztIj6wDbyQzbpR/YWZfvLJHtYL50AsFFjM7Am2lYXtx6DJ2eJRybacDK9zGlYeD8eUfh9Ec=\r\nX-Received: by 2002:adf:fc87:0:b0:2d1:e517:4992 with SMTP id\r\n g7-20020adffc87000000b002d1e5174992mr4524402wrr.69.1682097453980; Fri, 21 Apr\r\n 2023 10:17:33 -0700 (PDT)\r\nMIME-Version: 1.0\r\nFrom: Forks are better than spoons <herbertgronemeier5@gmail.com>\r\nDate: Fri, 21 Apr 2023 19:17:22 +0200\r\nMessage-ID: <CAPYZEoETTStugAp0SH_qF2qTZwXcL1fVVWQ9Zi1UW-K1Cxz+Ew@mail.gmail.com>\r\nSubject: tets\r\nTo: z-tsmyw5@developermail.com\r\nContent-Type: multipart/alternative; boundary=\"00000000000038e10805f9dbd5a5\"\r\n\r\n--00000000000038e10805f9dbd5a5\r\nContent-Type: text/plain; charset=\"UTF-8\"\r\n\r\nte\r\n\r\n--00000000000038e10805f9dbd5a5\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n<div dir=\"ltr\">te</div>\r\n\r\n--00000000000038e10805f9dbd5a5--"
// } 