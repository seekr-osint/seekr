const elements: {[key: string]: HTMLCollectionOf<Element>} = {
  "country-select": document.getElementsByClassName("country-select")
};

for (const [className, nodeList] of Object.entries(elements)) {
  for (let i = 0; i < nodeList.length; i++) {
    const x = nodeList[i] as HTMLElement;
    const selElmnt = x.getElementsByTagName("select")[0] as HTMLSelectElement;
    const ll = selElmnt.length;
    let a = document.createElement("DIV");
    a.setAttribute("class", "select-selected");
    a.innerHTML = selElmnt.options[selElmnt.selectedIndex].innerHTML;
    x.appendChild(a);
    let b = document.createElement("DIV");
    b.setAttribute("class", "select-items select-hide");
    for (let j = 1; j < ll; j++) {
      const c = document.createElement("DIV");
      c.innerHTML = selElmnt.options[j].innerHTML;
      c.addEventListener("click", function (e) {
        if (this.parentNode && this.parentNode.parentNode && this.parentNode.parentNode.querySelectorAll("select")[0]) {
          const y = this.parentNode.parentNode.querySelectorAll("select")[0] as HTMLSelectElement;
          const h = this.parentNode.previousSibling as HTMLElement;

          for (let k = 0; k < y.length; k++) {
            if (y.options[k].innerHTML == this.innerHTML) {
              y.selectedIndex = k;
              h.innerHTML = this.innerHTML;
              let yl = this.parentNode.querySelector(".same-as-selected") as HTMLSelectElement;
              if (yl) {
                for (let l = 0; l < yl.length; l++) {
                  yl[l].removeAttribute("class");
                }
                this.setAttribute("class", "same-as-selected");
                break;
              }
            }
          }
          h.click();
        }
      });
      b.appendChild(c);
    }
    x.appendChild(b);
    a.addEventListener("click", function (e) {
      e.stopPropagation();
      closeAllSelectOld(this);
      if (this.nextSibling) {
        const s = this.nextSibling as HTMLElement;
        s.classList.toggle("select-hide");
        this.classList.toggle("select-arrow-active");
      }
    });
  }
}

function closeAllSelectOld(elmnt: HTMLElement) {
  const arrNo = [];
  const x = document.getElementsByClassName("select-items") as HTMLCollectionOf<HTMLElement>;
  const y = document.getElementsByClassName("select-selected") as HTMLCollectionOf<HTMLElement>;
  for (let i = 0; i < y.length; i++) {
    if (elmnt == y[i]) {
      arrNo.push(i);
    } else {
      y[i].classList.remove("select-arrow-active");
    }
  }
  for (let i = 0; i < x.length; i++) {
    if (arrNo.indexOf(i)) {
      x[i].classList.add("select-hide");
    }
  }
}

document.addEventListener("click", function() {
  closeAllSelectOld(this.activeElement as HTMLElement);
});