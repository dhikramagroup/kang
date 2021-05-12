
window.onload = function () {
    var all_links = document.getElementById("nav").getElementsByTagName("a"),
        i = 0, len = all_links.length,
        full_path = location.href.split('#')[0]; //Ignore hashes?

    // Loop through each link.
    for (; i < len; i++) {
        if (all_links[i].href.split("#")[0] == full_path) {
            all_links[i].className += " active";
        }
    }
}

function myFunction(x) {
    const nav = document.querySelector("nav")
    x.classList.toggle("change");
    nav.classList.toggle("slide")
}
