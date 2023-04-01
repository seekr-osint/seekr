function delay(time) { // Because there is no default sleep function
  return new Promise(resolve => setTimeout(resolve, time));
}

function SaveAsFile(t, f, m) {
  // SaveAsFile("text","filename.txt","text/plain;charset=utf-8");

  try {
    var b = new Blob([t],{type:m});
    saveAs(b, f);
  } catch (e) {
    window.open("data:"+m+"," + encodeURIComponent(t), '_blank','');
  }
}


document.querySelectorAll("span").forEach(function (element) {
  element.addEventListener('paste', function (e) {
    // Prevent the default action
    e.preventDefault();

    // Get the copied text from the clipboard
    const text = e.clipboardData
      ? (e.originalEvent || e).clipboardData.getData('text/plain')
      : // For IE
      window.clipboardData
      ? window.clipboardData.getData('Text')
      : '';

    if (document.queryCommandSupported('insertText')) {
      document.execCommand('insertText', false, text);
    } else {
      // Insert text at the current position of caret
      const range = document.getSelection().getRangeAt(0);
      range.deleteContents();

      const textNode = document.createTextNode(text);
      range.insertNode(textNode);
      range.selectNodeContents(textNode);
      range.collapse(false);

      const selection = window.getSelection();
      selection.removeAllRanges();
      selection.addRange(range);
    }
  });
});

[document.getElementById("c-name-tag"), document.getElementById("acc-name-tag")].forEach(item => {
  item.addEventListener('paste', function (e) {
    // Prevent the default action
    e.preventDefault();
  
    // Get the copied text from the clipboard
    const text = e.clipboardData
        ? (e.originalEvent || e).clipboardData.getData('text/plain')
        : // For IE
        window.clipboardData
        ? window.clipboardData.getData('Text')
        : '';
  
    if (document.queryCommandSupported('insertText')) {
        document.execCommand('insertText', false, text);
    } else {
        // Insert text at the current position of caret
        const range = document.getSelection().getRangeAt(0);
        range.deleteContents();
  
        const textNode = document.createTextNode(text);
        range.insertNode(textNode);
        range.selectNodeContents(textNode);
        range.collapse(false);
  
        const selection = window.getSelection();
        selection.removeAllRanges();
        selection.addRange(range);
    }
  });
});


export { delay, SaveAsFile };