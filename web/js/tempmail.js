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

const testData1 = `Received: by mail-ed1-f52.google.com with SMTP id 4fb4d7f45d1cf-505934ccc35so4855281a12.2
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

// const testData2 = `
// Received: from github-highworker-4bd0f4f.ash1-iad.github.net (github-highworker-4bd0f4f.ash1-iad.github.net [10.56.101.39])
// 	by smtp.github.com (Postfix) with ESMTP id 84965600267
// 	for <z-tlfdj1@developermail.com>; Sun, 23 Apr 2023 04:23:41 -0700 (PDT)
// DKIM-Signature: v=1; a=rsa-sha256; c=relaxed/relaxed; d=github.com;
// 	s=pf2023; t=1682249021;
// 	bh=gV4eli2RcwC8Awt098BQtuQYbJLMi6j0s+4MYVSrMsQ=;
// 	h=Date:From:To:Subject:From;
// 	b=gRCOkHgxgI4+E55xLQ4RF8rV202u9jzWBU0w4LDIU4Mum1oMFk94Xvf5j+EOv514G
// 	 kUeMEGM2GY7EoWu4cKBGBaw1dXw6tH4UL97VpxBcIIsc4YKX9igQK07aXk0r05krHA
// 	 3yewxyKFgFFvgf9cD/8Tya/gt1fGhg9AtCe6q95o=
// Date: Sun, 23 Apr 2023 04:23:41 -0700
// From: GitHub <noreply@github.com>
// To: Hank3r-gh <z-tlfdj1@developermail.com>
// Message-ID: <6445153d76bc7_38bbc97c26a3@github-highworker-4bd0f4f.ash1-iad.github.net.mail>
// Subject: =?UTF-8?Q?=F0=9F=9A=80_Your_GitHub_launch_code?=
// Mime-Version: 1.0
// Content-Type: multipart/alternative;
//  boundary="--==_mimepart_6445153ca0e93_38bbc97c25a9";
//  charset=UTF-8
// Content-Transfer-Encoding: 7bit
// X-Auto-Response-Suppress: All


// ----==_mimepart_6445153ca0e93_38bbc97c25a9
// Content-Type: text/plain;
//  charset=UTF-8
// Content-Transfer-Encoding: quoted-printable

// Here's your GitHub launch code, @Hank3r-gh!

// Continue signing up for GitHub by entering the code below:

// 21337452

// You can enter it by visiting the link below:

// https://github.com/account_verifications?via_launch_code_email=3Dtrue

// You=E2=80=99re receiving this email because you recently created a new Gi=
// tHub account. If this wasn=E2=80=99t you, please ignore this email.

// Not able to enter the code? Paste the following link into your browser:

// https://github.com/users/Hank3r-gh/emails/253608093/confirm_verification/=
// 21337452?via_launch_code_email=3Dtrue

// ---
// Sent with <3 by GitHub.
// GitHub, Inc. 88 Colin P Kelly Jr Street
// San Francisco, CA 94107

// ----==_mimepart_6445153ca0e93_38bbc97c25a9
// Content-Type: text/html;
//  charset=UTF-8
// Content-Transfer-Encoding: 7bit

// <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
// <html xmlns="http://www.w3.org/1999/xhtml" xmlns="http://www.w3.org/1999/xhtml" lang="en" xml:lang="en" style="font-family: sans-serif; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%; box-sizing: border-box;" xml:lang="en">
//   <head>
//     <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
//     <meta name="viewport" content="width=device-width" />
//     <title></title>
    
//   </head>
//   <body style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot;; font-size: 14px; line-height: 1.5; color: #24292e; background-color: #fff; margin: 0;" bgcolor="#fff">
//     <table align="center" class="container-sm width-full" width="100%" style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; max-width: 544px; margin-right: auto; margin-left: auto; width: 100% !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//       <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//         <td class="center p-3" align="center" valign="top" style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 16px;">
//           <center style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//             <table border="0" cellspacing="0" cellpadding="0" align="center" class="width-full container-md" width="100%" style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; max-width: 768px; margin-right: auto; margin-left: auto; width: 100% !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <td align="center" style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">
//               <table style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tbody style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//       <td height="16" style="font-size: 16px; line-height: 16px; box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">&#160;</td>
//     </tr>
//   </tbody>
// </table>

//               <table border="0" cellspacing="0" cellpadding="0" align="left" width="100%" style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//                 <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//                   <td class="text-center" style="box-sizing: border-box; text-align: center !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;" align="center">
//                     <img src="https://github.githubassets.com/images/email/global/octocat-logo.png" alt="GitHub" width="32" style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; border-style: none;" />
//                     <h2 class="lh-condensed mt-2 text-normal" style="box-sizing: border-box; margin-top: 8px !important; margin-bottom: 0; font-size: 24px; font-weight: 400 !important; line-height: 1.25 !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//                         Here's your GitHub launch code, @Hank3r-gh!

