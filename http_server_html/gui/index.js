"use strict"; // ES6
window.onload = () => {
    let elems = {
        mainContent: document.getElementById("mainContent"),
        changeContentBtn: document.getElementById("changeContentBtn")
    };

    let state = {
        nClicks: 0,
    };

    elems.changeContentBtn.addEventListener("click", () => {
        state.nClicks += 1;
        if (state.nClicks%2 === 1) {
            elems.mainContent.innerHTML = "<h3>Hihi</h3>"
        } else {
            elems.mainContent.innerHTML = "<h1>Hello</h1>"
        }
    });
};
