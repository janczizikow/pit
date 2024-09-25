function data() {
  const lead = document.querySelector(".w-full.mb-12");
  const rows = [...lead.children].slice(1);

  var cls = rows[0].children[1].children[0].firstElementChild.src
    .split("icons/")[1]
    .replaceAll(".png", "");
  const name = rows[0].children[1].children[1].textContent;
  const tier = rows[0].children[2].children[0].textContent;
  const time = rows[0].children[2].children[1].textContent;
  const build = rows[0].children[4].children[0].href;
  const video = rows[0].children[4].children[1].href;

  return rows
    .map((el) => ({
      name: el.children[1].children[1].textContent,
      class: el.children[1].children[0].firstElementChild.src
        .split("icons/")[1]
        ?.replaceAll(".png", ""),
      tier: el.children[2].children[0].textContent,
      time: el.children[2].children[1].textContent,
      build: el.children[4].children[0].href || "",
      video: el.children[4].children[1].href,
      mode: "softcore",
      season_id: 5,
    }))
    .filter((res) => !!res.class);
}