//                     </h2>
//                   </td>
//                 </tr>
//               </table>
//               <table style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tbody style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//       <td height="16" style="font-size: 16px; line-height: 16px; box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">&#160;</td>
//     </tr>
//   </tbody>
// </table>

// </td>
//   </tr>
// </table>
//             <table width="100%" class="width-full" style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; width: 100% !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//               <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//                 <td class="border rounded-2 d-block" style="box-sizing: border-box; border-radius: 6px !important; display: block !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0; border: 1px solid #e1e4e8;">
//                   <table align="center" class="width-full text-center" style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; width: 100% !important; text-align: center !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//                     <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//                       <td class="p-3 p-sm-4" style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 24px;">
//                         <table border="0" cellspacing="0" cellpadding="0" align="center" class="width-full" width="100%" style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; width: 100% !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <td align="center" style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">
                          
// <img src="https://github.githubassets.com/images/email/signup/mona-launch-rocket.png" alt="an octocat standing next to a rocket" width="110" style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; border-style: none;" />

// <table style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tbody style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//       <td height="16" style="font-size: 16px; line-height: 16px; box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">&#160;</td>
//     </tr>
//   </tbody>
// </table>


// <span class="text-mono f4" style="box-sizing: border-box; font-size: 16px !important; font-family: &quot;SFMono-Regular&quot;,Consolas,&quot;Liberation Mono&quot;,Menlo,monospace !important;">
//   Continue signing up for GitHub by entering the code below:
// </span>

// <table style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tbody style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//       <td height="16" style="font-size: 16px; line-height: 16px; box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">&#160;</td>
//     </tr>
//   </tbody>
// </table>


// <span class="branch-name f00-light text-gray-dark text-mono" style="box-sizing: border-box; color: #24292e !important; display: inline-block; background-color: #eaf5ff; border-radius: 6px; padding: 2px 6px; font: 300 48px &quot;SFMono-Regular&quot;,Consolas,&quot;Liberation Mono&quot;, Menlo, monospace;">21337452</span>

// <table style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tbody style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//       <td height="16" style="font-size: 16px; line-height: 16px; box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">&#160;</td>
//     </tr>
//   </tbody>
// </table>


// <table width="100%" border="0" cellspacing="0" cellpadding="0" style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <td style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">
//       <table border="0" cellspacing="0" cellpadding="0" width="100%" style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//         <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//           <td align="center" style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">
//               <a href="https://github.com/account_verifications?via_launch_code_email=true" target="_blank" class="btn btn-primary btn-large" style="background-color: #28a745 !important; box-sizing: border-box; color: #fff; text-decoration: none; position: relative; display: inline-block; font-size: inherit; font-weight: 500; line-height: 1.5; white-space: nowrap; vertical-align: middle; cursor: pointer; -webkit-user-select: none; user-select: none; border-radius: .5em; -webkit-appearance: none; appearance: none; box-shadow: 0 1px 0 rgba(27,31,35,.1),inset 0 1px 0 rgba(255,255,255,.03); transition: background-color .2s cubic-bezier(0.3, 0, 0.5, 1); font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: .75em 1.5em; border: 1px solid #28a745;">Open GitHub</a>
//           </td>
//         </tr>
//       </table>
//     </td>
//   </tr>
// </table>




// </td>
//   </tr>
// </table>
//                       </td>
//                     </tr>
//                   </table>
//                 </td>
//               </tr>
//             </table>

//               <table border="0" cellspacing="0" cellpadding="0" align="center" class="width-full f5 text-gray-light" width="100%" style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; color: #6a737d !important; width: 100% !important; font-size: 14px !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <td align="center" style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">
//                 <table style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tbody style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//       <td height="16" style="font-size: 16px; line-height: 16px; box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">&#160;</td>
//     </tr>
//   </tbody>
// </table>

//                   <table class="width-full" style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; width: 100% !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tbody style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
      
//     <td style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">
//   <table style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//       <td class="text-left" style="box-sizing: border-box; text-align: left !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;" align="left">
//       Once completed, you can start using all of GitHub's features to explore, build, and share projects.

//       <table style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tbody style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//       <td height="16" style="font-size: 16px; line-height: 16px; box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">&#160;</td>
//     </tr>
//   </tbody>
// </table>


//       Not able to enter the code? Paste the following link into your browser: <br style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;" />
//       <a href="https://github.com/users/Hank3r-gh/emails/253608093/confirm_verification/21337452?via_launch_code_email=true" class="wb-break-all" style="background-color: transparent; box-sizing: border-box; color: #0366d6; text-decoration: none; word-break: break-all !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">https://github.com/users/Hank3r-gh/emails/253608093/confirm_verification/21337452?via_launch_code_email=true</a>
// </td>
//       <td style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;"></td>
//     </tr>
//   </table>
// </td>

