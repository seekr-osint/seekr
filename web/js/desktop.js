const frameContainer = document.querySelector('.frame-container');
const frame = document.querySelector('.frame');

let isDragging = false;
let currentX;
let currentY;
let initialX;
let initialY;
let xOffset = 0;
let yOffset = 0;

frame.addEventListener("mousedown", dragStart);
frame.addEventListener("mouseup", dragEnd);
frame.addEventListener("mouseout", dragEnd);
document.addEventListener("mousemove", drag);

function dragStart(e) {
  initialX = e.clientX - xOffset;
  initialY = e.clientY - yOffset;

  isDragging = true;
}

function dragEnd(e) {
  initialX = currentX;
  initialY = currentY;

  isDragging = false;
}

function drag(e) {
  if (isDragging) {
    e.preventDefault();
    currentX = e.clientX - initialX;
    currentY = e.clientY - initialY;

    xOffset = currentX;
    yOffset = currentY;

    setTranslate(currentX, currentY, frame);
  }
}

function setTranslate(xPos, yPos, el) {
  el.style.transform = "translate3d(" + xPos + "px, " + yPos + "px, 0)";
}
