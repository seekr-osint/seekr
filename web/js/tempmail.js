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

const testData = `Received: by mail-wm1-f53.google.com with SMTP id 5b1f17b1804b1-3f1950f5676so13982675e9.3
for <z-rptjp7@developermail.com>; Sat, 22 Apr 2023 03:37:23 -0700 (PDT)
DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;
d=gmail.com; s=20221208; t=1682159842; x=1684751842;
h=to:subject:message-id📅from:mime-version:from:to:cc:subject
 📅message-id:reply-to;
bh=z42V4dDYdrFpBK0is0aEjnMzMoSgnXIKTfsQHNl3GLA=;
b=Qc/cmtIboGSWl7TTf3qEGDwob9//E+5CBTTGQj+H5jGqIU8DBh3KBy7e2/P4HpRs5H
 qwVyxdQe82dcJG7Udwoxwm/bG7Ss5P9yKn9HsDMyF+bHaATUjW1D4342DFGEer3WFxuX
 zqFkcoiVq2sJJnaqGhTrw7SDtd2/Jgokwrn79J72wQf8UOYG31Ln/8BMy9+Ypr1BANZt
 dnnfwp6Z+YlUZWie1bgifXBdLM020A6DVzbojFuMKEVavsF1jcFD8dzzVX5+/UqD73SF
 xgVwHCWiN5ebjLLBc+VEq9N0mlZMcfCDGwaw0VLl7tmRY6DLi5BPhgZLLKCNNDuI1+XE
 9xOw==
X-Google-DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed;
d=1e100.net; s=20221208; t=1682159842; x=1684751842;
h=to:subject:message-id📅from:mime-version:x-gm-message-state
 :from:to:cc:subject📅message-id:reply-to;
bh=z42V4dDYdrFpBK0is0aEjnMzMoSgnXIKTfsQHNl3GLA=;
b=fc+y4uilysz+OMsjafjYH5wzWGEbmOh7xeyy1IrQSUSs+V+iJ2S7DWW1O+5IO3sDNn
 75qxVhITPFCp4MDZse6LeHSPr7WmYErk/cBiQih7mFXMVMGW1h98M3wRearf3wtqY/yx
 NVbdzmQBLeVJrvfeOY7Lh1j0+uIISjBPEtaLtwhzbfi+BVD3OToub/6NixIwj414v6rQ
 WxujOolaLakcrKARM+w2KNkPNSRJi3a1UJdZYoYGK33hPZ9teZWW5sYmVX23IDEyVeD0
 CdbCsLE5RIGqYlbaOZYi7VjUYBrbi2qU+onq8wZwKXe122NPW8yQ0wOcThT60tt8pF22
 Uw0g==
X-Gm-Message-State: AAQBX9d8vkgJm9TS8fnaOROsOz2uqE3kpWCRgvDIogtoSpBnUz4UOWSK
Q13V2BQJw5NaFNwOWnEkPsKARXJ3hx0nQh0WHtSd9rAY
X-Google-Smtp-Source: AKy350ZWaVCc1n1xSErkCq9qKudQJSsBafibxC5fvjUWkaMEis2GKuYOdqoUATgLpceWXRmOiFtcKmkJpSpH840Vap0=
X-Received: by 2002:adf:dc4e:0:b0:2f2:7a49:c6cc with SMTP id
m14-20020adfdc4e000000b002f27a49c6ccmr5547370wrj.70.1682159842209; Sat, 22
Apr 2023 03:37:22 -0700 (PDT)
MIME-Version: 1.0
From: Some guy <someemail@gmail.com>
Date: Sat, 22 Apr 2023 12:37:10 +0200
Message-ID: <CAPYZEoFs7AcRR7Z7qNy+C+Vb3=_3Nx95=_me_kAsnuDSNzjErA@mail.gmail.com>
Subject: tes
To: z-rptjp7@developermail.com
Content-Type: multipart/alternative; boundary="000000000000d9b24c05f9ea5bff"

--000000000000d9b24c05f9ea5bff
Content-Type: text/plain; charset="UTF-8"

pls
work

--000000000000d9b24c05f9ea5bff
Content-Type: text/html; charset="UTF-8"
Content-Transfer-Encoding: quoted-printable

<div dir=3D"ltr">pls=C2=A0<div>work</div></div>

--000000000000d9b24c05f9ea5bff--`

mailDetailsBody.innerHTML = testData.split(`--([0-9]+([A-Za-z]+[0-9]+)+)`)[0];


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
                      mailDetailsBody.innerHTML = data.result[i].value.split(`--([0-9]+([A-Za-z]+[0-9]+)+)`)[0]; //.split(`\r\n\r\n--\w{28}--`)[0]
                      console.log(mailDetailsBody.innerHTML);
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