//     </tr>
//   </tbody>
// </table>

//                 <hr style="box-sizing: content-box; height: 0; overflow: hidden; background-color: transparent; border-bottom-color: #dfe2e5; border-bottom-style: solid; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; margin: 15px 0; border-width: 0 0 1px;" />
// </td>
//   </tr>
// </table>

//             <table border="0" cellspacing="0" cellpadding="0" align="center" class="width-full text-center" width="100%" style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; width: 100% !important; text-align: center !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <td align="center" style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">
//               <table style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tbody style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//       <td height="16" style="font-size: 16px; line-height: 16px; box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">&#160;</td>
//     </tr>
//   </tbody>
// </table>

//                   <p class="f6" style="box-sizing: border-box; margin-top: 0; margin-bottom: 10px; font-size: 12px !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//                       <a href="https://github.com/settings/emails" class="d-inline-block" style="background-color: transparent; box-sizing: border-box; color: #0366d6; text-decoration: none; display: inline-block !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">Email preferences</a> &#12539;
//                       <a href="https://docs.github.com/articles/github-terms-of-service/" class="d-inline-block" style="background-color: transparent; box-sizing: border-box; color: #0366d6; text-decoration: none; display: inline-block !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">Terms</a> &#12539;
//                       <a href="https://docs.github.com/articles/github-privacy-policy/" class="d-inline-block" style="background-color: transparent; box-sizing: border-box; color: #0366d6; text-decoration: none; display: inline-block !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">Privacy</a> &#12539;
//                       <a href="https://github.com/login" class="d-inline-block" style="background-color: transparent; box-sizing: border-box; color: #0366d6; text-decoration: none; display: inline-block !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">Sign in to GitHub</a> 
//                   </p>
//               <table style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tbody style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//       <td height="16" style="font-size: 16px; line-height: 16px; box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">&#160;</td>
//     </tr>
//   </tbody>
// </table>

//               <p class="f5 text-gray-light" style="box-sizing: border-box; margin-top: 0; margin-bottom: 10px; color: #6a737d !important; font-size: 14px !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">  You&#8217;re receiving this email because you recently created a new GitHub account. If this wasn&#8217;t you, please ignore this email.
// </p>
// </td>
//   </tr>
// </table>
//             <table border="0" cellspacing="0" cellpadding="0" align="center" class="width-full text-center" width="100%" style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; width: 100% !important; text-align: center !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <td align="center" style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">
//   <table style="box-sizing: border-box; border-spacing: 0; border-collapse: collapse; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//   <tbody style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//     <tr style="box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">
//       <td height="16" style="font-size: 16px; line-height: 16px; box-sizing: border-box; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important; padding: 0;">&#160;</td>
//     </tr>
//   </tbody>
// </table>

//   <p class="f6 text-gray-light" style="box-sizing: border-box; margin-top: 0; margin-bottom: 10px; color: #6a737d !important; font-size: 12px !important; font-family: -apple-system,BlinkMacSystemFont,&quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot; !important;">GitHub, Inc. &#12539;88 Colin P Kelly Jr Street &#12539;San Francisco, CA 94107</p>
// </td>
//   </tr>
// </table>

//           </center>
//         </td>
//       </tr>
//     </table>
//     <!-- prevent Gmail on iOS font size manipulation -->
//    <div style="display: none; white-space: nowrap; box-sizing: border-box; font: 15px/0 apple-system, BlinkMacSystemFont, &quot;Segoe UI&quot;,Helvetica,Arial,sans-serif,&quot;Apple Color Emoji&quot;,&quot;Segoe UI Emoji&quot;;"> &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; &#160; </div>
//   </body>
// </html>

// ----==_mimepart_6445153ca0e93_38bbc97c25a9--
// `

// function extractData(data) {
//   const pattern = /Content-Type: multipart\/alternative; boundary="([^"]*)"/;
//   const match = data.match(pattern);
//   if (!match) {
//     throw new Error("Invalid data: missing boundary");
//   }
//   const boundary = match[1];
//   const split = data.split(`--${boundary}`);
//   if (split.length < 4) {
//     throw new Error("Invalid data: not enough parts");
//   }
//   return [ split[1].replace(/Content-Type:[^<]+/, ''), split[2].replace(/Content-Type:[^<]+/, '') ];
// }

// console.log(extractData(testData2));

// console.log(extractData(testData2)[0]);
// console.log(extractData(testData2)[1]);
// console.log(extractData(testData2)[2]);


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

