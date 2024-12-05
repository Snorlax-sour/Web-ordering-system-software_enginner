const Account_info = document.querySelector(".Account_info");
const Purple_button = document.querySelector(".Purple_button");
const form = document.getElementById("Input_bar");
function Read_content(){
    if(Account_info.value === ""){
        return ;
        
    }
    else{
        console.log(Account_info.value);

    }
}

Account_info.addEventListener("keyup", function(e){
    if(e.key == "Enter"){
        Read_content();
    }
});
Purple_button.addEventListener("click", Purple_button);