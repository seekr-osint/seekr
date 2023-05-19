// Sshhhh... This is a secret!

const specificTime = new Date();
specificTime.setHours(17);
specificTime.setMinutes(21);

const currentTime = new Date();

const webTitle = document.querySelector("title") as HTMLTitleElement;

const timeDifferenceInMilliseconds = Math.abs(currentTime.getTime() - specificTime.getTime());
const timeDifferenceInMinutes = Math.floor(timeDifferenceInMilliseconds / (1000 * 60));

if (timeDifferenceInMinutes <= 1) {
  webTitle.innerHTML = "GOX"; // GOX was supposed to be the name of the project before it was changed to seekr
}