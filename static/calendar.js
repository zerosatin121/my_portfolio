fetch("/download-logs")
  .then(res => res.json())
  .then(logs => {
    const calendar = document.getElementById("calendar");
    const today = new Date();
    const year = today.getFullYear();
    const month = today.getMonth();

    const daysInMonth = new Date(year, month + 1, 0).getDate();
    const loggedDays = logs.map(log => new Date(log.log_date).getDate());

    for (let day = 1; day <= daysInMonth; day++) {
      const isLogged = loggedDays.includes(day);
      const cell = document.createElement("div");
      cell.textContent = day;
      cell.className = `p-2 border rounded ${isLogged ? "bg-green-300 font-bold" : "bg-white"}`;
      calendar.appendChild(cell);
    }
  });
