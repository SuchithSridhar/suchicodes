let menu_icon = document.querySelector(".menu-icon");
let menu_list = document.querySelector(".links");

menu_icon.addEventListener("click", () => {
    menu_list.classList.toggle("active");
})

let gear_icon = document.querySelector(".gear-icon");
let settings_menu_container = document.querySelector(".settings-menu-container");
let settings_buttons = settings_menu_container.querySelectorAll("button");

function toggleSettingsMenu() {
    gear_icon.classList.toggle("active");
    settings_menu_container.classList.toggle("active");
    
    if (settings_menu_container.classList.contains("active")) {
        settings_buttons.forEach(button => {
            button.removeAttribute("tabindex");
        });
    } else {
        settings_buttons.forEach(button => {
            button.setAttribute("tabindex", "-1");
        });
    }
}

const enter_key_code = 13;
const space_key_code = 32;

gear_icon.addEventListener("click", toggleSettingsMenu);
gear_icon.addEventListener("keydown", event => {
    if (event.keyCode == enter_key_code || event.keyCode == space_key_code){
        event.preventDefault();
        toggleSettingsMenu();
    }
});

function detectColorScheme() {
    let theme = "light";

    let stored_theme = localStorage.getItem("theme");
    if (stored_theme && stored_theme == "dark") {
        theme = "dark";
    } else if (stored_theme && stored_theme == "light") {
        theme = "light";
    } else if (!window.matchMedia) {
        theme = "light";
    } else if (window.matchMedia("(prefers-color-scheme: dark)").matches) {
        theme = "dark";
    }

    if (theme == "dark") {
        document.documentElement.setAttribute("data-theme", "dark");
    }
}

function handleThemeChange(target) {
    let theme = target.getAttribute("data-mode").toLowerCase()

    if (theme == "system") {
        localStorage.setItem('theme', 'system');

        if (window.matchMedia("(prefers-color-scheme: dark)").matches) {
            document.documentElement.setAttribute('data-theme', 'dark');
        } else {
            document.documentElement.setAttribute('data-theme', 'light');
        }

    } else {
        localStorage.setItem('theme', theme);
        document.documentElement.setAttribute('data-theme', theme);
    }


    if (target.classList.contains("active")) {
        return
    } else {
        target.classList.add('active')
    }

    theme_buttons.forEach(button => {
        if (button != target && button.classList.contains('active')) {
            button.classList.remove('active');
        }
    })
}

detectColorScheme();
let stored_theme = localStorage.getItem("theme");

let theme_buttons = settings_menu_container.querySelectorAll('.theme-toggler button');
theme_buttons.forEach(button => {
    if (stored_theme == button.getAttribute("data-mode").toLowerCase()) {
        button.classList.add("active");
    }

    button.addEventListener('click', event => {
        handleThemeChange(event.target);
    });
    button.addEventListener('keydown', event => {
        if (event.keyCode == enter_key_code || event.keyCode == space_key_code){
            event.preventDefault();
            handleThemeChange(event.target);
        }
    })
})
