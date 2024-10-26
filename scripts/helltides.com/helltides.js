function data(isHC = false) {
  const lead = document.querySelector(".w-full.mb-12");
  const rows = [...lead.children].slice(1);
  return rows
    .map((el) => ({
      name: el.children[1].children[1].textContent,
      class: isHC
        ? el.children[1].children[0].children[1].src
            .split("icons/")[1]
            ?.replaceAll(".png", "")
        : el.children[1].children[0].firstElementChild.src
            .split("icons/")[1]
            ?.replaceAll(".png", ""),
      tier: el.children[2].children[0].textContent,
      time: el.children[2].children[1].textContent,
      build: el.children[4].children[0].href || "",
      video: el.children[4].children[1].href,
      mode: isHC ? "hardcore" : "softcore",
      season_id: 5,
    }))
    .filter((res) => !!res.class);
}
