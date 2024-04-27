let counter = document.getElementById("count");
let counterValue = document.getElementById("counter");
let updateCount = document.getElementById("updateCount");
let count = 0;

document.addEventListener("keypress", () => {
    count++;
    counterValue.value = count;
    updateCount.value = `Add ${count} to counter`;
});
updateCount.addEventListener("click", async () => {
    const request = await fetch("/increase-counter", {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: `counter=${count}`,
    }).then((response) => response.json());
    count = 0;
    counterValue.value = 0;
    counter.textContent = request;
    updateCount.value = "Add 0 to counter";
});
