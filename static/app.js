function play(e, i) {
    if (ap != undefined) {
        ap.destroy();
    }
    var ap = new APlayer({
        element: document.getElementById('player'),
        music: albums[i]
    });
    ap.play();
}
