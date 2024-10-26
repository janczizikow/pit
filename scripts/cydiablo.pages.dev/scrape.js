function data(seasonId = 6) {
  const table = document.querySelector("tbody");
  const rows = [...table.children];
  return rows
    .map((r) => ({
      name: r.children[1].textContent,
      class: r.children[2].textContent.toLowerCase(),
      tier: r.children[7].textContent,
      time: r.children[8].textContent,
      build:
        r.children[9].textContent === "-" ? "" : r.children[9].firstChild.href,
      video: r.children[10].firstChild.href,
      mode:
        r.children[1].className === "text-warning" ? "hardcore" : "softcore",
      season_id: seasonId,
    }))
    .filter((res) => !!res.class);
}
