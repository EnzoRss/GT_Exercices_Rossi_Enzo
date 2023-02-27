let container =document.getElementById("container");
let g =document.getElementById("g")
let d =document.getElementById("d");
let p=0;
let nbr= 9;

function afficherMasquer(){
if(p ===-nbr+1)
    d.style.visibility="hidden";
else
    d.style.visibility="visible";
    if(p===0)
        g.style.visibility="hidden";
    else
        g.style.visibility="visible";
}

document.body.onload =function (){

    container.style.width=(400*nbr)+"px";
    for (let i=1 ; i<=nbr ;i++){
        let  div=document.createElement("div")
        div.className ="photo"
        div.style.backgroundImage="url(image/"+i+".jpg";
        container.appendChild(div);
    }
    afficherMasquer()
}

d.onclick =function (){
    if (p>-nbr+1 )
        p--;
    container.style.transform="translate("+p*400+"px)";
    container.style.transition ="all 0.5s ease"
    afficherMasquer()
}

g.onclick =function (){
    if (p<0 )
        p++;
    container.style.transform="translate("+p*400+"px)";
    container.style.transition ="all 0.5s ease"
    afficherMasquer()
}

