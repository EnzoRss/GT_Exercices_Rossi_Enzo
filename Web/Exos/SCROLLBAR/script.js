


window.onload=function(){
    let scrolldiv = document.getElementById('scroll_div');
    let Is_Scrolling =false;
    let Is_over =false;
    let stop =null;

    scrolldiv.addEventListener('scroll',function () {
        console.log("ici")
        if (scrolldiv.scrollTop+scrolldiv.clientHeight>= scrolldiv.scrollHeight)
            Is_Scrolling = false;
        else if (scrolldiv.scrollTop === 0)
            Is_Scrolling =true;
    });

    scrolldiv.addEventListener("mouseenter", function (){
        Is_over=true;
    })
    scrolldiv.addEventListener("mouseleave", function (){
        Is_over=false;
    });

    document.addEventListener('wheel',function (Event){
        if (Is_Scrolling && Is_over){
            if ((Event.deltaY < 0 && scrolldiv.scrollTop === 0) || (Event.deltaY > 0 && scrolldiv.scrollTop + scrolldiv.clientHeight >= scrolldiv.scrollHeight)) {
                    Event.preventDefault()
            }
        }
        stop = setTimeout(function() {
            Is_Scrolling = false;
        }, 100);

        if (!Is_Scrolling && !Is_over) {
            clearTimeout(stop);
        }
    },{passive:false});
}