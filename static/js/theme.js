const btn = document.getElementById("theme-toggle");
const icon = document.getElementById("theme-icon");
const root = document.documentElement;

function updateIcon(theme) {
  if (theme === "light") {
    icon.className = "fas fa-moon"; // clair
  } else {
    icon.className = "fas fa-lightbulb";  // sombre
  }
}

let savedTheme = localStorage.getItem("theme");

if (!savedTheme) {
  savedTheme = "dark";
  localStorage.setItem("theme", "dark");
}

root.setAttribute("data-theme", savedTheme);
updateIcon(savedTheme);

btn.addEventListener("click", (e) => {
  e.preventDefault();

  const current = root.getAttribute("data-theme");
  const newTheme = current === "light" ? "dark" : "light";

  root.setAttribute("data-theme", newTheme);
  localStorage.setItem("theme", newTheme);

  updateIcon(newTheme);
});