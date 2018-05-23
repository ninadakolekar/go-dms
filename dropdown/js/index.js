$('.dropdown-button').dropdown({
    inDuration: 300,
    outDuration: 225,
    constrain_width: true, // Does not change width of dropdown to that of the activator
    hover: true, // Activate on hover
    gutter: 0, // Spacing from edge
    belowOrigin: true,// Displays dropdown below the button
});

$(document).ready(function(){
    if($("dropdown2").css('display') == 'block'){
        $("dropdown2").css("left","0px");
    }
});

