const checkboxesEl = document.querySelectorAll("input[type='checkbox'][data-day]")

const storedState = localStorage.getItem("openedWindows") != null
  ? JSON.parse(localStorage.getItem("openedWindows"))
  : {};

for (const checkboxEl of checkboxesEl) {
  const day = checkboxEl.getAttribute("data-day");
  if (day in storedState) {
    checkboxEl.checked = storedState[day];
  }

  checkboxEl.addEventListener("change", (e) => {
    const day = e.target.getAttribute("data-day");
    storedState[day] = e.target.checked;
    localStorage.setItem("openedWindows", JSON.stringify(storedState))
  })
}
