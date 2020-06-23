function byid(id) {
    return document.getElementById(id);
}
function valof(id) {
    return byid(id).checked ? 1 : 0;
}
function text(e) {
    if (e == null) {
        return "";
    }
    return e.value;
}

var lastQ = null
var ticking = false
var ticksSinceChange = 0

function update() {
    const q = theImage();

    const req = new XMLHttpRequest();
    req.onreadystatechange = function() {
        if(req.readyState === 4 && req.status === 200) {
            byid("output").innerHTML = req.responseText
            //	if(history && history.replaceState)
            //		history.replaceState(null, null, "?q=" + encodeURIComponent(q))
        }
    }
    req.open("GET", q, true)
    req.send(null)
}

function theImage() {
    return "/qr/draw?i=" + img + "&u=" + encodeURI(byid("url").value) + "&m=" + m + "&x=" + dx + "&y=" + dy + "&v=" + v + "&c=" + valof("control") + "&r=" + valof("rand") + "&d=" + valof("data") + "&t=" + valof("dither") + "&s=" + Math.floor(Math.random()*1e9) + "&z=" + z + "&o=" + rotation
}


var dx=4
var dy=4
var url="http://research.swtch.com/qart"
var m=2
var v=6
var z=0
var img="pjw"
var cheat=0
var rotation=0

function up() { dy++; update(); }
function down() { dy--; update(); }
function left() { dx++; update(); }
function right() { dx--; update(); }
function bigger() { if(v < 8) { v++; update(); } }
function smaller() { if(v > 1) { v--; update(); } }
function setimg(s) { img=s; update(); }
function togglemask() { if(m>=0) m=-1; else m=2; update(); }
function ibigger() { z++; update(); }
function ismaller() { z--; update(); }
function rotate() { rotation = (rotation+1) & 3; update(); }
