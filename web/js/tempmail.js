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

const testData = `Received: by mail-ed1-f52.google.com with SMTP id 4fb4d7f45d1cf-505934ccc35so4855281a12.2
for <z-tpwpj5@developermail.com>; Sat, 22 Apr 2023 11:25:56 -0700 (PDT)
DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;
d=gmail.com; s=20221208; t=1682187955; x=1684779955;
h=to:subject:message-id:date:from:mime-version:from:to:cc:subject
 :date:message-id:reply-to;
bh=lum9DbkH3N1hbHsgEq0AlY9Wwvjg7O89ppGKncH3YzU=;
b=VIkhYF2Nr6vXKnB5Qmh3F5c1q4HS5Eky65vXOBqWBEBcdrwZ20rV5s4nW+KpSEY3iA
 SahhN3+kqmNtCw8exeS5NYq2vU63HwpR0iFkizVD2LiiocngUOza0GuBdbLZSTk05Twj
 P2X4i58h9UGDW1DPhGgM1fi4FaFC3fhrT2i9lqJqVow9PVcc3D/WxE5XNkruiZFnqBxX
 CdJPMOm6szAUasdZnWc72gc4YtW1/x053VDfCZOyt37WJHOkKKpjT6dsuhrX0LMKaZUv
 jwzRzRbngZPUCdRBiFAAUACal3+iGtpYgFJ8zmIdZjRrXpf+MDjJw5S51hzXZvTniI2H
 WZLQ==
X-Google-DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;
d=1e100.net; s=20221208; t=1682187955; x=1684779955;
h=to:subject:message-id:date:from:mime-version:x-gm-message-state
 :from:to:cc:subject:date:message-id:reply-to;
bh=lum9DbkH3N1hbHsgEq0AlY9Wwvjg7O89ppGKncH3YzU=;
b=k2KaUIF2Mqg5DWQd2JtjsrM7s4WDXCAWhAB9XrSei/M84jVlRNAti/YTocJfxR/7Ef
 aCD56/prn40sMIuDJr/5lpHA+Rvs0tQmw2MXNSv8wS4MSJjW1tKBwACrTIKGquAlLkKq
 ulzMvMSFCVJ7pwnGuCdTINxS/lg4HB+CHQeK1S2Adq3cP38m4gZ4gO5OiGPOJUPFMyAY
 QdclJJTqhkA1PeiJltlGPUsvUSPI8Ej+GFrfxJvcyCLH7dp3/WPU75uz6B9iPHC3hxTd
 B7JWvbEqEDsToYz1z21riOLytN1zJspWbb4pBLJ329iMJ2tS+cqaNF/VmWkVskmb8Knw
 wncg==
X-Gm-Message-State: AAQBX9eNMYf4aE9hcokkjxO0vODcrA8Bbto7nPJydt6vtiX1PFVGz3vS
9RZcz6uYpf5Z7qh1+FtlHL7fBUFujaxfDVYrCRBlETltdPGug+5C
X-Google-Smtp-Source: AKy350YNtfiDyAjHCur4oE25fF/xOt7AVfoh7lXY8nYP/QG3IhYNHeyDSy52PsnjmQmjtDY/6TzQlNYaRBYFohA9IOA=
X-Received: by 2002:aa7:d281:0:b0:506:b88a:cab4 with SMTP id
w1-20020aa7d281000000b00506b88acab4mr7984149edq.3.1682187954715; Sat, 22 Apr
2023 11:25:54 -0700 (PDT)
MIME-Version: 1.0
From: Tom Spitz <tomspitz04@gmail.com>
Date: Sat, 22 Apr 2023 18:25:43 +0000
Message-ID: <CACrofBWnXXaW6-Zf1z3dDrTfREwbWTykmu6V0gmiPzxAKOFh8A@mail.gmail.com>
Subject: js suck
To: z-tpwpj5@developermail.com
Content-Type: multipart/alternative; boundary="0000000000007c7e1505f9f0e7ee"

--0000000000007c7e1505f9f0e7ee
Content-Type: text/plain; charset="UTF-8"

js
suck
suck
suck

--0000000000007c7e1505f9f0e7ee
Content-Type: text/html; charset="UTF-8"

<div dir="ltr"><div>js <br></div><div>suck</div><div>suck</div><div>suck<br></div></div>

--0000000000007c7e1505f9f0e7ee--`

function extractData(data) {
  const pattern = /Content-Type: multipart\/alternative; boundary="([^"]*)"/;
  const match = data.match(pattern);
  if (!match) {
    throw new Error("Invalid data: missing boundary");
  }
  const boundary = match[1];
  const split = data.split(`--${boundary}`);
  if (split.length < 4) {
    throw new Error("Invalid data: not enough parts");
  }
  return [ split[1].replace(/Content-Type:[^<]+/, ''), split[2].replace(/Content-Type:[^<]+/, '') ];
}

console.log(extractData(testData));

console.log(extractData(testData)[0]);
console.log(extractData(testData)[1]);
console.log(extractData(testData)[2]);


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
                      console.log(data.result[i].value);

                      mailDetailsDate.innerHTML = data.result[i].value.split("Date: ")[1].split("\r\nMessage-ID")[0].match(/\d{2}:\d{2}/)[0];
                      mailDetailsSender.innerHTML = data.result[i].value.split("From: ")[1].split("\r\nDate:")[0].match(/\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b/)[0];
                      mailDetailsHeader.innerHTML = data.result[i].value.split("Subject: ")[1].split("\r\nTo: ")[0];
                      let actualData = data.result[i].value
                      console.log(extractData(actualData)[0]);
                      console.log(extractData(actualData)[1]);

                      iframe = document.getElementById('content-mail-body-iframe');
                      iframe.srcdoc = extractData(actualData)[1]
                      //console.log(mailDetailsBody.innerHTML);
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

