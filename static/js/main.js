
//================================
//    chow and hide the nav
//================================

const HiddenNav = document.getElementsByClassName("hidden")[0]
const OpenFilter = document.getElementById("open")
const CloseFilter = document.getElementById("exit")

OpenFilter.addEventListener('click' , ()=> {
    HiddenNav.style.display = "flex"
})

CloseFilter.addEventListener('click' , ()=> {
    HiddenNav.style.display = "none"
})

//================================
//    Change value of range filter every time you change it
//================================

const Creation_Date_Input_min = document.getElementById("creationdate");
const Creation_Date_Label_min = document.getElementById("creationdatelabel");

const Creation_Date_Input_max = document.getElementById("creationdatemax");
const Creation_Date_Label_max = document.getElementById("creationdatelabelmax");

const First_Album_Input_min = document.getElementById("firstalbum");
const First_Album_Label_min = document.getElementById("firstalbumlabel");

const First_Album_Input_max = document.getElementById("firstalbummax");
const First_Album_Label_max = document.getElementById("firstalbumlabelmax");

// Event listeners for Creation Date range
Creation_Date_Input_min.addEventListener("input", () => {
    if (parseInt(Creation_Date_Input_min.value) > parseInt(Creation_Date_Input_max.value)) {
        Creation_Date_Input_min.value = Creation_Date_Input_max.value; // Prevent the min from being higher than the max
    }
    Creation_Date_Label_min.innerText = `min : ${Creation_Date_Input_min.value}`;
});

Creation_Date_Input_max.addEventListener("input", () => {
    if (parseInt(Creation_Date_Input_max.value) < parseInt(Creation_Date_Input_min.value)) {
        Creation_Date_Input_max.value = Creation_Date_Input_min.value; // Prevent the max from being lower than the min
    }
    Creation_Date_Label_max.innerText = `max : ${Creation_Date_Input_max.value}`;
});

// Event listeners for First Album range
First_Album_Input_min.addEventListener("input", () => {
    if (parseInt(First_Album_Input_min.value) > parseInt(First_Album_Input_max.value)) {
        First_Album_Input_min.value = First_Album_Input_max.value; // Prevent the min from being higher than the max
    }
    First_Album_Label_min.innerText = `min : ${First_Album_Input_min.value}`;
});

First_Album_Input_max.addEventListener("input", () => {
    if (parseInt(First_Album_Input_max.value) < parseInt(First_Album_Input_min.value)) {
        First_Album_Input_max.value = First_Album_Input_min.value; // Prevent the max from being lower than the min
    }
    First_Album_Label_max.innerText = `max : ${First_Album_Input_max.value}`;
});
