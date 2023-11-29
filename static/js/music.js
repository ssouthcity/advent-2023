const audioEl = document.querySelector("#audio-player audio");
const buttonsEl = document.querySelectorAll("button[data-play]");

for (const button of buttonsEl) {
  button.addEventListener("click", (event) => {
    const audioSrc = event.target.getAttribute("data-play");

    audioEl.setAttribute("src", audioSrc);
    audioEl.load();
    audioEl.play();
  });
}
