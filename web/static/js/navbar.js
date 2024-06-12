let menu_icon = document.querySelector(".menu-icon");
let menu_list = document.querySelector(".links");

menu_icon.addEventListener("click", () => {
    menu_list.classList.toggle("active");
})
