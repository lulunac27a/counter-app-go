let counter = document.getElementById("count"); //key count since `Add # to counter` button is pressed
let squareCounter = document.getElementById("squareCount"); //square count
let counterValue = document.getElementById("counter"); //counter value
let updateCount = document.getElementById("updateCount"); //
let count = 0; //set key count to 0 when page is loaded

document.addEventListener("keypress", () => {
    //increase key count when key is pressed
    count++;
    counterValue.value = count;
    updateCount.value = `Add ${count} to counter`; //update button text
});
updateCount.addEventListener("click", async () => {
    //when `Add # to counter` button is pressed
    const request = await fetch("/increase-counter", {
        method: "POST",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
        body: `counter=${count}`,
    }).then((response) => response.json());
    count = 0; //set key counter to 0
    counterValue.value = 0;
    counter.textContent = request["count"];
    squareCounter.textContent = request["squareCount"];
    updateCount.value = "Add 0 to counter";
});
